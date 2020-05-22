package main

import (
    "os"
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
    "golang-restful-api/app"
    "golang-restful-api/controllers"
)

func main() {
    router := mux.NewRouter()

    router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
    router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
    outer.HandleFunc("/api/contacts/new", controllers.CreateContact).Methods("POST")
    router.HandleFunc("/api/me/contacts", controllers.GetContactsFor).Methods("GET")

    router.Use(app.JwtAuthentication) //attach JWT auth middleware

    router.NotFoundHandler = app.NotFoundHandler

    port := os.Getenv("PORT")
    if port == "" {
        port = "8000"
    }

    fmt.Println("Server loaded on Port:" + port)

    // Launch the app, visit localhost:8000
    err := http.ListenAndServe(":"+port, router)
    if err != nil {
        fmt.Print(err)
    }
}
