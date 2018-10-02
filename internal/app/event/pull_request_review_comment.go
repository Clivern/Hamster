package event

import (
    "encoding/json"
)

// Any time a comment on a pull request's unified diff is created, edited, or deleted (in the Files Changed tab).
type PullRequestReviewComment struct {

}

func (e *PullRequestReviewComment) LoadFromJSON (data []byte) (bool, error) {
    err := json.Unmarshal(data, &e)
    if err != nil {
        return false, err
    }
    return true, nil
}

func (e *PullRequestReviewComment) ConvertToJSON () (string, error) {
    data, err := json.Marshal(&e)
    if err != nil {
        return "", err
    }
    return string(data), nil
}