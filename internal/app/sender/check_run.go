// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package sender

import (
	"encoding/json"
)

type CheckRun struct {
	Name        string   `json:"name,omitempty"`
	HeadSha     string   `json:"head_sha,omitempty"`
	DetailsURL  string   `json:"details_url,omitempty"`
	Status      string   `json:"status,omitempty"`
	ExternalID  string   `json:"external_id,omitempty"`
	Conclusion  string   `json:"conclusion,omitempty"`
	StartedAt   string   `json:"started_at,omitempty"`
	CompletedAt string   `json:"completed_at,omitempty"`
	Output      Output   `json:"output,omitempty"`
	Actions     []Action `json:"actions,omitempty"`
}

type Output struct {
	Title       string       `json:"title,omitempty"`
	Summary     string       `json:"summary,omitempty"`
	Text        string       `json:"text,omitempty"`
	Annotations []Annotation `json:"annotations,omitempty"`
	Images      []Image      `json:"images,omitempty"`
}

type Annotation struct {
	Path            string `json:"path,omitempty"`
	StartLine       int    `json:"start_line,omitempty"`
	EndLine         int    `json:"end_line,omitempty"`
	StartColumn     int    `json:"start_column,omitempty"`
	EndColumn       int    `json:"end_column,omitempty"`
	AnnotationLevel string `json:"annotation_level,omitempty"`
	Message         string `json:"message,omitempty"`
	Title           string `json:"title,omitempty"`
	RawDetails      string `json:"raw_details,omitempty"`
}

type Image struct {
	Alt      string `json:"alt,omitempty"`
	ImageURL string `json:"image_url,omitempty"`
	Caption  string `json:"caption,omitempty"`
}

type Action struct {
	Label       string `json:"label,omitempty"`
	Description string `json:"description,omitempty"`
	Identifier  string `json:"identifier,omitempty"`
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
