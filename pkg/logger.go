package pkg

import (
    "os"
    "fmt"
    "time"
    "github.com/google/logger"
)

const LOGS_PATH = "var/logs"

func Info(v ...interface{}) {

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

        out.Info(v...)
    }
}

func Infoln(v ...interface{}) {

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

        out.Infoln(v...)
    }
}

func Infof(format string, v ...interface{}) {

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

        out.Infof(format, v...)
    }
}

func Warning(v ...interface{}) {

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

        out.Warning(v...)
    }
}

func Warningln(v ...interface{}) {

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

        out.Warningln(v...)
    }
}


func Warningf(format string, v ...interface{}) {

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

        out.Warningf(format, v...)
    }
}

func Error(v ...interface{}) {

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

        out.Error(v...)
    }
}

func Errorln(v ...interface{}) {

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

        out.Errorln(v...)
    }
}

func Errorf(format string, v ...interface{}) {

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

        out.Errorf(format, v...)
    }
}

func Fatal(v ...interface{}) {

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

        out.Fatal(v...)
    }
}

func Fatalln(v ...interface{}) {

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

        out.Fatalln(v...)
    }
}

func Fatalf(format string, v ...interface{}) {

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

        out.Fatalf(format, v...)
    }
}