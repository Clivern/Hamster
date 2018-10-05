package event


// All Events
type Raw struct {
    Body    string
    Event   string
}

func (e *Raw) SetBody (body string) {
    e.Body = body
}

func (e *Raw) GetBody () string {
    return e.Body
}

func (e *Raw) SetEvent (event string) {
    e.Event = event
}

func (e *Raw) GetEvent () string {
    return e.Event
}