package event

import (
    "testing"
    "io/ioutil"
)

func TestStatus(t *testing.T) {

    var status Status

    dat, err := ioutil.ReadFile("../../../samples/status.json")

    if err != nil{
        t.Errorf("File samples/status.json is invalid!")
    }

    ok, _ := status.LoadFromJSON(dat)

    if !ok {
        t.Errorf("Testing with file samples/status.json is invalid")
    }

    got := status.Commit.Commit.Author.Name
    want := "Clivern"

    if !ok || got != want {
        t.Errorf("got '%s' want '%s'", got, want)
    }
}