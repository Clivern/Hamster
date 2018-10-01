package pkg

import (
    _ "log"
    _ "os"
    "fmt"
)

type Logger struct {
    FileName    string
    FilePath    string
}

func (e Logger) OpenLog() string {
    return fmt.Sprintf("%s/%s", e.FilePath, e.FileName)
}