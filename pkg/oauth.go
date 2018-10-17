// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package pkg

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	GithubOAuthURL   = "https://github.com/login/oauth/authorize"
	OAuthAccessToken = "https://github.com/login/oauth/access_token"
)

type GithubOAuthApp struct {
	ClientID     string   `json:"client_id"`
	RedirectURI  string   `json:"redirect_uri"`
	Scope        string   `json:"scope"`
	Scopes       []string `json:"scopes"`
	State        string   `json:"state"`
	AllowSignup  string   `json:"allow_signup"`
	ClientSecret string   `json:"client_secret"`
	AccessToken  string   `json:"access_token"`
	TokenType    string   `json:"token_type"`
}

type GithubOAuthClient struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

type GithubAccessToken struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Code         string `json:"code"`
	State        string `json:"state"`
}

func (e *GithubOAuthApp) GenerateState() {
	val, err := e.RandomString(20)
	if err == nil {
		e.State = val
	}
}

func (e *GithubOAuthApp) GetState() string {
	return e.State
}

func (e *GithubOAuthApp) SetState(state string) {
	e.State = state
}

func (e *GithubOAuthApp) AddScope(scope string) {
	e.Scopes = append(e.Scopes, scope)
	e.Scope = strings.Join(e.Scopes, ",")
}

func (e *GithubOAuthApp) AddScopes(scopes []string) {
	e.Scopes = scopes
	e.Scope = strings.Join(e.Scopes, ",")
}

func (e *GithubOAuthApp) BuildAuthorizeURL() string {
	e.Scope = strings.Join(e.Scopes, ",")

	u, err := url.Parse(GithubOAuthURL)

	if err != nil {
		return ""
	}

	q := u.Query()
	q.Set("client_id", e.ClientID)
	q.Set("redirect_uri", e.RedirectURI)
	q.Set("scope", e.Scope)
	q.Set("state", e.State)
	q.Set("allow_signup", e.AllowSignup)
	u.RawQuery = q.Encode()

	return u.String()
}

func (e *GithubOAuthApp) RandomString(len int) (string, error) {
	bytes := make([]byte, len)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func (e *GithubOAuthApp) FetchAccessToken(code string, state string) (bool, error) {

	accessTokenRequest := &GithubAccessToken{
		ClientID:     e.ClientID,
		ClientSecret: e.ClientSecret,
		Code:         code,
		State:        e.State,
	}

	jsonBody, err := accessTokenRequest.ConvertToJSON()

	if err != nil {
		return false, err
	}

	githubOAuthClient := &GithubOAuthClient{}

	if state != e.State {
		return false, fmt.Errorf(
			"Invalid state provided %s, original one is %s",
			state,
			e.State,
		)
	}

	client := &http.Client{}
	req, err := http.NewRequest(
		"POST",
		OAuthAccessToken,
		bytes.NewBufferString(jsonBody),
	)

	if err != nil {
		return false, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		return false, err
	}

	defer resp.Body.Close()

	bodyByte, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return false, err
	}

	err = json.Unmarshal(bodyByte, &githubOAuthClient)

	if err != nil {
		return false, err
	}

	e.AccessToken = githubOAuthClient.AccessToken
	e.TokenType = githubOAuthClient.TokenType

	return true, nil
}

func (e *GithubOAuthApp) GetAccessToken() string {
	return e.AccessToken
}

func (e *GithubOAuthApp) GetTokenType() string {
	return e.TokenType
}

func (e *GithubAccessToken) ConvertToJSON() (string, error) {
	data, err := json.Marshal(&e)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
