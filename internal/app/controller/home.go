package controller

import (
    "github.com/gin-gonic/gin"
    "github.com/clivern/hamster/pkg"
    "os"
    "fmt"
)

func Index(c *gin.Context) {

    github_api := &pkg.GithubAPI{
        Token: os.Getenv("GithubToken"),
        Author: os.Getenv("RepositoryAuthor"),
        Repository: os.Getenv("RepositoryName"),
    }

    label, _ := github_api.GetLabel("Release-2.0.0")

    fmt.Println(label)

    c.JSON(200, gin.H{
        "status": "ok",
    })
}