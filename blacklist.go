//
// blacklist.go
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
	"strings"
)

// Data assets.
//
//go:embed data/*
var assets embed.FS

// ParseBlacklist parses the blacklist from the data.
func ParseBlacklist(data []byte) ([]Labels, error) {
	var result []Labels

	scanner := bufio.NewScanner(bytes.NewReader(data))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		if line[0] == '#' {
			continue
		}

		result = append(result, strings.Split(line, "."))
	}
	return result, scanner.Err()
}
