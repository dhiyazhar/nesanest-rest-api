package controller

import (
    "net/http"
)

type GlobalChatController interface {
    SendMessage(w http.ResponseWriter, r *http.Request)
    GetMessages(w http.ResponseWriter, r *http.Request)
}