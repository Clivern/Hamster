// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package pkg

//revive:disable:exported

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Request(method string, url string, body string, token string) (string, error) {

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBufferString(body))

	if err != nil {
		return "", err
	}

	req.Header.Add("Authorization", fmt.Sprintf("token %s", token))

	resp, err := client.Do(req)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body_byte, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	return string(body_byte), nil
}
