package event

import (
    "encoding/json"
)

// Any time a Repository is created, deleted (organization hooks only), archived, unarchived, made public, or made private.
type Repository struct {

}

func (e *Repository) LoadFromJSON (data []byte) (bool, error) {
    err := json.Unmarshal(data, &e)
    if err != nil {
        return false, err
    }
    return true, nil
}

func (e *Repository) ConvertToJSON () (string, error) {
    data, err := json.Marshal(&e)
    if err != nil {
        return "", err
    }
    return string(data), nil
}