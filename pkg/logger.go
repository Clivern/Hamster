package pkg

import (
    "os"
    "fmt"
    "time"
    "github.com/google/logger"
)

const LOGS_PATH = "var/logs"

func Info(msg string) {

    logLevel := os.Getenv("AppLogLevel")
    ok := logLevel == "info"

    if ok {
        current_time := time.Now().Local()
        file := fmt.Sprintf("%s/%s.log", LOGS_PATH, current_time.Format("2006-01-02"))
        lf, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)

        if err != nil {
            logger.Fatalf("Failed to open log file: %v", err)
        }

        defer lf.Close()

        out := logger.Init("Hamster", false, false, lf)
        defer out.Close()

        out.Info(msg)
    }
}

func Warning(msg string) {

    logLevel := os.Getenv("AppLogLevel")
    ok := logLevel == "info" || logLevel == "warning"

    if ok {
        current_time := time.Now().Local()
        file := fmt.Sprintf("%s/%s.log", LOGS_PATH, current_time.Format("2006-01-02"))
        lf, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)

        if err != nil {
            logger.Fatalf("Failed to open log file: %v", err)
        }

        defer lf.Close()

        out := logger.Init("Hamster", false, false, lf)
        defer out.Close()

        out.Warning(msg)
    }
}

func Error(msg string) {

    logLevel := os.Getenv("AppLogLevel")
    ok := logLevel == "info" || logLevel == "warning" || logLevel == "error"

    if ok {
        current_time := time.Now().Local()
        file := fmt.Sprintf("%s/%s.log", LOGS_PATH, current_time.Format("2006-01-02"))
        lf, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)

        if err != nil {
            logger.Fatalf("Failed to open log file: %v", err)
        }

        defer lf.Close()

        out := logger.Init("Hamster", false, false, lf)
        defer out.Close()

        out.Error(msg)
    }
}

func Fatal(msg string) {

    logLevel := os.Getenv("AppLogLevel")
    ok := logLevel == "info" || logLevel == "warning" || logLevel == "error" || logLevel == "fatal"

    if ok {
        current_time := time.Now().Local()
        file := fmt.Sprintf("%s/%s.log", LOGS_PATH, current_time.Format("2006-01-02"))
        lf, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)

        if err != nil {
            logger.Fatalf("Failed to open log file: %v", err)
        }

        defer lf.Close()

        out := logger.Init("Hamster", false, false, lf)
        defer out.Close()

        out.Fatal(msg)
    }
}