package event

import (
    "encoding/json"
)

// Any time a user accepts an invitation or is removed as a collaborator to a repository, or has their permissions modified.
type Member struct {

}

func (e *Member) LoadFromJSON (data []byte) (bool, error) {
    err := json.Unmarshal(data, &e)
    if err != nil {
        return false, err
    }
    return true, nil
}

func (e *Member) ConvertToJSON () (string, error) {
    data, err := json.Marshal(&e)
    if err != nil {
        return "", err
    }
    return string(data), nil
}