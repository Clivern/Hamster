package controller

import (
    "github.com/clivern/hamster/internal/app/receiver"
    "github.com/gin-gonic/gin"
)

func Listen(c *gin.Context) {

    var commit receiver.Commit

    b, _ := c.GetRawData()

    ok, _ := commit.LoadFromJSON(b)

    if !ok {
        c.JSON(200, gin.H{
            "status": "not ok",
        })
    }else{
        c.JSON(200, gin.H{
            "status": "ok",
            "author": commit.Commit.Commit.Author.Name,
        })
    }
}