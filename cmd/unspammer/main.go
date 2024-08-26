//
// Copyright (c) 2024 Markku Rossi
//
// All rights reserved.
//

package main

import (
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"flag"
	"fmt"
	"html"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/markkurossi/nap/pki"
)

var (
	certificates = make(map[string]*tls.Certificate)
)

func main() {
	caName := flag.String("ca", "", "The name of the CA")
	createCA := flag.Bool("create-ca", false, "Create CA")
	addr := flag.String("addr", ":443", "Address to listen")
	flag.Parse()

	if len(*caName) == 0 {
		log.Fatal("CA name not specified")
	}

	var ca *pki.CA
	var err error

	if *createCA {
		ca, err = pki.CreateCA(*caName)
	} else {
		ca, err = pki.OpenCA(*caName)
	}
	if err != nil {
		log.Fatal(err)
	}
	eePriv, eePub, err := ca.CreateEEKey()
	if err != nil {
		log.Fatal(err)
	}
	eeTmpl := &x509.Certificate{}

	tlsConfig := &tls.Config{
		GetCertificate: func(info *tls.ClientHelloInfo) (*tls.Certificate, error) {
			fmt.Printf("ServerName: %v\n", info.ServerName)
			tlsCert, ok := certificates[info.ServerName]
			if !ok {
				eeTmpl.Subject = pkix.Name{
					CommonName: info.ServerName,
				}
				eeTmpl.DNSNames = []string{info.ServerName}

				cert, err := ca.CreateCertificate(eeTmpl, eePub)
				if err != nil {
					return nil, err
				}
				tlsCert = &tls.Certificate{
					Certificate: [][]byte{
						cert.Raw,
						ca.Cert.Raw,
					},
					PrivateKey: eePriv,
					Leaf:       cert,
				}
				certificates[info.ServerName] = tlsCert
			}
			return tlsCert, nil
		},
	}

	s := &http.Server{
		Addr:           *addr,
		Handler:        http.HandlerFunc(handler),
		TLSConfig:      tlsConfig,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServeTLS("", ""))
}

var cors = map[string]string{
	"Access-Control-Allow-Methods":     "GET,POST,OPTIONS",
	"Access-Control-Allow-Origin":      "https://www.mtvuutiset.fi",
	"Access-Control-Max-Age":           "1728000",
	"Date":                             "Fri, 23 Aug 2024 04:39:12 GMT",
	"Access-Control-Allow-Credentials": "true",
	"Access-Control-Allow-Headers":     "content-type, depth, user-agent, x-file-size, x-requested-with, if-modified-since, x-file-name, cache-control",
	"P3p":                              `policyref="https://www.freewheel.tv/w3c/p3p.xml",CP="ALL DSP COR NID"`,
	// "Connection":                       "keep-alive",
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s: %s\n", r.Method, r.URL.Path)
	if false {
		for k, values := range r.Header {
			for _, v := range values {
				fmt.Printf(" - %v: %v\n", k, v)
			}
		}
	}

	if r.Method == "OPTIONS" {

		if false {
			setHdr(r, w, "Access-Control-Request-Private-Network",
				"Private-Network")
			w.Header().Set("Access-Control-Allow-Headers", "*")
			w.Header().Set("Access-Control-Allow-Methods",
				"GET,HEAD,PUT,PATCH,POST,DELETE,OPTIONS")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Max-Age", "240")
			w.WriteHeader(204)
		}
		for k, v := range cors {
			w.Header().Set(k, v)
		}
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(200)
		return
	}

	if strings.HasPrefix(r.URL.Path, "/ad/") {
		for k, v := range cors {
			w.Header().Set(k, v)
		}
		w.Header().Set("Content-Type", "text/xml")

		q := r.URL.Query()
		switch q.Get("resp") {
		case "vast4":
			fmt.Printf(" - returning vast4\n")
			fmt.Fprint(w, `<VAST version='4.1'
      xmlns:xsi='http://www.w3.org/2001/XMLSchema-instance'
      xsi:noNamespaceSchemaLocation='vast.xsd'>
</VAST>
`)
		case "vmap1+vast4":
			fmt.Printf(" - returning vmap1+vast4\n")
			fmt.Fprint(w, `<vmap:VMAP version='1.0'
           xmlns:vmap='http://www.iab.net/vmap-1.0'>
  <vmap:AdBreak breakId='0.0.0.2046985237'
                breakType='linear'
                timeOffset='start'>
    <vmap:AdSource allowMultipleAds='true' followRedirects='true' id='1'>
      <vmap:VASTAdData>
        <VAST version='4.1'
              xmlns:xsi='http://www.w3.org/2001/XMLSchema-instance'
              xsi:noNamespaceSchemaLocation='vast.xsd'/>
      </vmap:VASTAdData>
    </vmap:AdSource>
    <vmap:TrackingEvents>
      <vmap:Tracking event='breakEnd'>
        <![CDATA[https://805ba.v.fwmrm.net/ad/l/1?s=l0d8e&n=525754%3B525754%3B512166%3B512167%3B512188%3B517424&t=1724414563520472336&f=262144&cn=videoView&et=i&uxnw=&uxss=&uxct=&init=1&vcid2=1f928695-eaa3-4f7c-a51d-09a877c1164e]]>
      </vmap:Tracking>
      <vmap:Tracking event='breakStart'>
        <![CDATA[https://805ba.v.fwmrm.net/ad/l/1?s=l0d8e&n=525754%3B525754%3B512166%3B512167%3B512188%3B517424&t=1724414563520472336&f=262144&cn=slotImpression&et=i&tpos=0&init=1&slid=0,1,2]]>
      </vmap:Tracking>
    </vmap:TrackingEvents>
  </vmap:AdBreak>
</vmap:VMAP>`)

		default:
			fmt.Printf(" - unknown resp: %s\n", q.Get("resp"))
			fmt.Printf(" - q: %s\n", r.RequestURI)
			for k, v := range q {
				fmt.Printf(" - %s=%v\n", k, v)
			}
		}
	} else {
		fmt.Fprintf(w, "Hello, %q\n", html.EscapeString(r.URL.String()))
	}
}

func setHdr(r *http.Request, w http.ResponseWriter, req, resp string) {
	val := r.Header.Get(req)
	if len(val) > 0 {
		w.Header().Set("Access-Control-Allow-"+resp, val)
	}
}
