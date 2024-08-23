//
// Copyright (c) 2024 Markku Rossi
//
// All rights reserved.
//

package pki

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math"
	"math/big"
	"os"
	"time"
)

type CA struct {
	name string
	priv *ecdsa.PrivateKey
	Cert *x509.Certificate
}

func CreateCA(name string) (*CA, error) {
	priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, err
	}
	serial, err := rand.Int(rand.Reader, big.NewInt(math.MaxInt64))
	if err != nil {
		return nil, err
	}

	subject := pkix.Name{
		Country:      []string{"FI"},
		Organization: []string{"Ephemelier"},
		CommonName:   "Ephemelier Root X1",
	}
	now := time.Now()

	caTmpl := &x509.Certificate{
		SignatureAlgorithm: x509.ECDSAWithSHA512,
		SerialNumber:       serial,
		Subject:            subject,
		NotBefore:          now,
		NotAfter:           now.Add(time.Hour * 24 * 365 * 20),
		KeyUsage: x509.KeyUsageDigitalSignature |
			x509.KeyUsageCertSign | x509.KeyUsageCRLSign,
		ExtKeyUsage: []x509.ExtKeyUsage{
			x509.ExtKeyUsageServerAuth,
			x509.ExtKeyUsageClientAuth,
		},
		BasicConstraintsValid: true,
		IsCA:                  true,
		MaxPathLenZero:        true,
		CRLDistributionPoints: []string{
			"http://ephemelier.com/crl",
		},
	}
	der, err := x509.CreateCertificate(rand.Reader, caTmpl, caTmpl,
		&priv.PublicKey, priv)
	if err != nil {
		return nil, err
	}
	cert, err := x509.ParseCertificate(der)
	if err != nil {
		return nil, err
	}

	err = savePrivateKey(privateKeyName(name), priv)
	if err != nil {
		return nil, err
	}
	err = os.WriteFile(certName(name), cert.Raw, 0666)
	if err != nil {
		return nil, err
	}

	return &CA{
		name: name,
		priv: priv,
		Cert: cert,
	}, nil
}

func OpenCA(name string) (*CA, error) {
	priv, err := loadPrivateKey(privateKeyName(name))
	if err != nil {
		return nil, err
	}
	der, err := os.ReadFile(certName(name))
	if err != nil {
		return nil, err
	}
	cert, err := x509.ParseCertificate(der)
	if err != nil {
		return nil, err
	}

	return &CA{
		name: name,
		priv: priv,
		Cert: cert,
	}, nil
}

func (ca *CA) CreateEEKey() (crypto.PrivateKey, crypto.PublicKey, error) {
	priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, nil, err
	}
	return priv, &priv.PublicKey, nil
}

func (ca *CA) CreateCertificate(tmpl *x509.Certificate, pub any) (
	*x509.Certificate, error) {

	serial, err := rand.Int(rand.Reader, big.NewInt(math.MaxInt64))
	if err != nil {
		return nil, err
	}
	now := time.Now()
	tmpl.SignatureAlgorithm = x509.ECDSAWithSHA512
	tmpl.SerialNumber = serial
	tmpl.NotBefore = now
	tmpl.NotAfter = now.Add(time.Hour * 24)
	tmpl.KeyUsage = x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment
	tmpl.ExtKeyUsage = []x509.ExtKeyUsage{
		x509.ExtKeyUsageServerAuth,
		x509.ExtKeyUsageClientAuth,
	}
	tmpl.BasicConstraintsValid = true

	der, err := x509.CreateCertificate(rand.Reader, tmpl, ca.Cert, pub, ca.priv)
	if err != nil {
		return nil, err
	}
	return x509.ParseCertificate(der)
}

func privateKeyName(name string) string {
	return name + ".prv"
}

func certName(name string) string {
	return name + ".crt"
}

func savePrivateKey(name string, privateKey *ecdsa.PrivateKey) error {
	x509Encoded, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		return err
	}
	data := pem.EncodeToMemory(&pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509Encoded,
	})
	return os.WriteFile(name, data, 0600)
}

func loadPrivateKey(name string) (*ecdsa.PrivateKey, error) {
	pemEncoded, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(pemEncoded)
	if block == nil {
		return nil, fmt.Errorf("invalid PEM data")
	}
	return x509.ParseECPrivateKey(block.Bytes)
}
