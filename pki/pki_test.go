//
// Copyright (c) 2020-2024 Markku Rossi
//
// All rights reserved.
//

package pki

import (
	"testing"
)

func TestCreateEEKey(t *testing.T) {
	ca := &CA{}
	priv, pub, err := ca.CreateEEKey()
	if err != nil {
		t.Error(err)
	}
	_ = priv
	_ = pub
}
