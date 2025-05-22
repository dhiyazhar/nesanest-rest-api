package service

import (
	"context"
	"nesanest-rest-api/model/web"
)

type RestoranService interface {
	Create(ctx context.Context, request web.RestoranCreateRequest) web.RestoranResponse
	Update(ctx context.Context, request web.RestoranUpdateRequest) web.RestoranResponse
	Delete(ctx context.Context, categoryId int)
	FindById(ctx context.Context, categoryId int) web.RestoranResponse
	FindAll(ctx context.Context) []web.RestoranResponse
}
