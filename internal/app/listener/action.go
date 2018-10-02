package listener

import (
    "github.com/clivern/hamster/internal/app/event"
    "encoding/json"
)

type Action struct {
    Commit          []func(commit event.Commit)(bool, error)
    Issue           []func(issue event.Issue)(bool, error)
    IssueComment    []func(issue_comment event.IssueComment)(bool, error)
}

func (e *Action) LoadFromJSON (data []byte) (bool, error) {
    err := json.Unmarshal(data, &e)
    if err != nil {
        return false, err
    }
    return true, nil
}

func (e *Action) ConvertToJSON () (string, error) {
    data, err := json.Marshal(&e)
    if err != nil {
        return "", err
    }
    return string(data), nil
}

func (e *Action) RegisterCommitAction (f func(commit event.Commit)(bool, error)) {
    e.Commit = append(e.Commit, f)
}

func (e *Action) RegisterIssueAction (f func(issue event.Issue)(bool, error)) {
    e.Issue = append(e.Issue, f)
}

func (e *Action) RegisterIssueCommentAction (f func(issue_comment event.IssueComment)(bool, error)) {
    e.IssueComment = append(e.IssueComment, f)
}

func (e *Action) ExecuteCommitActions (commit event.Commit) (bool, error) {
    for _, fun := range e.Commit{
        ok, err := fun(commit)
        if !ok {
            return false, err
        }
    }
    return true, nil
}

func (e *Action) ExecuteIssueActions (issue event.Issue) (bool, error) {
    for _, fun := range e.Issue{
        ok, err := fun(issue)
        if !ok {
            return false, err
        }
    }
    return true, nil
}

func (e *Action) ExecuteIssueCommentActions (issue_comment event.IssueComment) (bool, error) {
    for _, fun := range e.IssueComment{
        ok, err := fun(issue_comment)
        if !ok {
            return false, err
        }
    }
    return true, nil
}