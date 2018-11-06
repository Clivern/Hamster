// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"github.com/clivern/hamster/internal/app/pkg/github"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

// Login controller
func Login(c *gin.Context) {

	githubOauth := &github.OAuthApp{
		ClientID:     os.Getenv("GithubAppClientID"),
		RedirectURI:  os.Getenv("GithubAppRedirectURI"),
		AllowSignup:  os.Getenv("GithubAppAllowSignup"),
		Scope:        os.Getenv("GithubAppScope"),
		ClientSecret: os.Getenv("GithubAppClientSecret"),
	}

	state, err := c.Cookie("gh_oauth_state")

	if err != nil || state == "" {
		githubOauth.GenerateState()
		c.SetCookie("gh_oauth_state", githubOauth.GetState(), 3600, "/", os.Getenv("AppDomain"), true, true)
	} else {
		githubOauth.SetState(state)
	}

	c.HTML(http.StatusOK, "login.tmpl", gin.H{
		"title":    "Hamster",
		"oauthURL": githubOauth.BuildAuthorizeURL(),
	})
}
