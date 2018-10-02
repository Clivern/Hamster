package event

import (
    "encoding/json"
)

// Any time a pull request review is submitted, edited, or dismissed.
type PullRequestReview struct {

}

func (e *PullRequestReview) LoadFromJSON (data []byte) (bool, error) {
    err := json.Unmarshal(data, &e)
    if err != nil {
        return false, err
    }
    return true, nil
}

func (e *PullRequestReview) ConvertToJSON () (string, error) {
    data, err := json.Marshal(&e)
    if err != nil {
        return "", err
    }
    return string(data), nil
}