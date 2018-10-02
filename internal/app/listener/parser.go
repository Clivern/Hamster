package listener

import (
    "github.com/clivern/hamster/internal/app/receiver"
    "encoding/json"
)

const OPEN_ISSUE = "OPEN_ISSUE"
const CLOSE_ISSUE = "CLOSE_ISSUE"
const NEW_COMMIT = "NEW_COMMIT"
const NEW_ISSUE_COMMENT = "ISSUE_COMMENT"

type Parser struct {
    Type    string
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

func (e *Parser) GetType (body string) (string, error) {
    e.Type = OPEN_ISSUE
    return e.Type, nil
}

