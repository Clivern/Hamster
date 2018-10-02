package event

import (
    "testing"
    "io/ioutil"
)

func TestCommit(t *testing.T) {

    var commit Commit

    dat, err := ioutil.ReadFile("../../../samples/commit.json")

    if err != nil{
        t.Errorf("File samples/commit.json is invalid!")
    }

    ok, _ := commit.LoadFromJSON(dat)

    if !ok {
        t.Errorf("Testing with file samples/commit.json is invalid")
    }

    got := commit.Commit.Commit.Author.Name
    want := "Clivern"

    if !ok || got != want {
        t.Errorf("got '%s' want '%s'", got, want)
    }
}