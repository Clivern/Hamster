// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package main

import (
	"github.com/clivern/hamster/internal/app/controller"
	"github.com/clivern/hamster/internal/app/pkg/utils"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

func main() {
	// Load config.json file and store on env
	config := &utils.Config{}
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
	r.Static("/assets", "./assets")
	r.LoadHTMLGlob("templates/*")
	r.GET("/", controller.Index)
	r.GET("/_healthcheck", controller.HealthCheck)
	r.GET("/login", controller.Login)
	r.GET("/auth", controller.Auth)
	r.POST("/listen", controller.Listen)
	r.GET("/favicon.ico", func(c *gin.Context) {
		c.String(http.StatusNoContent, "")
	})

	r.Run()
}
