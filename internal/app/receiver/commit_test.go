package receiver

import (
    "testing"
    "io/ioutil"
)

func TestCommit(t *testing.T) {

    var commit Commit
    dat, _ := ioutil.ReadFile("../../../samples/commit.json")

    ok := commit.LoadFromJSON(dat)

    if !ok {
        t.Errorf("Testing file samples/commit.json is invalid")
    }

    got := commit.Commit.Commit.Author.Name
    want := "Clivern"

    if !ok || got != want {
        t.Errorf("got '%s' want '%s'", got, want)
    }
}