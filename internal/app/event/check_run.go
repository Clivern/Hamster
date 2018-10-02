package event

import (
    "encoding/json"
)

// Any time a check run is created, requested, or rerequested.
type CheckRun struct {

}

func (e *CheckRun) LoadFromJSON (data []byte) (bool, error) {
    err := json.Unmarshal(data, &e)
    if err != nil {
        return false, err
    }
    return true, nil
}

func (e *CheckRun) ConvertToJSON () (string, error) {
    data, err := json.Marshal(&e)
    if err != nil {
        return "", err
    }
    return string(data), nil
}