//
// Copyright (c) 2019-2024 Markku Rossi
//
// All rights reserved.
//

package nap

import (
	"testing"
)

var blacklistData = `# Test blacklist

*.more-adds-please.com
**.fwmrm.net => google.com
**.addthis.com => 10.1.9.49
`

func TestBlacklist(t *testing.T) {
	bl, err := ParseBlacklist([]byte(blacklistData))
	if err != nil {
		t.Fatal(err)
	}
	if len(bl) != 3 {
		t.Errorf("expected 3 items, got %v", len(bl))
	}
}
