package service

import (
	"context"
	"database/sql"
	"log/slog"
	"nesanest-rest-api/exception"
	"nesanest-rest-api/helper"
	"nesanest-rest-api/model/domain"
	"nesanest-rest-api/model/web"
	"nesanest-rest-api/repository"

	"github.com/go-playground/validator/v10"
)

type RestoranServiceImpl struct {
	RestoranRepository repository.RestoranRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewRestoranService(restoranRepository repository.RestoranRepository, DB *sql.DB, validate *validator.Validate) RestoranService {
	return &RestoranServiceImpl{
		RestoranRepository: restoranRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

func (service *RestoranServiceImpl) Create(ctx context.Context, request web.RestoranCreateRequest) web.RestoranResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	restoran := domain.Restoran{
		Name:        request.Name,
		Description: request.Description,
	}

	restoran = service.RestoranRepository.Save(ctx, tx, restoran)

	return helper.ToRestoranResponse(restoran)
}

func (service *RestoranServiceImpl) Update(ctx context.Context, request web.RestoranUpdateRequest) web.RestoranResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	slog.Info("Finding ID on the database...", slog.Any("", request.Id))
	restoran, err := service.RestoranRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	restoran.Id = request.Id
	restoran.Name = request.Name
	restoran.Description = request.Description

	slog.Info("Updating on the database", restoran.Name, restoran.Description)
	restoran = service.RestoranRepository.Update(ctx, tx, restoran)

	return helper.ToRestoranResponse(restoran)
}

func (service *RestoranServiceImpl) Delete(ctx context.Context, categoryId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	restoran, err := service.RestoranRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.RestoranRepository.Delete(ctx, tx, restoran)
}

func (service *RestoranServiceImpl) FindById(ctx context.Context, categoryId int) web.RestoranResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	restoran, err := service.RestoranRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToRestoranResponse(restoran)
}

func (service *RestoranServiceImpl) FindAll(ctx context.Context) []web.RestoranResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	restorans := service.RestoranRepository.FindAll(ctx, tx)

	return helper.ToRestoranResponses(restorans)
}
