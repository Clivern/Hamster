// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package listener

import (
	"encoding/json"
	"github.com/clivern/hamster/internal/app/event"
)

// Action struct
type Action struct {
	Status                   []func(status event.Status) (bool, error)
	Issues                   []func(issue event.Issues) (bool, error)
	IssueComment             []func(issueComment event.IssueComment) (bool, error)
	Watch                    []func(watch event.Watch) (bool, error)
	Push                     []func(watch event.Push) (bool, error)
	Create                   []func(watch event.Create) (bool, error)
	Label                    []func(label event.Label) (bool, error)
	Delete                   []func(delete event.Delete) (bool, error)
	Milestone                []func(milestone event.Milestone) (bool, error)
	PullRequest              []func(pullRequest event.PullRequest) (bool, error)
	PullRequestReview        []func(pullRequestReview event.PullRequestReview) (bool, error)
	PullRequestReviewComment []func(pullRequestReviewComment event.PullRequestReviewComment) (bool, error)
	Raw                      []func(raw event.Raw) (bool, error)
}

// LoadFromJSON update object from json
func (e *Action) LoadFromJSON(data []byte) (bool, error) {
	err := json.Unmarshal(data, &e)
	if err != nil {
		return false, err
	}
	return true, nil
}

// ConvertToJSON convert object to json
func (e *Action) ConvertToJSON() (string, error) {
	data, err := json.Marshal(&e)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// RegisterRawAction adds action for all events
func (e *Action) RegisterRawAction(f func(raw event.Raw) (bool, error)) {
	e.Raw = append(e.Raw, f)
}

// RegisterStatusAction adds action for status event
func (e *Action) RegisterStatusAction(f func(status event.Status) (bool, error)) {
	e.Status = append(e.Status, f)
}

// RegisterIssuesAction adds action for issues event
func (e *Action) RegisterIssuesAction(f func(issue event.Issues) (bool, error)) {
	e.Issues = append(e.Issues, f)
}

// RegisterIssueCommentAction adds action for issue comment event
func (e *Action) RegisterIssueCommentAction(f func(issueComment event.IssueComment) (bool, error)) {
	e.IssueComment = append(e.IssueComment, f)
}

// RegisterWatchAction adds action for watch event
func (e *Action) RegisterWatchAction(f func(watch event.Watch) (bool, error)) {
	e.Watch = append(e.Watch, f)
}

// RegisterPushAction adds action for push event
func (e *Action) RegisterPushAction(f func(watch event.Push) (bool, error)) {
	e.Push = append(e.Push, f)
}

// RegisterCreateAction adds action for create event
func (e *Action) RegisterCreateAction(f func(create event.Create) (bool, error)) {
	e.Create = append(e.Create, f)
}

// RegisterLabelAction adds action for label event
func (e *Action) RegisterLabelAction(f func(label event.Label) (bool, error)) {
	e.Label = append(e.Label, f)
}

// RegisterDeleteAction adds action for delete event
func (e *Action) RegisterDeleteAction(f func(delete event.Delete) (bool, error)) {
	e.Delete = append(e.Delete, f)
}

// RegisterMilestoneAction adds action for milestone event
func (e *Action) RegisterMilestoneAction(f func(milestone event.Milestone) (bool, error)) {
	e.Milestone = append(e.Milestone, f)
}

// RegisterPullRequestAction adds action for pull request event
func (e *Action) RegisterPullRequestAction(f func(pullRequest event.PullRequest) (bool, error)) {
	e.PullRequest = append(e.PullRequest, f)
}

// RegisterPullRequestReviewAction adds action for pull request review event
func (e *Action) RegisterPullRequestReviewAction(f func(pullRequestReview event.PullRequestReview) (bool, error)) {
	e.PullRequestReview = append(e.PullRequestReview, f)
}

// RegisterPullRequestReviewCommentAction adds action for pull request review comment event
func (e *Action) RegisterPullRequestReviewCommentAction(f func(pullRequestReviewComment event.PullRequestReviewComment) (bool, error)) {
	e.PullRequestReviewComment = append(e.PullRequestReviewComment, f)
}

// ExecuteRawActions executes actions for all events
func (e *Action) ExecuteRawActions(raw event.Raw) (bool, error) {
	for _, fun := range e.Raw {
		ok, err := fun(raw)
		if !ok {
			return false, err
		}
	}
	return true, nil
}

// ExecuteStatusActions executes actions for status events
func (e *Action) ExecuteStatusActions(status event.Status) (bool, error) {
	for _, fun := range e.Status {
		ok, err := fun(status)
		if !ok {
			return false, err
		}
	}
	return true, nil
}

// ExecuteIssuesActions executes actions for issues events
func (e *Action) ExecuteIssuesActions(issue event.Issues) (bool, error) {
	for _, fun := range e.Issues {
		ok, err := fun(issue)
		if !ok {
			return false, err
		}
	}
	return true, nil
}

// ExecuteIssueCommentActions executes actions for issue comment events
func (e *Action) ExecuteIssueCommentActions(issueComment event.IssueComment) (bool, error) {
	for _, fun := range e.IssueComment {
		ok, err := fun(issueComment)
		if !ok {
			return false, err
		}
	}
	return true, nil
}

// ExecuteWatchActions executes actions for watch events
func (e *Action) ExecuteWatchActions(watch event.Watch) (bool, error) {
	for _, fun := range e.Watch {
		ok, err := fun(watch)
		if !ok {
			return false, err
		}
	}
	return true, nil
}

// ExecutePushActions executes actions for push events
func (e *Action) ExecutePushActions(push event.Push) (bool, error) {
	for _, fun := range e.Push {
		ok, err := fun(push)
		if !ok {
			return false, err
		}
	}
	return true, nil
}

// ExecuteCreateActions executes actions for create events
func (e *Action) ExecuteCreateActions(create event.Create) (bool, error) {
	for _, fun := range e.Create {
		ok, err := fun(create)
		if !ok {
			return false, err
		}
	}
	return true, nil
}

// ExecuteLabelActions executes actions for label events
func (e *Action) ExecuteLabelActions(label event.Label) (bool, error) {
	for _, fun := range e.Label {
		ok, err := fun(label)
		if !ok {
			return false, err
		}
	}
	return true, nil
}

// ExecuteDeleteActions executes actions for delete events
func (e *Action) ExecuteDeleteActions(delete event.Delete) (bool, error) {
	for _, fun := range e.Delete {
		ok, err := fun(delete)
		if !ok {
			return false, err
		}
	}
	return true, nil
}

// ExecuteMilestoneActions executes actions for milestone events
func (e *Action) ExecuteMilestoneActions(milestone event.Milestone) (bool, error) {
	for _, fun := range e.Milestone {
		ok, err := fun(milestone)
		if !ok {
			return false, err
		}
	}
	return true, nil
}

// ExecutePullRequestActions executes actions for pull request events
func (e *Action) ExecutePullRequestActions(pullRequest event.PullRequest) (bool, error) {
	for _, fun := range e.PullRequest {
		ok, err := fun(pullRequest)
		if !ok {
			return false, err
		}
	}
	return true, nil
}

// ExecutePullRequestReviewActions executes actions for pull requests review events
func (e *Action) ExecutePullRequestReviewActions(pullRequestReview event.PullRequestReview) (bool, error) {
	for _, fun := range e.PullRequestReview {
		ok, err := fun(pullRequestReview)
		if !ok {
			return false, err
		}
	}
	return true, nil
}

// ExecutePullRequestReviewCommentActions executes actions for pull requests review comment events
func (e *Action) ExecutePullRequestReviewCommentActions(pullRequestReviewComment event.PullRequestReviewComment) (bool, error) {
	for _, fun := range e.PullRequestReviewComment {
		ok, err := fun(pullRequestReviewComment)
		if !ok {
			return false, err
		}
	}
	return true, nil
}
