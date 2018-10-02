package event

import (
    "encoding/json"
)

// Any time a repository is added or removed from an installation.
type InstallationRepositories struct {

}

func (e *InstallationRepositories) LoadFromJSON (data []byte) (bool, error) {
    err := json.Unmarshal(data, &e)
    if err != nil {
        return false, err
    }
    return true, nil
}

func (e *InstallationRepositories) ConvertToJSON () (string, error) {
    data, err := json.Marshal(&e)
    if err != nil {
        return "", err
    }
    return string(data), nil
}