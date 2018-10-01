package receiver

import (
    "testing"
    "io/ioutil"
)

func TestIssueComment(t *testing.T) {

    var issue_comment IssueComment
    dat, _ := ioutil.ReadFile("../../../samples/issue_comment.json")

    ok := issue_comment.LoadFromJSON(dat)

    if !ok {
        t.Errorf("Testing file samples/issue_comment.json is invalid")
    }

    got := issue_comment.Issue.User.Login
    want := "Clivern"

    if !ok || got != want {
        t.Errorf("got '%s' want '%s'", got, want)
    }
}