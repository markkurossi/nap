//
// Copyright (c) 2019-2024 Markku Rossi
//
// All rights reserved.
//

package nap

import (
	"bufio"
	"bytes"
	"embed"
	"fmt"
	"net"
	"strings"
)

// Data assets.
//
//go:embed data/*
var assets embed.FS

// Blacklist implements a blacklist entry.
type Blacklist struct {
	Labels  Labels
	Name    string
	Address net.IP
}

// Block tests if the blacklist entry is a block entry.
func (b Blacklist) Block() bool {
	return len(b.Name) == 0 && b.Address == nil
}

func (b Blacklist) String() string {
	if len(b.Name) > 0 {
		return fmt.Sprintf("%s => %s", b.Labels, b.Name)
	}
	if b.Address != nil {
		return fmt.Sprintf("%s => %s", b.Labels, b.Address)
	}
	return b.Labels.String()
}

// ParseBlacklist parses the blacklist from the data.
func ParseBlacklist(data []byte) ([]Blacklist, error) {
	var result []Blacklist

	scanner := bufio.NewScanner(bytes.NewReader(data))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		if line[0] == '#' {
			continue
		}

		parts := strings.Split(line, "=>")
		switch len(parts) {
		case 1:
			result = append(result, Blacklist{
				Labels: strings.Split(parts[0], "."),
			})

		case 2:
			bl := Blacklist{
				Labels: strings.Split(strings.TrimSpace(parts[0]), "."),
			}
			ip := net.ParseIP(parts[1])
			if ip != nil {
				bl.Address = ip
			} else {
				bl.Name = parts[1]
			}
			result = append(result, bl)
		default:
			return nil, fmt.Errorf("malformed line: %s", line)
		}
	}
	return result, scanner.Err()
}
