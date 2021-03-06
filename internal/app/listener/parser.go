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

// Parser struct
type Parser struct {
	UserAgent      string
	GithubDelivery string
	GitHubEvent    string
	HubSignature   string
	Headers        map[string]string
	Body           string
}

// LoadFromJSON update object from json
func (e *Parser) LoadFromJSON(data []byte) (bool, error) {
	err := json.Unmarshal(data, &e)
	if err != nil {
		return false, err
	}
	return true, nil
}

// ConvertToJSON convert object to json
func (e *Parser) ConvertToJSON() (string, error) {
	data, err := json.Marshal(&e)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// SetUserAgent sets user agent
func (e *Parser) SetUserAgent(userAgent string) {
	e.UserAgent = userAgent
}

// GetUserAgent gets user agent
func (e *Parser) GetUserAgent() string {
	return e.UserAgent
}

// SetGithubDelivery sets github delivery
func (e *Parser) SetGithubDelivery(githubDelivery string) {
	e.GithubDelivery = githubDelivery
}

// GetGithubDelivery gets github delivery
func (e *Parser) GetGithubDelivery() string {
	return e.GithubDelivery
}

// SetGitHubEvent sets github event
func (e *Parser) SetGitHubEvent(githubEvent string) {
	e.GitHubEvent = githubEvent
}

// GetGitHubEvent gets github event
func (e *Parser) GetGitHubEvent() string {
	return e.GitHubEvent
}

// SetHubSignature sets hub signature
func (e *Parser) SetHubSignature(hubSignature string) {
	e.HubSignature = hubSignature
}

// GetHubSignature gets hub signature
func (e *Parser) GetHubSignature() string {
	return e.HubSignature
}

// SetBody sets body
func (e *Parser) SetBody(body string) {
	e.Body = body
}

// GetBody gets body
func (e *Parser) GetBody() string {
	return e.Body
}

// SetHeader sets header
func (e *Parser) SetHeader(key string, value string) {
	e.Headers[key] = value
}

// GetHeader gets header
func (e *Parser) GetHeader(key string) string {
	return e.Headers[key]
}

// VerifySignature verify signature
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

// SignBody signs body
func (e *Parser) SignBody(secret, body []byte) []byte {
	computed := hmac.New(sha1.New, secret)
	computed.Write(body)
	return []byte(computed.Sum(nil))
}
