package event

import (
    "encoding/json"
)

// Any time a check suite is completed, requested, or rerequested.
type CheckSuite struct {

}

func (e *CheckSuite) LoadFromJSON (data []byte) (bool, error) {
    err := json.Unmarshal(data, &e)
    if err != nil {
        return false, err
    }
    return true, nil
}

func (e *CheckSuite) ConvertToJSON () (string, error) {
    data, err := json.Marshal(&e)
    if err != nil {
        return "", err
    }
    return string(data), nil
}