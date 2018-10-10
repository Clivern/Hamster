// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package event

//revive:disable:exported

import (
	"io/ioutil"
	"testing"
)

func TestWatch(t *testing.T) {

	var watch Watch

	dat, err := ioutil.ReadFile("../../../samples/watch.json")

	if err != nil {
		t.Errorf("File samples/watch.json is invalid!")
	}

	ok, _ := watch.LoadFromJSON(dat)

	if !ok {
		t.Errorf("Testing with file samples/watch.json is invalid")
	}

	got := watch.Repository.Owner.Login
	want := "Clivern"

	if !ok || got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
