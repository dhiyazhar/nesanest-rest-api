package repository

import (
	"context"
	"database/sql"
	"nesanest-rest-api/model/domain"
)

type RestoranRepository interface {
	Save(ctx context.Context, tx *sql.Tx, restoran domain.Restoran) domain.Restoran
	Update(ctx context.Context, tx *sql.Tx, restoran domain.Restoran) domain.Restoran
	Delete(ctx context.Context, tx *sql.Tx, restoran domain.Restoran)
	FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Restoran, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Restoran
}
