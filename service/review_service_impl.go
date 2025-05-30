package service

import (
    "context"
	"database/sql"
    "nesanest-rest-api/helper"
    "nesanest-rest-api/model/domain"
    "nesanest-rest-api/model/web"
    "nesanest-rest-api/repository"
)

type ReviewServiceImpl struct {
    ReviewRepository repository.ReviewRepository
    DB               *sql.DB
}

func NewReviewService(reviewRepository repository.ReviewRepository, db *sql.DB) ReviewService {
    return &ReviewServiceImpl{
        ReviewRepository: reviewRepository,
        DB:               db,
    }
}

func (s *ReviewServiceImpl) CreateReview(ctx context.Context, userId int, request web.ReviewCreateRequest, imageUrl string) web.ReviewResponse {
    tx, err := s.DB.Begin()
    helper.PanicIfError(err)
    defer helper.CommitOrRollback(tx)

    review := domain.Review{
        UserId:     userId,
        RestoranId: request.RestoranId,
        Rating:     request.Rating,
        Comment:    request.Comment,
        ImageUrl:   imageUrl,
    }
    review = s.ReviewRepository.Save(ctx, tx, review)
    return helper.ToReviewResponse(review)
}

func (s *ReviewServiceImpl) GetReviewsByRestoran(ctx context.Context, restoranId int) []web.ReviewResponse {
    tx, err := s.DB.Begin()
    helper.PanicIfError(err)
    defer helper.CommitOrRollback(tx)

    reviews := s.ReviewRepository.FindByRestoranId(ctx, tx, restoranId)
    return helper.ToReviewResponses(reviews)
}

func (s *ReviewServiceImpl) GetReviewsByUser(ctx context.Context, userId int) []web.ReviewResponse {
    tx, err := s.DB.Begin()
    helper.PanicIfError(err)
    defer helper.CommitOrRollback(tx)

    reviews := s.ReviewRepository.FindByUserId(ctx, tx, userId)
    return helper.ToReviewResponses(reviews)
}