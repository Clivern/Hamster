// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package event

// Raw event
type Raw struct {
	Body  string
	Event string
}

// SetBody set event body
func (e *Raw) SetBody(body string) {
	e.Body = body
}

// GetBody gets event body
func (e *Raw) GetBody() string {
	return e.Body
}

// SetEvent sets event
func (e *Raw) SetEvent(event string) {
	e.Event = event
}

// GetEvent gets event
func (e *Raw) GetEvent() string {
	return e.Event
}
