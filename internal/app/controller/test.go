package controller

import (
    "github.com/gin-gonic/gin"
    "github.com/clivern/hamster/pkg"
    "github.com/clivern/hamster/internal/app/event"
    "github.com/clivern/hamster/internal/app/listener"
    "os"
    "fmt"
)

func CreateCommentTest(c *gin.Context) {
    // export GITHUB_TOKEN=b1.....
    github_api := &pkg.GithubAPI{
        Token: os.Getenv("GITHUB_TOKEN"),
        Author:"Clivern",
        Repository:"Hamster",
    }

    created_comment, err := github_api.NewComment("Hi Buddy", 1)

    if err == nil {
        c.JSON(200, gin.H{
            "status": "ok",
            "id": created_comment.ID,
        })
    }else{
        c.JSON(200, gin.H{
            "status": "not ok",
            "error": err.Error(),
        })
    }
}


func ActionsTest(c *gin.Context) {

    var commit event.Commit
    var actions listener.Action

    commit.ID = 1
    commit.Sha = "Hi"

    actions.RegisterCommitAction(func(commit event.Commit)(bool, error){
        fmt.Printf("Action 1 ID: %d\n", commit.ID)
        fmt.Printf("Action 1 SHA: %s\n", commit.Sha)
        return true, nil
    })

    actions.RegisterCommitAction(func(commit event.Commit)(bool, error){
        fmt.Printf("Action 2 ID: %d\n", commit.ID)
        fmt.Printf("Action 3 SHA: %s\n", commit.Sha)
        return true, nil
    })

    actions.ExecuteCommitActions(commit)

    c.JSON(200, gin.H{
        "status": "ok",
        "data": fmt.Sprintf("%s -> %d", commit.Sha, commit.ID),
    })
}


func ParserTest(c *gin.Context) {

    rawBody, _ := c.GetRawData()

    parser := &listener.Parser{
        UserAgent: c.GetHeader("User-Agent"),
        GithubDelivery: c.GetHeader("X-GitHub-Delivery"),
        GitHubEvent: c.GetHeader("X-GitHub-Event"),
        HubSignature: c.GetHeader("X-Hub-Signature"),
        Body: string(rawBody),
    }

    fmt.Println(parser.GetUserAgent())
    fmt.Println(parser.GetGithubDelivery())
    fmt.Println(parser.GetGitHubEvent())
    fmt.Println(parser.GetHubSignature())
    fmt.Println(parser.VerifySignature(os.Getenv("GITHUB_WEBHOOK_SECRET")))

    c.JSON(200, gin.H{
        "status": "ok",
    })
}