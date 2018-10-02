package controller

import (
    "github.com/clivern/hamster/internal/app/listener"
    "github.com/clivern/hamster/internal/app/event"
    "github.com/clivern/hamster/plugin"
    "github.com/gin-gonic/gin"
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

    if ok {
        switch evt := parser.GetGitHubEvent(); evt {
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
        default:
            fmt.Printf("Unknown or Unsupported Event %s", evt)
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