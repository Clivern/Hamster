package event

import (
    "testing"
    "io/ioutil"
)

func TestIssueComment(t *testing.T) {

    var issue_comment IssueComment

    dat, err := ioutil.ReadFile("../../../samples/issue_comment.json")

    if err != nil{
        t.Errorf("File samples/issue_comment.json is invalid!")
    }

    ok, _ := issue_comment.LoadFromJSON(dat)

    if !ok {
        t.Errorf("Testing with file samples/issue_comment.json is invalid")
    }

    got := issue_comment.Issue.User.Login
    want := "Clivern"

    if !ok || got != want {
        t.Errorf("got '%s' want '%s'", got, want)
    }
}