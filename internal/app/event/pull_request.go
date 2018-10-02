package event

import (
    "encoding/json"
)

// Any time a pull request is assigned, unassigned, labeled, unlabeled, opened, edited, closed, reopened, or synchronized (updated due to a new push in the branch that the pull request is tracking).
// Also any time a pull request review is requested, or a review request is removed.
type PullRequest struct {

}

func (e *PullRequest) LoadFromJSON (data []byte) (bool, error) {
    err := json.Unmarshal(data, &e)
    if err != nil {
        return false, err
    }
    return true, nil
}

func (e *PullRequest) ConvertToJSON () (string, error) {
    data, err := json.Marshal(&e)
    if err != nil {
        return "", err
    }
    return string(data), nil
}