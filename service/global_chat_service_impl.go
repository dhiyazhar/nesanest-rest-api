package service

import (
    "context"
    "database/sql"
    "nesanest-rest-api/helper"
    "nesanest-rest-api/model/domain"
    "nesanest-rest-api/model/web"
    "nesanest-rest-api/repository"
)

type GlobalChatServiceImpl struct {
    GlobalChatRepository repository.GlobalChatRepository
    DB                   *sql.DB
}

func NewGlobalChatService(repo repository.GlobalChatRepository, db *sql.DB) GlobalChatService {
    return &GlobalChatServiceImpl{
        GlobalChatRepository: repo,
        DB:                   db,
    }
}

func (s *GlobalChatServiceImpl) SendMessage(ctx context.Context, anonId string, request web.GlobalChatRequest) web.GlobalChatResponse {
    tx, err := s.DB.Begin()
    helper.PanicIfError(err)
    defer helper.CommitOrRollback(tx)

    chat := domain.GlobalChat{
        Message: request.Message,
        AnonId:  anonId,
    }
    chat = s.GlobalChatRepository.Save(ctx, tx, chat)
    return helper.ToGlobalChatResponse(chat)
}

func (s *GlobalChatServiceImpl) GetMessages(ctx context.Context) []web.GlobalChatResponse {
    tx, err := s.DB.Begin()
    helper.PanicIfError(err)
    defer helper.CommitOrRollback(tx)

    chats := s.GlobalChatRepository.FindAll(ctx, tx)
    return helper.ToGlobalChatResponses(chats)
}