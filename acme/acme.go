//
// Copyright (c) 2024-2025 Markku Rossi
//
// All rights reserved.
//

// Package acme implements Let's Encrypt ACME integration.
package acme

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"

	"github.com/go-acme/lego/v4/certcrypto"
	"github.com/go-acme/lego/v4/certificate"
	"github.com/go-acme/lego/v4/challenge/http01"
	"github.com/go-acme/lego/v4/lego"
	"github.com/go-acme/lego/v4/registration"
)

// CA endpoints.
var (
	StagingDirectory = "https://acme-staging-v02.api.letsencrypt.org/directory"
	Directory        = "https://acme-v02.api.letsencrypt.org/directory"
)

var (
	_ registration.User = &Client{}
)

// Client implements the ACME client. It also implements the
// registration.User interface.
type Client struct {
	hostname     string
	email        string
	dryRun       bool
	priv         crypto.PrivateKey
	acme         *lego.Client
	registration *registration.Resource
}

// GetEmail implements User.GetEmail.
func (c *Client) GetEmail() string {
	return c.email
}

// GetRegistration implements User.GetRegistration.
func (c *Client) GetRegistration() *registration.Resource {
	return c.registration
}

// GetPrivateKey implements User.GetPrivateKey.
func (c *Client) GetPrivateKey() crypto.PrivateKey {
	return c.priv
}

// New creates a new ACME client.
func New(hostname, email string, dryRun bool) (*Client, error) {
	priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, err
	}
	client := &Client{
		hostname: hostname,
		email:    email,
		dryRun:   dryRun,
		priv:     priv,
	}

	config := lego.NewConfig(client)

	if dryRun {
		config.CADirURL = StagingDirectory
	} else {
		config.CADirURL = Directory
	}
	config.Certificate.KeyType = certcrypto.EC256

	client.acme, err = lego.NewClient(config)
	if err != nil {
		return nil, err
	}

	err = client.acme.Challenge.SetHTTP01Provider(
		http01.NewProviderServer("", "80"))
	if err != nil {
		return nil, err
	}

	// New users will need to register
	reg, err := client.acme.Registration.Register(registration.RegisterOptions{
		TermsOfServiceAgreed: true,
	})
	if err != nil {
		return nil, err
	}
	client.registration = reg

	request := certificate.ObtainRequest{
		Domains: []string{hostname},
		Bundle:  true,
	}
	resource, err := client.acme.Certificate.Obtain(request)
	if err != nil {
		return nil, err
	}

	// Each certificate comes back with the cert bytes, the bytes of
	// the client's private key, and a certificate URL. SAVE THESE TO
	// DISK.
	fmt.Printf(`Certificate: {
    Domain           : %v
    CertURL          : %v
    CertStableURL    : %v
    PrivateKey       : %x
    Certificate      : %x
    IssuerCertificate: %x
    CSR              : %x
}
`,
		resource.Domain,
		resource.CertURL,
		resource.CertStableURL,
		resource.PrivateKey,
		resource.Certificate,
		resource.IssuerCertificate,
		resource.CSR)

	return client, nil
}
