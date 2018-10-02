package controller

import (
    "github.com/clivern/hamster/internal/app/listener"
    "github.com/clivern/hamster/internal/app/event"
    "github.com/clivern/hamster/plugin"
    "github.com/gin-gonic/gin"
    "os"
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

    if ok {
        if parser.GetGitHubEvent() == "status" {
            var commit event.Commit
            commit.LoadFromJSON(rawBody)
            actions.RegisterCommitAction(plugin.CommitListener)
            actions.ExecuteCommitActions(commit)
        }
        c.JSON(200, gin.H{
            "status": "Nice!",
        })
    }else{
        c.JSON(200, gin.H{
            "status": "Oops!",
        })
    }
}