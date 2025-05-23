package repository

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"
	"nesanest-rest-api/helper"
	"nesanest-rest-api/model/domain"
)

type RestoranRepositoryImpl struct {
}

func NewRestoranRepository() RestoranRepository {
	return &RestoranRepositoryImpl{}
}

func (repository *RestoranRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, restoran domain.Restoran) domain.Restoran {
	SQL := "INSERT INTO restoran(name, description, address, image_url) VALUES ($1, $2, $3, $4) RETURNING id"
	slog.Info("Added to database", restoran.Name, restoran.Description, restoran.Address, restoran.ImageUrl)
	err := tx.QueryRowContext(ctx, SQL, restoran.Name, restoran.Description, restoran.Address, restoran.ImageUrl).Scan(&restoran.Id)
	helper.PanicIfError(err)

	return restoran
}

func (repository *RestoranRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, restoran domain.Restoran) domain.Restoran {
	SQL := "UPDATE restoran SET name = $1, description = $2 WHERE id = $3"
	_, err := tx.ExecContext(ctx, SQL, restoran.Name, restoran.Description, restoran.Id)
	helper.PanicIfError(err)

	return restoran
}

func (repository *RestoranRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, restoran domain.Restoran) {
	SQL := "DELETE FROM restoran WHERE id = $1"
	_, err := tx.ExecContext(ctx, SQL, restoran.Id)
	helper.PanicIfError(err)
}

func (repository *RestoranRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, restoranId int) (domain.Restoran, error) {
	SQL := "SELECT id FROM restoran WHERE id = $1"
	rows, err := tx.QueryContext(ctx, SQL, restoranId)
	helper.PanicIfError(err)
	defer rows.Close()

	restoran := domain.Restoran{}
	if rows.Next() {
		err := rows.Scan(&restoran.Id)
		helper.PanicIfError(err)
		return restoran, nil
	} else {
		return restoran, errors.New("restoran tidak ditemukan")
	}
}

func (repository *RestoranRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Restoran {
	SQL := "SELECT id, name, description FROM restoran ORDER BY id ASC"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var restorans []domain.Restoran
	for rows.Next() {
		restoran := domain.Restoran{}
		err := rows.Scan(&restoran.Id, &restoran.Name, &restoran.Description)
		helper.PanicIfError(err)
		restorans = append(restorans, restoran)
	}

	return restorans
}
