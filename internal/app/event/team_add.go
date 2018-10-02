package event

import (
    "encoding/json"
)

// Any time a team is added or modified on a Repository.
type TeamAdd struct {

}

func (e *TeamAdd) LoadFromJSON (data []byte) (bool, error) {
    err := json.Unmarshal(data, &e)
    if err != nil {
        return false, err
    }
    return true, nil
}

func (e *TeamAdd) ConvertToJSON () (string, error) {
    data, err := json.Marshal(&e)
    if err != nil {
        return "", err
    }
    return string(data), nil
}