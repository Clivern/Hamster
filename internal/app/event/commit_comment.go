package event

import (
    "encoding/json"
)

// Any time a Commit is commented on.
type CommitComment struct {

}

func (e *CommitComment) LoadFromJSON (data []byte) (bool, error) {
    err := json.Unmarshal(data, &e)
    if err != nil {
        return false, err
    }
    return true, nil
}

func (e *CommitComment) ConvertToJSON () (string, error) {
    data, err := json.Marshal(&e)
    if err != nil {
        return "", err
    }
    return string(data), nil
}