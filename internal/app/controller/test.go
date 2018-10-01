package controller

import (
    "github.com/gin-gonic/gin"
    "github.com/clivern/hamster/pkg"
    "os"
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