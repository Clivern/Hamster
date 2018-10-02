package event

import (
    "encoding/json"
)

// Any time a User is added or removed from a team. Organization hooks only.
type Membership struct {

}

func (e *Membership) LoadFromJSON (data []byte) (bool, error) {
    err := json.Unmarshal(data, &e)
    if err != nil {
        return false, err
    }
    return true, nil
}

func (e *Membership) ConvertToJSON () (string, error) {
    data, err := json.Marshal(&e)
    if err != nil {
        return "", err
    }
    return string(data), nil
}