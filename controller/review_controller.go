package controller

import (
    "net/http"
)

type ReviewController interface {
    CreateReview(w http.ResponseWriter, r *http.Request)
    GetReviewsByRestoran(w http.ResponseWriter, r *http.Request, restoranId string)
    GetReviewsByUser(w http.ResponseWriter, r *http.Request)
    
}