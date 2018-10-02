package main

import (
    "github.com/gin-gonic/gin"
    "github.com/clivern/hamster/internal/app/controller"
    "net/http"
    "github.com/clivern/hamster/pkg"
)

func main() {
    // Load config.json file and store on env
    config := &pkg.Config{}
    config.Load("config.dist.json")
    config.Cache()

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
    r.POST("/parser-test", controller.ParserTest)

    r.Run()
}