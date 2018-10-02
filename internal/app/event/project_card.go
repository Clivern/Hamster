package event

import (
    "encoding/json"
)

// Any time a Project Card is created, edited, moved, converted to an issue, or deleted.
type ProjectCard struct {

}

func (e *ProjectCard) LoadFromJSON (data []byte) (bool, error) {
    err := json.Unmarshal(data, &e)
    if err != nil {
        return false, err
    }
    return true, nil
}

func (e *ProjectCard) ConvertToJSON () (string, error) {
    data, err := json.Marshal(&e)
    if err != nil {
        return "", err
    }
    return string(data), nil
}