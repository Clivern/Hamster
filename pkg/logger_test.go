package pkg

import (
    "testing"
)

func TestHello(t *testing.T) {

    log := Logger{FileName: "file", FilePath: "path"}
    got := log.OpenLog()
    want := "path/file"

    if got != want {
        t.Errorf("got '%s' want '%s'", got, want)
    }
}