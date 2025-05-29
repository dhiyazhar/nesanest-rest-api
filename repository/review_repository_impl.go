package repository

import (
    "context"
    "database/sql"
    "nesanest-rest-api/model/domain"
)

type ReviewRepositoryImpl struct{}

func NewReviewRepository() ReviewRepository {
    return &ReviewRepositoryImpl{}
}

func (r *ReviewRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, review domain.Review) domain.Review {
    SQL := "INSERT INTO reviews(user_id, restoran_id, rating, comment, image_url, created_at) VALUES($1, $2, $3, $4, $5, NOW()) RETURNING id, created_at"
    err := tx.QueryRowContext(ctx, SQL, review.UserId, review.RestoranId, review.Rating, review.Comment, review.ImageUrl).
        Scan(&review.Id, &review.CreatedAt)
    if err != nil {
        panic(err)
    }
    return review
}

func (r *ReviewRepositoryImpl) FindByRestoranId(ctx context.Context, tx *sql.Tx, restoranId int) []domain.Review {
    SQL := "SELECT id, user_id, restoran_id, rating, comment, image_url, created_at FROM reviews WHERE restoran_id = $1"
    rows, err := tx.QueryContext(ctx, SQL, restoranId)
    if err != nil {
        panic(err)
    }
    defer rows.Close()

    var reviews []domain.Review
    for rows.Next() {
        var review domain.Review
        err := rows.Scan(&review.Id, &review.UserId, &review.RestoranId, &review.Rating, &review.Comment, &review.ImageUrl, &review.CreatedAt)
        if err != nil {
            panic(err)
        }
        reviews = append(reviews, review)
    }
    return reviews
}

func (r *ReviewRepositoryImpl) FindByUserId(ctx context.Context, tx *sql.Tx, userId int) []domain.Review {
    SQL := "SELECT id, user_id, restoran_id, rating, comment, image_url, created_at FROM reviews WHERE user_id = $1"
    rows, err := tx.QueryContext(ctx, SQL, userId)
    if err != nil {
        panic(err)
    }
    defer rows.Close()

    var reviews []domain.Review
    for rows.Next() {
        var review domain.Review
        err := rows.Scan(&review.Id, &review.UserId, &review.RestoranId, &review.Rating, &review.Comment, &review.ImageUrl, &review.CreatedAt)
        if err != nil {
            panic(err)
        }
        reviews = append(reviews, review)
    }
    return reviews
}

