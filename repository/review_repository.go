package repository

import (
    "context"
    "database/sql"
    "nesanest-rest-api/model/domain"
)

type ReviewRepository interface {
    Save(ctx context.Context, tx *sql.Tx, review domain.Review) domain.Review
    FindByRestoranId(ctx context.Context, tx *sql.Tx, restoranId int) []domain.Review
    FindByUserId(ctx context.Context, tx *sql.Tx, userId int) []domain.Review
}