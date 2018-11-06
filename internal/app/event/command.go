// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package event

import (
	"encoding/json"
	"strings"
)

// Command struct
type Command struct {
	Data       string
	Name       string
	Parameters []string
}

// LoadFromJSON update object from json
func (e *Command) LoadFromJSON(data []byte) (bool, error) {
	err := json.Unmarshal(data, &e)
	if err != nil {
		return false, err
	}
	return true, nil
}

// ConvertToJSON convert object to json
func (e *Command) ConvertToJSON() (string, error) {
	data, err := json.Marshal(&e)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// Parse parses the incoming command
func (e *Command) Parse() {
	if strings.Contains(e.Data, "{") && strings.Contains(e.Data, "}") {
		items := strings.Split(e.Data, "{")
		if len(items) == 2 {
			e.Name = strings.Trim(items[0], "/")
			e.Parameters = strings.Split(strings.Trim(items[1], "}"), ",")
		}
	} else {
		e.Name = strings.Trim(e.Data, "/")
	}
}
