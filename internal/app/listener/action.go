package listener

import (
    "github.com/clivern/hamster/internal/app/event"
    "encoding/json"
)

type Action struct {
    Status              []func(status event.Status)(bool, error)
    Issues              []func(issue event.Issues)(bool, error)
    IssueComment        []func(issue_comment event.IssueComment)(bool, error)
    Watch               []func(watch event.Watch)(bool, error)
    Raw                 []func(raw event.Raw)(bool, error)
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

func (e *Action) RegisterRawAction (f func(raw event.Raw)(bool, error)) {
    e.Raw = append(e.Raw, f)
}

func (e *Action) RegisterStatusAction (f func(status event.Status)(bool, error)) {
    e.Status = append(e.Status, f)
}

func (e *Action) RegisterIssuesAction (f func(issue event.Issues)(bool, error)) {
    e.Issues = append(e.Issues, f)
}

func (e *Action) RegisterIssueCommentAction (f func(issue_comment event.IssueComment)(bool, error)) {
    e.IssueComment = append(e.IssueComment, f)
}

func (e *Action) RegisterWatchAction (f func(watch event.Watch)(bool, error)) {
    e.Watch = append(e.Watch, f)
}

func (e *Action) ExecuteRawActions (raw event.Raw) (bool, error) {
    for _, fun := range e.Raw{
        ok, err := fun(raw)
        if !ok {
            return false, err
        }
    }
    return true, nil
}

func (e *Action) ExecuteStatusActions (status event.Status) (bool, error) {
    for _, fun := range e.Status{
        ok, err := fun(status)
        if !ok {
            return false, err
        }
    }
    return true, nil
}

func (e *Action) ExecuteIssuesActions (issue event.Issues) (bool, error) {
    for _, fun := range e.Issues{
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

func (e *Action) ExecuteWatchActions (watch event.Watch) (bool, error) {
    for _, fun := range e.Watch{
        ok, err := fun(watch)
        if !ok {
            return false, err
        }
    }
    return true, nil
}