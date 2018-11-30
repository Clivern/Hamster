// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package event

import (
	"encoding/json"
)

// Gollum event received any time a Wiki page is updated.
type Gollum struct {
}

// LoadFromJSON update object from json
func (e *Gollum) LoadFromJSON(data []byte) (bool, error) {
	err := json.Unmarshal(data, &e)
	if err != nil {
		return false, err
	}
	return true, nil
}

// ConvertToJSON convert object to json
func (e *Gollum) ConvertToJSON() (string, error) {
	data, err := json.Marshal(&e)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
