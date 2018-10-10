// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package listener

//revive:disable:exported

import (
	"github.com/clivern/hamster/internal/app/event"
	"regexp"
)

type Commands struct {
	Incoming     []event.Command
	Issues       map[string]func(command event.Command, issues event.Issues) (bool, error)
	IssueComment map[string]func(command event.Command, issue_comment event.IssueComment) (bool, error)
}

// This will fetch all commands and parameters within the issue or issue comment
// /fire become fire & []
// /run{test,cases} become run & [test,cases]
func (e *Commands) Fetch(body string) {
	re := regexp.MustCompile(`\S*(/[a-zA-Z0-9])\S*`)
	re.MatchString(body)

	submatchall := re.FindAllString(body, -1)
	for _, element := range submatchall {
		command := event.Command{Data: element}
		command.Parse()
		e.Incoming = append(e.Incoming, command)
	}
}

func (e *Commands) RegisterIssuesAction(command string, callback func(command event.Command, issues event.Issues) (bool, error)) {
	if e.Issues == nil {
		e.Issues = make(map[string]func(command event.Command, issues event.Issues) (bool, error))
	}
	e.Issues[command] = callback
}

func (e *Commands) RegisterIssueCommentAction(command string, callback func(command event.Command, issue_comment event.IssueComment) (bool, error)) {
	if e.IssueComment == nil {
		e.IssueComment = make(map[string]func(command event.Command, issue_comment event.IssueComment) (bool, error))
	}
	e.IssueComment[command] = callback
}

func (e *Commands) ExecuteIssuesActions(issues event.Issues) (bool, error) {
	e.Fetch(issues.Issue.Body)
	for _, command := range e.Incoming {
		if fun, ok := e.Issues[command.Name]; ok {
			ok, err := fun(command, issues)
			if !ok {
				return false, err
			}
		}

	}
	return true, nil
}

func (e *Commands) ExecuteIssueCommentActions(issue_comment event.IssueComment) (bool, error) {
	e.Fetch(issue_comment.Comment.Body)
	for _, command := range e.Incoming {
		if fun, ok := e.IssueComment[command.Name]; ok {
			ok, err := fun(command, issue_comment)
			if !ok {
				return false, err
			}
		}

	}
	return true, nil
}
