// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"github.com/clivern/hamster/internal/app/pkg/github"
	"github.com/gin-gonic/gin"
	"os"
)

// Auth controller
func Auth(c *gin.Context) {

	githubOauth := &github.OAuthApp{
		ClientID:     os.Getenv("GithubAppClientID"),
		RedirectURI:  os.Getenv("GithubAppRedirectURI"),
		AllowSignup:  os.Getenv("GithubAppAllowSignup"),
		Scope:        os.Getenv("GithubAppScope"),
		ClientSecret: os.Getenv("GithubAppClientSecret"),
	}

	state, err := c.Cookie("gh_oauth_state")

	if err == nil && state != "" {
		githubOauth.SetState(state)
	}

	ok, err := githubOauth.FetchAccessToken(
		c.DefaultQuery("code", ""),
		c.DefaultQuery("state", ""),
	)

	if ok && err == nil {
		c.JSON(200, gin.H{
			"status":      "ok",
			"accessToken": githubOauth.GetAccessToken(),
		})
	} else {
		c.JSON(200, gin.H{
			"status": "not ok",
			"error":  err.Error(),
		})
	}
}
