package event

import (
    "encoding/json"
)

// Any Git push to a Repository, including editing tags or branches.
// Commits via API actions that update references are also counted. This is the default event.
type Push struct {

}

func (e *Push) LoadFromJSON (data []byte) (bool, error) {
    err := json.Unmarshal(data, &e)
    if err != nil {
        return false, err
    }
    return true, nil
}

func (e *Push) ConvertToJSON () (string, error) {
    data, err := json.Marshal(&e)
    if err != nil {
        return "", err
    }
    return string(data), nil
}