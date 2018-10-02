package event

import (
    "encoding/json"
)

// Any time a Milestone is created, closed, opened, edited, or deleted.
type Milestone struct {

}

func (e *Milestone) LoadFromJSON (data []byte) (bool, error) {
    err := json.Unmarshal(data, &e)
    if err != nil {
        return false, err
    }
    return true, nil
}

func (e *Milestone) ConvertToJSON () (string, error) {
    data, err := json.Marshal(&e)
    if err != nil {
        return "", err
    }
    return string(data), nil
}