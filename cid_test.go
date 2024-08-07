//
// Copyright (c) 2020-2024 Markku Rossi
//
// All rights reserved.
//

package nap

import (
	"fmt"
	"testing"
)

var testCIDData = `# Comment

02da8b774066fd61	# Desktop
f2727277cdbaa234	# Mobile

3de38580d02586b5
7d5a28d772778cee
`

func TestParseCIDs(t *testing.T) {
	cids, err := ParseCIDs([]byte(testCIDData))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("got %v CIDs:\n", len(cids))
	for k, v := range cids {
		fmt.Printf("%v: %v\n", k, v)
	}
}
