package repository

import (
    "context"
    "database/sql"
    "errors"
    "log/slog"
    "nesanest-rest-api/helper"
    "nesanest-rest-api/model/domain"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
    return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
    SQL := "INSERT INTO users(username, password, email, image_url) VALUES($1, $2, $3, $4) RETURNING id"
    slog.Info("User registered", slog.String("username", user.Username))
    err := tx.QueryRowContext(ctx, SQL, user.Username, user.Password, user.Email, user.ProfileImg).Scan(&user.Id)
    helper.PanicIfError(err)
    return user
}

func (repository *UserRepositoryImpl) UpdateUsername(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
    query := "UPDATE users SET username = $1 WHERE id = $2"
    _, err := tx.ExecContext(ctx, query, user.Username, user.Id)
    if err != nil {
        panic(err)
    }
    return user
}

func (repository *UserRepositoryImpl) UpdatePassword(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
    SQL := "UPDATE users SET password = $1 WHERE id = $2"
    _, err := tx.ExecContext(ctx, SQL, user.Password, user.Id)
    helper.PanicIfError(err)
    return user
}

func (repository *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, user domain.User) {
    SQL := "DELETE FROM users WHERE id = $1"
    _, err := tx.ExecContext(ctx, SQL, user.Id)
    helper.PanicIfError(err)
}

func (repository *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error) {
    SQL := "SELECT id, username, password, email, image_url FROM users WHERE id = $1"
    rows, err := tx.QueryContext(ctx, SQL, userId)
    helper.PanicIfError(err)
    defer rows.Close()

    user := domain.User{}
    if rows.Next() {
        err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.ProfileImg)
        helper.PanicIfError(err)
        return user, nil
    }
    return user, errors.New("user tidak ditemukan")
}

func (r *UserRepositoryImpl) FindByEmail(ctx context.Context, tx *sql.Tx, email string) (domain.User, error) {
    SQL := "SELECT id, username, email, password FROM users WHERE email = $1"
    rows, err := tx.QueryContext(ctx, SQL, email)
    helper.PanicIfError(err)
    defer rows.Close()

    user := domain.User{}
    if rows.Next() {
        err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.Password)
        helper.PanicIfError(err)
        return user, nil
    } else {
        return user, errors.New("user tidak ditemukan")
    }
}

func (repository *UserRepositoryImpl) FindByUsername(ctx context.Context, tx *sql.Tx, username string) (domain.User, error) {
    SQL := "SELECT id, username, password, email, image_url FROM users WHERE username = $1"
    rows, err := tx.QueryContext(ctx, SQL, username)
    helper.PanicIfError(err)
    defer rows.Close()

    user := domain.User{}
    if rows.Next() {
        err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.ProfileImg)
        helper.PanicIfError(err)
        return user, nil
    }
    return user, errors.New("user tidak ditemukan")
}