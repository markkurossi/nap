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
	"net/http"

	"github.com/gopacket/gopacket"
	"github.com/gopacket/gopacket/layers"
)

const dohServer = "https://cloudflare-dns.com/dns-query"

var decodeOptions = gopacket.DecodeOptions{
	Lazy:   true,
	NoCopy: true,
}

var serializeOptions = gopacket.SerializeOptions{
	FixLengths:       true,
	ComputeChecksums: true,
}

// DNSQuery implements handler for DNS queries.
func DNSQuery(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	if len(cids) > 0 {
		cid := q.Get("client_id")
		_, ok := cids[cid]
		if !ok {
			Errorf(w, http.StatusUnauthorized, "Invalid client ID")
			return
		}
	}
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
	name := q.Get("block")
	if len(name) == 0 {
		name = "default"
	}
	blacklist, ok := blacklists[name]
	if !ok {
		logError.Printf("unknown blacklist: %s", name)
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
		for _, black := range blacklist {
			if labels.Match(black) {
				logInfo.Printf("block: %s (%s)", labels, black)
				response, err := nonExistingDomain(dns)
				if err != nil {
					Errorf(w, http.StatusInternalServerError,
						"non-existing domain: %v", err)
					return
				}
				w.Header().Set("Content-Type", "application/dns-message")
				w.Write(response)
				return
			}
		}
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

	return dnsRespData, true
}

func nonExistingDomain(q *layers.DNS) ([]byte, error) {
	var responseLayers []gopacket.SerializableLayer

	responseLayers = append(responseLayers, &layers.DNS{
		ID:           q.ID,
		QR:           true,
		OpCode:       q.OpCode,
		AA:           true,
		TC:           false,
		RD:           q.RD,
		RA:           false,
		ResponseCode: layers.DNSResponseCodeNXDomain,
		Questions:    q.Questions,
	})

	buffer := gopacket.NewSerializeBuffer()
	err := gopacket.SerializeLayers(buffer, serializeOptions, responseLayers...)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func cname(q *layers.DNS, cname string) ([]byte, error) {
	var responseLayers []gopacket.SerializableLayer

	if len(q.Questions) != 1 {
		return nil, fmt.Errorf("cname: expected 1 question, got %v",
			len(q.Questions))
	}

	responseLayers = append(responseLayers, &layers.DNS{
		ID:           q.ID,
		QR:           true,
		OpCode:       q.OpCode,
		AA:           true,
		TC:           false,
		RD:           q.RD,
		RA:           true,
		ResponseCode: layers.DNSResponseCodeNoErr,
		Questions:    q.Questions,
		Answers: []layers.DNSResourceRecord{
			{
				Name:  q.Questions[0].Name,
				Type:  layers.DNSTypeCNAME,
				Class: layers.DNSClassIN,
				TTL:   60,
				CNAME: []byte(cname),
			},
		},
	})

	buffer := gopacket.NewSerializeBuffer()
	err := gopacket.SerializeLayers(buffer, serializeOptions, responseLayers...)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
