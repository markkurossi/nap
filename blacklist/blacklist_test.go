//
// Copyright (c) 2019-2024 Markku Rossi
//
// All rights reserved.
//

package blacklist

import (
	"fmt"
	"testing"
)

var blacklistData = `# Test blacklist

*.more-adds-please.com
**.fwmrm.net => bing.com
**.fwmrm.net => google.com VAST
**.addthis.com => 10.1.9.49 Googlebot
`

func TestBlacklist(t *testing.T) {
	bl, err := ParseData([]byte(blacklistData))
	if err != nil {
		t.Fatal(err)
	}
	const count = 4
	if len(bl.Entries) != count {
		t.Errorf("expected %d items, got %v", count, len(bl.Entries))
	}
	for idx, b := range bl.Entries {
		fmt.Printf("%d: %s\n", idx, b)
	}
}
