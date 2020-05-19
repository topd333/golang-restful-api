package main

import (
    "fmt"
    "github.com/gorilla/mux"
    "net/http"
    "os"
)

func main() {
    router := mux.NewRouter()

    port := os.Getenv("PORT")
    if port == "" {
        port = "8000"
    }

    // Launch the app, visit localhost:8000
    err := http.ListenAndServe(":"+port, router)
    if err != nil {
        fmt.Print(err)
    }
}
