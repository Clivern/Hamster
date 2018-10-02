package main

import (
    "github.com/gin-gonic/gin"
    "github.com/clivern/hamster/internal/app/controller"
    "net/http"
    "github.com/clivern/hamster/pkg"
    "os"
)

func main() {
    // Load config.json file and store on env
    config := &pkg.Config{}
    config.Load("config.dist.json")
    config.Cache()

    if os.Getenv("AppMode") == "prod" {
        gin.SetMode(gin.ReleaseMode)
    }

    r := gin.Default()
    r.GET("/", controller.Index)
    r.POST("/listen", controller.Listen)
    r.GET("/favicon.ico", func(c *gin.Context) {
        c.String(http.StatusNoContent, "")
    })

    r.Run()
}