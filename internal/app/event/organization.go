package event

import (
    "encoding/json"
)

// Any time a user is added, removed, or invited to an Organization. Organization hooks only.
type Organization struct {

}

func (e *Organization) LoadFromJSON (data []byte) (bool, error) {
    err := json.Unmarshal(data, &e)
    if err != nil {
        return false, err
    }
    return true, nil
}

func (e *Organization) ConvertToJSON () (string, error) {
    data, err := json.Marshal(&e)
    if err != nil {
        return "", err
    }
    return string(data), nil
}