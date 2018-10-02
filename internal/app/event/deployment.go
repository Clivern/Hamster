package event

import (
    "encoding/json"
)

// Any time a Repository has a new deployment created from the API.
type Deployment struct {

}

func (e *Deployment) LoadFromJSON (data []byte) (bool, error) {
    err := json.Unmarshal(data, &e)
    if err != nil {
        return false, err
    }
    return true, nil
}

func (e *Deployment) ConvertToJSON () (string, error) {
    data, err := json.Marshal(&e)
    if err != nil {
        return "", err
    }
    return string(data), nil
}