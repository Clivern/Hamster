package event

import (
    "testing"
    "io/ioutil"
)

func TestIssue(t *testing.T) {

    var issue Issue
    dat, err := ioutil.ReadFile("../../../samples/issue.json")

    if err != nil{
        t.Errorf("File samples/issue.json is invalid!")
    }

    ok, _ := issue.LoadFromJSON(dat)

    if !ok {
        t.Errorf("Testing with file samples/issue.json is invalid")
    }

    got := issue.Issue.User.Login
    want := "Clivern"

    if !ok || got != want {
        t.Errorf("got '%s' want '%s'", got, want)
    }
}