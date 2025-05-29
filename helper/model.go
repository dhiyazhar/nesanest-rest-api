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


func ToReviewResponse(review domain.Review) web.ReviewResponse {
    return web.ReviewResponse{
        Id:         review.Id,
        UserId:     review.UserId,
        RestoranId: review.RestoranId,
        Rating:     review.Rating,
        Comment:    review.Comment,
        ImageUrl:   review.ImageUrl,
        CreatedAt:  review.CreatedAt,
    }
}

func ToReviewResponses(reviews []domain.Review) []web.ReviewResponse {
    var responses []web.ReviewResponse
    for _, review := range reviews {
        responses = append(responses, ToReviewResponse(review))
    }
    return responses
}

func ToGlobalChatResponse(chat domain.GlobalChat) web.GlobalChatResponse {
    return web.GlobalChatResponse{
        AnonId:    chat.AnonId,
        Message:   chat.Message,
        CreatedAt: chat.CreatedAt,
    }
}

func ToGlobalChatResponses(chats []domain.GlobalChat) []web.GlobalChatResponse {
    var responses []web.GlobalChatResponse
    for _, chat := range chats {
        responses = append(responses, ToGlobalChatResponse(chat))
    }
    return responses
}