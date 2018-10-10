package main

import (
    "os"
    "io"
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/clivern/hamster/internal/app/controller"
    "github.com/clivern/hamster/pkg"
)

func main() {
    // Load config.json file and store on env
    config := &pkg.Config{}
    config.Load("config.dist.json")
    // This will never override ENV Vars if exists
    config.Cache()
    config.GinEnv()

    if os.Getenv("AppMode") == "prod" {
        gin.SetMode(gin.ReleaseMode)
        gin.DisableConsoleColor()
        f, _ := os.Create("var/logs/gin.log")
        gin.DefaultWriter = io.MultiWriter(f)
    }

    r := gin.Default()
    r.GET("/", controller.Index)
    r.POST("/listen", controller.Listen)
    r.GET("/favicon.ico", func(c *gin.Context) {
        c.String(http.StatusNoContent, "")
    })

    r.Run()
}