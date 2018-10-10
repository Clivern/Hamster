// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package event

//revive:disable:exported

import (
	"encoding/json"
)

// Any time a successful or unsuccessful repository import finishes for a GitHub organization or a personal repository.
type RepositoryImport struct {
}

func (e *RepositoryImport) LoadFromJSON(data []byte) (bool, error) {
	err := json.Unmarshal(data, &e)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (e *RepositoryImport) ConvertToJSON() (string, error) {
	data, err := json.Marshal(&e)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
