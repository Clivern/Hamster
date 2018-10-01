package receiver

import (
    "testing"
    "io/ioutil"
)

func TestIssue(t *testing.T) {

    var issue Issue
    dat, _ := ioutil.ReadFile("../../../samples/issue.json")

    ok := issue.LoadFromJSON(dat)

    if !ok {
        t.Errorf("Testing file samples/issue.json is invalid")
    }

    got := issue.Issue.User.Login
    want := "Clivern"

    if !ok || got != want {
        t.Errorf("got '%s' want '%s'", got, want)
    }
}