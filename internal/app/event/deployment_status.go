package event

import (
    "encoding/json"
)

// Any time a deployment for a Repository has a status update from the API.
type DeploymentStatus struct {

}

func (e *DeploymentStatus) LoadFromJSON (data []byte) (bool, error) {
    err := json.Unmarshal(data, &e)
    if err != nil {
        return false, err
    }
    return true, nil
}

func (e *DeploymentStatus) ConvertToJSON () (string, error) {
    data, err := json.Marshal(&e)
    if err != nil {
        return "", err
    }
    return string(data), nil
}