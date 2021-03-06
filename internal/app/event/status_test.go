// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package event

import (
	"github.com/nbio/st"
	"io/ioutil"
	"testing"
)

// TestStatus test cases
func TestStatus(t *testing.T) {

	var status Status

	dat, err := ioutil.ReadFile("../../../fixtures/status.json")

	if err != nil {
		t.Errorf("File fixtures/status.json is invalid!")
	}

	ok, _ := status.LoadFromJSON(dat)

	if !ok {
		t.Errorf("Testing with file fixtures/status.json is invalid")
	}

	st.Expect(t, status.Commit.Commit.Author.Name, "Clivern")
}
