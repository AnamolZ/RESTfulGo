package request

import (
    "fmt"
    "net/http"
    "webservices/services/logger"
)

func Request(url string) {
    resp, err := http.Get(url)
    if err != nil {
        logger.LogRequest(url, 0)
        return
    }
    defer resp.Body.Close()

    logger.LogRequest(url, resp.StatusCode)
    fmt.Println("HTTP Status Code:", resp.StatusCode)
}
