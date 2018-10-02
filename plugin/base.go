package plugin


// Here we define all our custom actions and it will
// get executed once we get a request from github

import (
    "github.com/clivern/hamster/internal/app/event"
    "fmt"
)

// Status Action
func StatusListener(status event.Status)(bool, error){
    fmt.Printf("StatusListener Fired: %s \n", status.Sha)
    return true, nil
}

// Issue Action
func IssuesListener(issues event.Issues)(bool, error){
    fmt.Printf("IssuesListener Fired")
    return true, nil
}

// Issue Comment Action
func IssueCommentListener(issue_comment event.IssueComment)(bool, error){
    fmt.Printf("IssueCommentListener Fired")
    return true, nil
}