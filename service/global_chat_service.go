package service

import (
    "context"
    "nesanest-rest-api/model/web"
)

type GlobalChatService interface {
    SendMessage(ctx context.Context, anonId string, request web.GlobalChatRequest) web.GlobalChatResponse
    GetMessages(ctx context.Context) []web.GlobalChatResponse
}