package helper

import (
	"nesanest-rest-api/model/domain"
	"nesanest-rest-api/model/web"
)

func ToRestoranResponse(restoran domain.Restoran) web.RestoranResponse {
	return web.RestoranResponse{
		Id:          restoran.Id,
		Name:        restoran.Name,
		Description: restoran.Description,
		Address:     restoran.Address,
		ImageUrl:    restoran.ImageUrl,
	}
}

func ToRestoranResponses(restorans []domain.Restoran) []web.RestoranResponse {
	var restoranResponses []web.RestoranResponse

	for _, restoran := range restorans {
		restoranResponses = append(restoranResponses, ToRestoranResponse(restoran))
	}

	return restoranResponses
}


func ToUserResponse(user domain.User) web.UserResponse {
	return web.UserResponse{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
		ProfileImg: user.ProfileImg,
	}
}

func ToUserResponses(users []domain.User) []web.UserResponse {
	var userResponses []web.UserResponse

	for _, user := range users {
		userResponses = append(userResponses, ToUserResponse(user))
	}

	return userResponses
}