package event

import (
    "encoding/json"
)

// Any time a User stars a Repository.
type Watch struct {

}

func (e *Watch) LoadFromJSON (data []byte) (bool, error) {
    err := json.Unmarshal(data, &e)
    if err != nil {
        return false, err
    }
    return true, nil
}

func (e *Watch) ConvertToJSON () (string, error) {
    data, err := json.Marshal(&e)
    if err != nil {
        return "", err
    }
    return string(data), nil
}