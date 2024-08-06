//
// Copyright (c) 2020-2024 Markku Rossi
//
// All rights reserved.
//

package nap

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gopacket/gopacket"
	"github.com/gopacket/gopacket/layers"
)

const dohServer = "https://cloudflare-dns.com/dns-query"

var decodeOptions = gopacket.DecodeOptions{
	Lazy:   true,
	NoCopy: true,
}

// DNSQuery implements handler for DNS queries.
func DNSQuery(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		Errorf(w, http.StatusBadRequest, "Invalid method %s", r.Method)
		return
	}
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		Errorf(w, http.StatusInternalServerError,
			"Error reading request body: %s", err)
		return
	}

	packet := gopacket.NewPacket(data, layers.LayerTypeDNS, decodeOptions)
	layer := packet.Layer(layers.LayerTypeDNS)
	if layer == nil {
		Errorf(w, http.StatusBadRequest, "Request did not contain DNS query")
		return
	}
	dns := layer.(*layers.DNS)
	for _, q := range dns.Questions {
		labels := NewLabels(string(q.Name))
		log.Printf("q: %s\n", labels)
	}

	response, ok := doh(w, data)
	if !ok {
		return
	}

	w.Header().Set("Content-Type", "application/dns-message")
	w.Write(response)
}

func doh(w http.ResponseWriter, data []byte) ([]byte, bool) {
	dnsReq, err := http.NewRequest("POST", dohServer, bytes.NewReader(data))
	if err != nil {
		Errorf(w, http.StatusInternalServerError, "HTTP new request: %s", err)
		return nil, false
	}
	dnsReq.Header.Set("Content-Type", "application/dns-message")

	dnsResp, err := httpClient.Do(dnsReq)
	if err != nil {
		Errorf(w, http.StatusInternalServerError, "HTTP request: %s", err)
		return nil, false
	}
	defer dnsResp.Body.Close()

	dnsRespData, err := ioutil.ReadAll(dnsResp.Body)
	if err != nil {
		Errorf(w, http.StatusBadGateway,
			"error reading server response: %s", err)
		return nil, false
	}
	if dnsResp.StatusCode != http.StatusOK {
		Errorf(w, http.StatusBadGateway, "status=%s, content:\n%s",
			dnsResp.Status, hex.Dump(dnsRespData))
	}

	fmt.Printf("dnsRespData:\n%s", hex.Dump(dnsRespData))

	return dnsRespData, true
}
