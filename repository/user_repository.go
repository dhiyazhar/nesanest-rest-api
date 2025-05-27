package repository

import (
	"context"
	"database/sql"
	"nesanest-rest-api/model/domain"
)

type UserRepository interface {
	Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	UpdateUsername(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	UpdatePassword(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	Delete(ctx context.Context, tx *sql.Tx, user domain.User)
	FindById(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error)
    FindAll(ctx context.Context, tx *sql.Tx) ([]domain.User, error)
	FindByEmail(ctx context.Context, tx *sql.Tx, email string) (domain.User, error)
	FindByUsername(ctx context.Context, tx *sql.Tx, username string) (domain.User, error)
}