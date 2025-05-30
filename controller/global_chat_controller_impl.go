package controller

import (
    "encoding/json"
    "nesanest-rest-api/helper"
    "nesanest-rest-api/model/web"
    "nesanest-rest-api/service"
    "net/http"
)

type GlobalChatControllerImpl struct {
    GlobalChatService service.GlobalChatService
}

func NewGlobalChatController(service service.GlobalChatService) GlobalChatController {
    return &GlobalChatControllerImpl{GlobalChatService: service}
}

func (c *GlobalChatControllerImpl) SendMessage(w http.ResponseWriter, r *http.Request) {
    var req web.GlobalChatRequest
    err := json.NewDecoder(r.Body).Decode(&req)
    helper.PanicIfError(err)

    // Generate anonId per user/session, atau gunakan "Anonymous"
    anonId := "Anonymous"
    // Bisa juga generate random string per user/session

    resp := c.GlobalChatService.SendMessage(r.Context(), anonId, req)
    helper.WriteToResponseBody(w, resp)
}

func (c *GlobalChatControllerImpl) GetMessages(w http.ResponseWriter, r *http.Request) {
    resp := c.GlobalChatService.GetMessages(r.Context())
    helper.WriteToResponseBody(w, resp)
}