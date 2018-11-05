// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	GithubToken           string `json:"github_token"`
	GithubWebhookSecret   string `json:"github_webhook_secret"`
	RepositoryAuthor      string `json:"repository_author"`
	RepositoryName        string `json:"repository_name"`
	AppMode               string `json:"app_mode"`
	AppPort               string `json:"app_port"`
	AppLogLevel           string `json:"app_log_level"`
	AppDomain             string `json:"app_domain"`
	GithubAppClientID     string `json:"github_app_client_id"`
	GithubAppRedirectURI  string `json:"github_app_redirect_uri"`
	GithubAppAllowSignup  string `json:"github_app_allow_signup"`
	GithubAppScope        string `json:"github_app_scope"`
	GithubAppClientSecret string `json:"github_app_client_secret"`
}

func (e *Config) Load(file string) (bool, error) {

	_, err := os.Stat(file)

	if err != nil {
		return false, fmt.Errorf("config file %s not found", file)
	}

	data, err := ioutil.ReadFile(file)

	if err != nil {
		return false, err
	}

	err = json.Unmarshal(data, &e)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (e *Config) Cache() {
	if os.Getenv("AppMode") == "" {
		os.Setenv("GithubToken", e.GithubToken)
		os.Setenv("GithubWebhookSecret", e.GithubWebhookSecret)
		os.Setenv("RepositoryAuthor", e.RepositoryAuthor)
		os.Setenv("RepositoryName", e.RepositoryName)
		os.Setenv("AppMode", e.AppMode)
		os.Setenv("AppLogLevel", e.AppLogLevel)
		os.Setenv("AppPort", e.AppPort)
		os.Setenv("GithubAppClientID", e.GithubAppClientID)
		os.Setenv("GithubAppRedirectURI", e.GithubAppRedirectURI)
		os.Setenv("GithubAppAllowSignup", e.GithubAppAllowSignup)
		os.Setenv("GithubAppScope", e.GithubAppScope)
		os.Setenv("GithubAppClientSecret", e.GithubAppClientSecret)
		os.Setenv("AppDomain", e.AppDomain)
	}
}

func (e *Config) GinEnv() {
	// Used by gin framework
	// https://github.com/gin-gonic/gin/blob/d510595aa58c2417373d89a8d8ffa21cf58673cb/utils.go#L140
	os.Setenv("PORT", os.Getenv("AppPort"))
}
