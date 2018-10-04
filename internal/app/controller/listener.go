package controller

import (
    "github.com/clivern/hamster/internal/app/listener"
    "github.com/clivern/hamster/internal/app/event"
    "github.com/clivern/hamster/plugin"
    "github.com/gin-gonic/gin"
    "github.com/clivern/hamster/pkg"
    "os"
    "fmt"
)

func Listen(c *gin.Context) {
    var actions listener.Action
    rawBody, _ := c.GetRawData()
    body := string(rawBody)

    parser := &listener.Parser{
        UserAgent: c.GetHeader("User-Agent"),
        GithubDelivery: c.GetHeader("X-GitHub-Delivery"),
        GitHubEvent: c.GetHeader("X-GitHub-Event"),
        HubSignature: c.GetHeader("X-Hub-Signature"),
        Body: body,
    }

    ok := parser.VerifySignature(os.Getenv("GithubWebhookSecret"))
    evt := parser.GetGitHubEvent()

    pkg.Info(fmt.Sprintf("Incoming event %s with payload %s!", evt, body))

    if ok {
        switch evt {
        case "status":
            var status event.Status
            status.LoadFromJSON(rawBody)
            actions.RegisterStatusAction(plugin.StatusListener)
            actions.ExecuteStatusActions(status)
        case "watch":
            var watch event.Watch
            watch.LoadFromJSON(rawBody)
            actions.RegisterWatchAction(plugin.WatchListener)
            actions.ExecuteWatchActions(watch)
        case "issues":
            var issues event.Issues
            issues.LoadFromJSON(rawBody)
            actions.RegisterIssuesAction(plugin.IssuesListener)
            actions.ExecuteIssuesActions(issues)
        case "issue_comment":
            var issue_comment event.IssueComment
            issue_comment.LoadFromJSON(rawBody)
            actions.RegisterIssueCommentAction(plugin.IssueCommentListener)
            actions.ExecuteIssueCommentActions(issue_comment)
        default:
            pkg.Info(fmt.Sprintf("Unknown or unsupported event %s!", evt))
        }

        var raw event.Raw
        raw.SetEvent(evt)
        raw.SetBody(body)
        actions.RegisterRawAction(plugin.RawListener)
        actions.ExecuteRawActions(raw)

        c.JSON(200, gin.H{
            "status": "Nice!",
        })
    }else{
        c.JSON(200, gin.H{
            "status": "Oops!",
        })
    }
}