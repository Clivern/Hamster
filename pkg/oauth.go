// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package pkg

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
)

const (
	GithubOAuthURL   = "https://github.com/login/oauth/authorize"
	OAuthAccessToken = "https://github.com/login/oauth/access_token"
)

// OAuth is a Representation of a Github OAuth API
type GithubOAuth struct {
	ClientID     string   `json:"client_id"`
	RedirectURI  string   `json:"redirect_uri"`
	Scope        string   `json:"scope"`
	Scopes       []string `json:"scopes"`
	State        string   `json:"state"`
	AllowSignup  string   `json:"allow_signup"`
	ClientSecret string   `json:"client_secret"`
	Code         string   `json:"code"`
	AccessToken  string   `json:"access_token"`
	TokenType    string   `json:"token_type"`
}

func (e *GithubOAuth) LoadFromJSON(data []byte) (bool, error) {
	err := json.Unmarshal(data, &e)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (e *GithubOAuth) ConvertToJSON() (string, error) {
	data, err := json.Marshal(&e)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (e *GithubOAuth) GenerateState() {
	val, err := e.RandomString(20)
	if err == nil {
		e.State = val
	}
}

func (e *GithubOAuth) GetState() string {
	return e.State
}

func (e *GithubOAuth) AddScope(scope string) {
	e.Scopes = append(e.Scopes, scope)
}

func (e *GithubOAuth) AddScopes(scopes []string) {
	e.Scopes = scopes
}

// Build github authorization url to send user to
func (e *GithubOAuth) BuildAuthorizeURL() string {
	e.Scope = strings.Join(e.Scopes, ",")
	return fmt.Sprintf(
		"%s?client_id=%s&redirect_uri=%s&scope=%s&state=%s&allow_signup=%s",
		GithubOAuthURL,
		e.ClientID,
		e.RedirectURI,
		e.Scope,
		e.State,
		e.AllowSignup,
	)
}

// Verify state and fetch the access token with incoming code
func (e *GithubOAuth) GetAccessToken(code string, state string) (string, error) {
	return "", nil
}

// Generate a random string
func (e *GithubOAuth) RandomString(len int) (string, error) {
	bytes := make([]byte, len)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
