package service

import (
    "context"
    "nesanest-rest-api/model/web"
)

type ReviewService interface {
    CreateReview(ctx context.Context, userId int, request web.ReviewCreateRequest, imageUrl string) web.ReviewResponse
    GetReviewsByRestoran(ctx context.Context, restoranId int) []web.ReviewResponse
    GetReviewsByUser(ctx context.Context, userId int) []web.ReviewResponse
}