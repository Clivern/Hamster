package event

import (
    "encoding/json"
)

// Any time a team is created, deleted, modified, or added to or removed from a repository. Organization hooks only
type Team struct {

}

func (e *Team) LoadFromJSON (data []byte) (bool, error) {
    err := json.Unmarshal(data, &e)
    if err != nil {
        return false, err
    }
    return true, nil
}

func (e *Team) ConvertToJSON () (string, error) {
    data, err := json.Marshal(&e)
    if err != nil {
        return "", err
    }
    return string(data), nil
}