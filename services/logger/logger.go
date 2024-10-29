package logger

import (
    "log"
    "os"
)

var logFile *os.File

func InitializeLog() error {
    var err error
    logFile, err = os.OpenFile("logger.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        return err
    }
    log.SetOutput(logFile)
    return nil
}

func LogRequest(url string, statusCode int) {
    log.Printf("Requested URL: %s, HTTP Status Code: %d\n", url, statusCode)
}

func CloseLog() {
    if logFile != nil {
        logFile.Close()
    }
}
