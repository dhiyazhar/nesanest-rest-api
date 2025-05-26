package helper

import (
	"nesanest-rest-api/model/domain"
	"nesanest-rest-api/model/web"
)

func ToRestoranResponse(restoran domain.Restoran) web.RestoranResponse {
	imageUrl := restoran.ImageUrl
	if imageUrl != "" {
		imageUrl = AppConfig.AppBaseUrl + "/" + imageUrl
	}

	return web.RestoranResponse{
		Id:          restoran.Id,
		Name:        restoran.Name,
		Description: restoran.Description,
		Address:     restoran.Address,
		ImageUrl:    imageUrl,
	}
}

func ToRestoranResponses(restorans []domain.Restoran) []web.RestoranResponse {
	var restoranResponses []web.RestoranResponse

	for _, restoran := range restorans {
		restoranResponses = append(restoranResponses, ToRestoranResponse(restoran))
	}

	return restoranResponses
}
