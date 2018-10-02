package event

import (
    "testing"
    "io/ioutil"
)

func TestIssues(t *testing.T) {

    var issues Issues
    dat, err := ioutil.ReadFile("../../../samples/issues.json")

    if err != nil{
        t.Errorf("File samples/issues.json is invalid!")
    }

    ok, _ := issues.LoadFromJSON(dat)

    if !ok {
        t.Errorf("Testing with file samples/issues.json is invalid")
    }

    got := issues.Issue.User.Login
    want := "Clivern"

    if !ok || got != want {
        t.Errorf("got '%s' want '%s'", got, want)
    }
}