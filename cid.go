//
// Copyright (c) 2024 Markku Rossi
//
// All rights reserved.
//

package nap

import (
	"bufio"
	"bytes"
	"strings"
)

// CIDs define known client IDs.
type CIDs map[string]string

func ParseCIDs(data []byte) (CIDs, error) {
	result := make(map[string]string)

	scanner := bufio.NewScanner(bytes.NewReader(data))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		var id, comment string
		sep := strings.IndexByte(line, '#')
		if sep < 0 {
			id = line
			comment = line
		} else {
			id = strings.TrimSpace(line[:sep])
			comment = strings.TrimSpace(line[sep+1:])
		}
		if len(id) == 0 {
			// Comment line.
			continue
		}
		result[id] = comment
	}
	return result, nil
}
