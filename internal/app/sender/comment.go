package sender

import (
    "encoding/json"
)

type Comment struct {
    Body string `json:"body"`
}


func (e *Comment) LoadFromJSON (data []byte) bool {
    err := json.Unmarshal(data, &e)
    if err != nil {
        return false
    }
    return true
}

func (e *Comment) ConvertToJSON () (error, string) {
    data, err := json.Marshal(&e)
    if err != nil {
        return err, ""
    }
    return err, string(data)
}