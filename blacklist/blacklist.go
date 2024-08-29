//
// Copyright (c) 2019-2024 Markku Rossi
//
// All rights reserved.
//

package blacklist

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"regexp"
	"strings"
)

var reSpace = regexp.MustCompilePOSIX("[[:space:]]+")

// Blacklist defines a blacklist.
type Blacklist struct {
	Entries []*Entry
}

// Match matches the name to blacklist entries and returns the match
// or nil.
func (bl *Blacklist) Match(name string) *Entry {
	labels := NewLabels(name)
	for _, entry := range bl.Entries {
		if labels.Match(entry.Labels) {
			return entry
		}
	}
	return nil
}

// Entry implements a blacklist entry.
type Entry struct {
	Labels   Labels
	Name     string
	Address  net.IP
	ProxyCmd ProxyCmd
}

// ProxyCmd defines the actions from proxied URLs.
type ProxyCmd byte

// Proxy actions.
const (
	ProxyBlock ProxyCmd = iota
	ProxyVAST
	ProxyGooglebot
)

func (cmd ProxyCmd) String() string {
	switch cmd {
	case ProxyVAST:
		return "VAST"
	case ProxyGooglebot:
		return "Googlebot"
	default:
		return "block"
	}
}

// Block tests if the blacklist entry is a block entry.
func (b Entry) Block() bool {
	return len(b.Name) == 0 && b.Address == nil
}

func (b Entry) String() string {
	if len(b.Name) > 0 {
		return fmt.Sprintf("%s => %s %s", b.Labels, b.Name, b.ProxyCmd)
	}
	if b.Address != nil {
		return fmt.Sprintf("%s => %s %s", b.Labels, b.Address, b.ProxyCmd)
	}
	return b.Labels.String()
}

// ParseData parses the blacklist from the data.
func ParseData(data []byte) (*Blacklist, error) {
	result := new(Blacklist)

	scanner := bufio.NewScanner(bytes.NewReader(data))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		if line[0] == '#' {
			continue
		}

		parts := reSpace.Split(line, -1)

		switch len(parts) {
		case 1:
			result.Entries = append(result.Entries, &Entry{
				Labels: strings.Split(parts[0], "."),
			})

		case 3, 4:
			if parts[1] != "=>" {
				return nil, fmt.Errorf("malformed line: %s", line)
			}
			bl := &Entry{
				Labels: strings.Split(parts[0], "."),
			}
			ip := net.ParseIP(parts[2])
			if ip != nil {
				bl.Address = ip
			} else {
				bl.Name = parts[2]
			}
			if len(parts) == 4 {
				switch parts[3] {
				case "VAST":
					bl.ProxyCmd = ProxyVAST

				case "Googlebot":
					bl.ProxyCmd = ProxyGooglebot

				default:
					return nil, fmt.Errorf("unknown proxy command: %s",
						parts[3])
				}
			}
			result.Entries = append(result.Entries, bl)

		default:
			return nil, fmt.Errorf("malformed line: %s", line)
		}
	}
	return result, scanner.Err()
}
