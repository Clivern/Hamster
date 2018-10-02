package event

import (
    "encoding/json"
)

// Any time an organization blocks or unblocks a user. Organization hooks only.
type OrgBlock struct {

}

func (e *OrgBlock) LoadFromJSON (data []byte) (bool, error) {
    err := json.Unmarshal(data, &e)
    if err != nil {
        return false, err
    }
    return true, nil
}

func (e *OrgBlock) ConvertToJSON () (string, error) {
    data, err := json.Marshal(&e)
    if err != nil {
        return "", err
    }
    return string(data), nil
}