package listener

import (
    "regexp"
    "github.com/clivern/hamster/internal/app/event"
)

type Commands struct{
    Incoming            []event.Command
    Issues              map[string]func(command event.Command, issue event.Issues)(bool, error)
    IssueComment        map[string]func(command event.Command, issue_comment event.IssueComment)(bool, error)
}

// This will fetch all commands and parameters within the issue or issue comment
// /fire become fire & []
// /run{test,cases} become run & [test,cases]
func (e *Commands) Fetch(body string) {
    re := regexp.MustCompile(`\S*(/[a-zA-Z0-9])\S*`)
    re.MatchString(body)

    submatchall := re.FindAllString(body,-1)
    for _, element := range submatchall {
        command := event.Command{Data: element}
        command.Parse()
        e.Incoming = append(e.Incoming, command)
    }
}
