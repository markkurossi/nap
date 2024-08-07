//
// Copyright (c) 2020-2024 Markku Rossi
//
// All rights reserved.
//

package nap

import (
	"fmt"
	"testing"

	"github.com/gopacket/gopacket"
	"github.com/gopacket/gopacket/layers"
)

var msgs = [][]byte{
	{
		0x00, 0x00, 0x01, 0x00, 0x00, 0x01, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x01, 0x07, 0x63, 0x6f, 0x6e,
		0x73, 0x6f, 0x6c, 0x65, 0x05, 0x63, 0x6c, 0x6f,
		0x75, 0x64, 0x06, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
		0x65, 0x03, 0x63, 0x6f, 0x6d, 0x00, 0x00, 0x41,
		0x00, 0x01, 0x00, 0x00, 0x29, 0x10, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x4b, 0x00, 0x0c, 0x00,
		0x47, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	},
}

func TestParse(t *testing.T) {
	for idx, msg := range msgs {
		packet := gopacket.NewPacket(msg, layers.LayerTypeDNS, decodeOptions)
		layer := packet.Layer(layers.LayerTypeDNS)
		if layer == nil {
			t.Errorf("msg-%d: not a DNS packet", idx)
			continue
		}
		fmt.Printf("msg: %v\n", layer)
	}
}
