package listener

import (
    "github.com/clivern/hamster/internal/app/receiver"
    "encoding/json"
)

type Action struct {
    Commit          []func(commit receiver.Commit)(bool, error)
    Issue           []func(issue receiver.Issue)(bool, error)
    IssueComment    []func(issue_comment receiver.IssueComment)(bool, error)
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

func (e *Action) RegisterCommitAction (f func(commit receiver.Commit)(bool, error)) (bool, error) {
    e.Commit = append(e.Commit, f)
    return true, nil
}

func (e *Action) RegisterIssueAction (f func(issue receiver.Issue)(bool, error)) (bool, error) {
    e.Issue = append(e.Issue, f)
    return true, nil
}

func (e *Action) RegisterIssueCommentAction (f func(issue_comment receiver.IssueComment)(bool, error)) (bool, error) {
    e.IssueComment = append(e.IssueComment, f)
    return true, nil
}

func (e *Action) ExecuteCommitActions (commit receiver.Commit) (bool, error) {
    for _, fun := range e.Commit{
        ok, err := fun(commit)
        if !ok {
            return false, err
        }
    }
    return true, nil
}

func (e *Action) ExecuteIssueActions (issue receiver.Issue) (bool, error) {
    for _, fun := range e.Issue{
        ok, err := fun(issue)
        if !ok {
            return false, err
        }
    }
    return true, nil
}

func (e *Action) ExecuteIssueCommentActions (issue_comment receiver.IssueComment) (bool, error) {
    for _, fun := range e.IssueComment{
        ok, err := fun(issue_comment)
        if !ok {
            return false, err
        }
    }
    return true, nil
}