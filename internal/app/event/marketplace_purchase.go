// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package event

import (
	"encoding/json"
)

// MarketplacePurchase event received any time a user purchases, cancels, or changes their GitHub Marketplace plan.
type MarketplacePurchase struct {
}

// LoadFromJSON update object from json
func (e *MarketplacePurchase) LoadFromJSON(data []byte) (bool, error) {
	err := json.Unmarshal(data, &e)
	if err != nil {
		return false, err
	}
	return true, nil
}

// ConvertToJSON convert object to json
func (e *MarketplacePurchase) ConvertToJSON() (string, error) {
	data, err := json.Marshal(&e)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
