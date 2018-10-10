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

    labels, err := github_api.GetRepositoryIssueLabels(11)

    if err == nil {
        fmt.Println(labels)
    }else{
        fmt.Println(err.Error())
    }

    c.JSON(200, gin.H{
        "status": "ok",
    })
}