package event

import (
    "encoding/json"
)

// Any time a Pages site is built or results in a failed build.
type PageBuild struct {

}

func (e *PageBuild) LoadFromJSON (data []byte) (bool, error) {
    err := json.Unmarshal(data, &e)
    if err != nil {
        return false, err
    }
    return true, nil
}

func (e *PageBuild) ConvertToJSON () (string, error) {
    data, err := json.Marshal(&e)
    if err != nil {
        return "", err
    }
    return string(data), nil
}