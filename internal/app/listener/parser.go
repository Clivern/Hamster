// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package listener

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"strings"
)

type Parser struct {
	UserAgent      string
	GithubDelivery string
	GitHubEvent    string
	HubSignature   string
	Headers        map[string]string
	Body           string
}

func (e *Parser) LoadFromJSON(data []byte) (bool, error) {
	err := json.Unmarshal(data, &e)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (e *Parser) ConvertToJSON() (string, error) {
	data, err := json.Marshal(&e)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (e *Parser) SetUserAgent(userAgent string) {
	e.UserAgent = userAgent
}

func (e *Parser) GetUserAgent() string {
	return e.UserAgent
}

func (e *Parser) SetGithubDelivery(githubDelivery string) {
	e.GithubDelivery = githubDelivery
}

func (e *Parser) GetGithubDelivery() string {
	return e.GithubDelivery
}

func (e *Parser) SetGitHubEvent(githubEvent string) {
	e.GitHubEvent = githubEvent
}

func (e *Parser) GetGitHubEvent() string {
	return e.GitHubEvent
}

func (e *Parser) SetHubSignature(hubSignature string) {
	e.HubSignature = hubSignature
}

func (e *Parser) GetHubSignature() string {
	return e.HubSignature
}

func (e *Parser) SetBody(body string) {
	e.Body = body
}

func (e *Parser) GetBody() string {
	return e.Body
}

func (e *Parser) SetHeader(key string, value string) {
	e.Headers[key] = value
}

func (e *Parser) GetHeader(key string) string {
	return e.Headers[key]
}

func (e *Parser) VerifySignature(secret string) bool {

	signature := e.GetHubSignature()
	body := e.GetBody()

	if len(signature) != 45 || !strings.HasPrefix(signature, "sha1=") {
		return false
	}

	actual := make([]byte, 20)
	hex.Decode(actual, []byte(signature[5:]))

	return hmac.Equal(e.SignBody([]byte(secret), []byte(body)), actual)
}

func (e *Parser) SignBody(secret, body []byte) []byte {
	computed := hmac.New(sha1.New, secret)
	computed.Write(body)
	return []byte(computed.Sum(nil))
}
