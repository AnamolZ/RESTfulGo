package main

import (
    "fmt"
    "log"
    "net/http"
    "webservices/services/logger"
    "webservices/services/request"
    "github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Fprintln(w, "Welcome to the Index Page!")
}

func main() {
    // Initialize the logger
    if err := logger.InitializeLog(); err != nil {
        log.Fatalf("Error initializing log file: %v", err)
    }
    defer logger.CloseLog()

    // Set up the router
    router := httprouter.New()
    router.GET("/", Index)

    // Log the server start
    log.Println("Starting server on :8080")

    // Start the server
    if err := http.ListenAndServe(":8080", router); err != nil {
        log.Fatalf("Error starting server: %v", err)
    }

    // Make an HTTP request
    url := "https://www.google.com"
    request.Request(url)
}
