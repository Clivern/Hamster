package event

import (
    "encoding/json"
)

// Any time a Project Column is created, edited, moved, or deleted.
type ProjectColumn struct {

}

func (e *ProjectColumn) LoadFromJSON (data []byte) (bool, error) {
    err := json.Unmarshal(data, &e)
    if err != nil {
        return false, err
    }
    return true, nil
}

func (e *ProjectColumn) ConvertToJSON () (string, error) {
    data, err := json.Marshal(&e)
    if err != nil {
        return "", err
    }
    return string(data), nil
}