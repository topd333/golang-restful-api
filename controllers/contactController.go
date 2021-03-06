package controllers

import (
    "encoding/json"
    "net/http"
    "strconv"
    "github.com/gorilla/mux"
    "golang-restful-api/models"
    u "golang-restful-api/utils"
)

var GetContacts = func(w http.ResponseWriter, r *http.Request) {

    id := r.Context().Value("user").(uint)
    data := models.GetContacts(id)
    resp := u.Message(true, "success")
    resp["data"] = data
    u.Respond(w, resp)
}

var GetContact = func(w http.ResponseWriter, r *http.Request) {

    vars := mux.Vars(r)
    key := vars["id"]
    id, err := strconv.ParseUint(key, 10, 64)
    if err != nil {
        u.Respond(w, u.Message(false, "Error while decoding request body"))
        return
    }
    data := models.GetContact(uint(id))
    resp := u.Message(true, "success")
    resp["data"] = data
    u.Respond(w, resp)
}

var CreateContact = func(w http.ResponseWriter, r *http.Request) {

    user := r.Context().Value("user").(uint)
    contact := &models.Contact{}

    err := json.NewDecoder(r.Body).Decode(contact)
    if err != nil {
        u.Respond(w, u.Message(false, "Error while decoding request body"))
        return
    }

    contact.UserId = user
    resp := contact.Create()
    u.Respond(w, resp)
}

var UpdateContact = func(w http.ResponseWriter, r *http.Request) {

    vars := mux.Vars(r)
    key := vars["id"]
    id, err := strconv.ParseUint(key, 10, 64)
    if err != nil {
        u.Respond(w, u.Message(false, "Error while decoding request body"))
        return
    }

    contact := &models.Contact{}
    err = json.NewDecoder(r.Body).Decode(contact)
    if err != nil {
        u.Respond(w, u.Message(false, "Error while decoding request body"))
        return
    }

    data := models.UpdateContact(uint(id), contact)
    resp := u.Message(true, "success")
    resp["data"] = data
    u.Respond(w, resp)
}

var DeleteContact = func(w http.ResponseWriter, r *http.Request) {

    vars := mux.Vars(r)
    key := vars["id"]
    id, err := strconv.ParseUint(key, 10, 64)
    if err != nil {
        u.Respond(w, u.Message(false, "Error while decoding request body"))
        return
    }
    data := models.DeleteContact(uint(id))
    resp := u.Message(true, "success")
    resp["data"] = data
    u.Respond(w, resp)
}
