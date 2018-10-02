package listener

import (
    _"github.com/clivern/hamster/internal/app/event"
    "encoding/json"
    "crypto/hmac"
    "crypto/sha1"
    "encoding/hex"
    "strings"
)

type Parser struct {
    UserAgent           string
    GithubDelivery      string
    GitHubEvent         string
    HubSignature        string
    Headers             map[string]string
    Body                string
}


func (e *Parser) LoadFromJSON (data []byte) (bool, error) {
    err := json.Unmarshal(data, &e)
    if err != nil {
        return false, err
    }
    return true, nil
}

func (e *Parser) ConvertToJSON () (string, error) {
    data, err := json.Marshal(&e)
    if err != nil {
        return "", err
    }
    return string(data), nil
}

func (e *Parser) SetUserAgent (user_agent string) {
    e.UserAgent = user_agent
}

func (e *Parser) GetUserAgent () (string) {
    return e.UserAgent
}

func (e *Parser) SetGithubDelivery (github_delivery string) {
    e.GithubDelivery = github_delivery
}

func (e *Parser) GetGithubDelivery () (string) {
    return e.GithubDelivery
}

func (e *Parser) SetGitHubEvent (github_event string) {
    e.GitHubEvent = github_event
}

func (e *Parser) GetGitHubEvent () (string) {
    return e.GitHubEvent
}

func (e *Parser) SetHubSignature (hub_signature string) () {
    e.HubSignature = hub_signature
}

func (e *Parser) GetHubSignature () (string) {
    return e.HubSignature
}

func (e *Parser) SetBody (body string) {
    e.Body = body
}

func (e *Parser) GetBody () (string) {
    return e.Body
}

func (e *Parser) SetHeader (key string, value string) {
    e.Headers[key] = value
}

func (e *Parser) GetHeader (key string) (string) {
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

func (e *Parser) Parse () (bool, error) {
    // Define the incoming github action

    return true, nil
}