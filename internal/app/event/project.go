package event

import (
    "encoding/json"
)

// Any time a Project is created, edited, closed, reopened, or deleted.
type Project struct {

}

func (e *Project) LoadFromJSON (data []byte) (bool, error) {
    err := json.Unmarshal(data, &e)
    if err != nil {
        return false, err
    }
    return true, nil
}

func (e *Project) ConvertToJSON () (string, error) {
    data, err := json.Marshal(&e)
    if err != nil {
        return "", err
    }
    return string(data), nil
}