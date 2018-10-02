package event

import (
    "encoding/json"
)

// Any time a GitHub App is installed or uninstalled.
type Installation struct {

}

func (e *Installation) LoadFromJSON (data []byte) (bool, error) {
    err := json.Unmarshal(data, &e)
    if err != nil {
        return false, err
    }
    return true, nil
}

func (e *Installation) ConvertToJSON () (string, error) {
    data, err := json.Marshal(&e)
    if err != nil {
        return "", err
    }
    return string(data), nil
}