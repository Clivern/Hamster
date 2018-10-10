package sender

import (
    "encoding/json"
)

type Label struct {
    Name    string `json:"body"`
    Color   string `json:"color"`
}


func (e *Label) LoadFromJSON (data []byte) (bool, error) {
    err := json.Unmarshal(data, &e)
    if err != nil {
        return false, err
    }
    return true, nil
}

func (e *Label) ConvertToJSON () (string, error) {
    data, err := json.Marshal(&e)
    if err != nil {
        return "", err
    }
    return string(data), nil
}