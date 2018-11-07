// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package event

import (
	"github.com/nbio/st"
	"io/ioutil"
	"testing"
)

// TestIssues test cases
func TestIssues(t *testing.T) {

	var issues Issues
	dat, err := ioutil.ReadFile("../../../fixtures/issues.json")

	if err != nil {
		t.Errorf("File fixtures/issues.json is invalid!")
	}

	ok, _ := issues.LoadFromJSON(dat)

	if !ok {
		t.Errorf("Testing with file fixtures/issues.json is invalid")
	}

	st.Expect(t, issues.Issue.User.Login, "Clivern")
}
