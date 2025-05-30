package repository

import (
    "context"
    "database/sql"
    "nesanest-rest-api/model/domain"
)

type GlobalChatRepository interface {
    Save(ctx context.Context, tx *sql.Tx, chat domain.GlobalChat) domain.GlobalChat
    FindAll(ctx context.Context, tx *sql.Tx) []domain.GlobalChat
}