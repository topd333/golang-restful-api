package app

import (
    "net/http"
    u "golang-restful-api/utils"
)

var NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    response := make(map[string]interface{})
    response = u.Message(false, "This resources was not found on the server")
    w.WriteHeader(http.StatusNotFound)
    u.Respond(w, response)
    return
})
