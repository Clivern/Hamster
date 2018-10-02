package event

import (
    "encoding/json"
)

// Any time someone revokes their authorization of a GitHub App.
// GitHub Apps receive this webhook by default and cannot unsubscribe from this event.
type GithubAppAuthorization struct {

}

func (e *GithubAppAuthorization) LoadFromJSON (data []byte) (bool, error) {
    err := json.Unmarshal(data, &e)
    if err != nil {
        return false, err
    }
    return true, nil
}

func (e *GithubAppAuthorization) ConvertToJSON () (string, error) {
    data, err := json.Marshal(&e)
    if err != nil {
        return "", err
    }
    return string(data), nil
}