// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
    "net/http"
	"github.com/clivern/hamster/pkg"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	githubOauth := &pkg.GithubOAuthApp {
        ClientID: "ClientID",
        RedirectURI: "RedirectURI",
        Scope: "Scope",
        State: "State",
        AllowSignup: "AllowSignup",
    }
	githubOauth.AddScope("scope1")
	githubOauth.AddScope("scope2")
	githubOauth.AddScope("scope3")
	githubOauth.GenerateState()

    c.HTML(http.StatusOK, "login.tmpl", gin.H{
        "title": "Hamster",
        "oauthURL": githubOauth.BuildAuthorizeURL(),
    })
}
