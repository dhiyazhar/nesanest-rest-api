package repository

import (
    "context"
    "database/sql"
    "nesanest-rest-api/model/domain"
)

type GlobalChatRepositoryImpl struct{}

func NewGlobalChatRepository() GlobalChatRepository {
    return &GlobalChatRepositoryImpl{}
}

func (r *GlobalChatRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, chat domain.GlobalChat) domain.GlobalChat {
    row := tx.QueryRowContext(ctx, "INSERT INTO global_chats (message, anon_id, created_at) VALUES ($1, $2, NOW()) RETURNING id, created_at",
        chat.Message, chat.AnonId)
    err := row.Scan(&chat.Id, &chat.CreatedAt)
    if err != nil {
        panic(err)
    }
    return chat
}

func (r *GlobalChatRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.GlobalChat {
    rows, err := tx.QueryContext(ctx, "SELECT id, message, anon_id, created_at FROM global_chats ORDER BY created_at ASC")
    if err != nil {
        panic(err)
    }
    defer rows.Close()

    var chats []domain.GlobalChat
    for rows.Next() {
        var chat domain.GlobalChat
        err := rows.Scan(&chat.Id, &chat.Message, &chat.AnonId, &chat.CreatedAt)
        if err != nil {
            panic(err)
        }
        chats = append(chats, chat)
    }
    return chats
}