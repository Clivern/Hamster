// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package sender

import (
	"encoding/json"
	"time"
)

type CheckRun struct {
	Name       string    `json:"name"`
	HeadSha    string    `json:"head_sha"`
	Status     string    `json:"status"`
	ExternalID string    `json:"external_id"`
	DetailsURL string    `json:"details_url"`
	StartedAt  time.Time `json:"started_at"`
	Output     struct {
		Title   string `json:"title,omitempty"`
		Summary string `json:"summary,omitempty"`
		Text    string `json:"text,omitempty"`
	} `json:"output,omitempty"`
}

func (e *CheckRun) LoadFromJSON(data []byte) (bool, error) {
	err := json.Unmarshal(data, &e)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (e *CheckRun) ConvertToJSON() (string, error) {
	data, err := json.Marshal(&e)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
