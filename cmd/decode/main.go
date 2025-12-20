//
// Copyright (c) 2024-2025 Markku Rossi
//
// All rights reserved.
//

package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

var reHex = regexp.MustCompilePOSIX(`^[[:xdigit:]]{8}  ([^|]+)  |`)
var reSep = regexp.MustCompilePOSIX(`[[:space:]]+`)

func main() {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	var digits []string

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		m := reHex.FindStringSubmatch(line)
		if m != nil {
			digits = append(digits, reSep.Split(m[1], -1)...)
		}
	}
	fmt.Printf("digits: %v\n", digits)

	for idx, d := range digits {
		if idx%8 == 0 {
			fmt.Printf(",\n")
		} else {
			fmt.Printf(", ")
		}
		fmt.Printf("0x%s", d)
	}
	fmt.Println()
}
