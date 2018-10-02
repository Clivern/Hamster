package main

import (
    "github.com/gin-gonic/gin"
    "github.com/clivern/hamster/internal/app/controller"
    "net/http"
)

func main() {
    //gin.SetMode(gin.ReleaseMode)
    r := gin.Default()
    r.GET("/", controller.Index)
    r.POST("/listen", controller.Listen)
    r.GET("/favicon.ico", func(c *gin.Context) {
        c.String(http.StatusNoContent, "")
    })

    // Test Routes
    r.GET("/create-comment-test", controller.CreateCommentTest)
    r.GET("/actions-test", controller.ActionsTest)

    r.Run()
}