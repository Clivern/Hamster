package plugin


// Here we define all our custom actions and it will
// get executed once we get a request from github

import (
    "github.com/clivern/hamster/internal/app/receiver"
    "fmt"
)

// Commit Action
func CommitListener(commit receiver.Commit)(bool, error){
    fmt.Printf("CommitListener Fired")
    return true, nil
}

// Issue Action
func IssueListener(issue receiver.Issue)(bool, error){
    fmt.Printf("IssueListener Fired")
    return true, nil
}

// Issue Comment Action
func IssueCommentListener(issue_comment receiver.IssueComment)(bool, error){
    fmt.Printf("IssueCommentListener Fired")
    return true, nil
}