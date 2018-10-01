package controller

import (
    "github.com/gin-gonic/gin"
    "fmt"
)

func Listen(c *gin.Context) {

    b, _ := c.GetRawData()

    fmt.Println(string(b))

    c.JSON(200, gin.H{
        "status": "ok",
    })
}