package controller

import (
    "fmt"
    "io"
    "net/http"
    "os"
    "strconv"
    "strings"
    "time"
    "nesanest-rest-api/helper"
    "nesanest-rest-api/model/web"
    "nesanest-rest-api/service"
)

type ReviewControllerImpl struct {
    ReviewService service.ReviewService
}

func NewReviewController(reviewService service.ReviewService) ReviewController {
    return &ReviewControllerImpl{
        ReviewService: reviewService,
    }
}

// Mendukung form-data (dengan upload gambar) dan JSON (tanpa gambar)
func (c *ReviewControllerImpl) CreateReview(w http.ResponseWriter, r *http.Request) {
    authHeader := r.Header.Get("Authorization")
    tokenString := ""
    if strings.HasPrefix(authHeader, "Bearer ") {
        tokenString = strings.TrimPrefix(authHeader, "Bearer ")
    }
    claims, err := helper.ParseJWT(tokenString)
    if err != nil {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }
    userId := int(claims["user_id"].(float64))

    contentType := r.Header.Get("Content-Type")
    var restoranId, rating int
    var comment, imageUrl string

    if strings.HasPrefix(contentType, "application/json") {
        var req web.ReviewCreateRequest
        helper.ReadFromRequestBody(r, &req)
        restoranId = req.RestoranId
        rating = req.Rating
        comment = req.Comment
        imageUrl = ""
    } else {
        // Ambil data dari form
        restoranId, _ = strconv.Atoi(r.FormValue("restoran_id"))
        rating, _ = strconv.Atoi(r.FormValue("rating"))
        comment = r.FormValue("comment")

        // Proses upload file jika ada
        file, handler, err := r.FormFile("image")
        if err == nil {
            defer file.Close()
            filename := fmt.Sprintf("uploads/review/%d_%s", time.Now().UnixNano(), handler.Filename)
            f, err := os.Create(filename)
            if err == nil {
                defer f.Close()
                io.Copy(f, file)
                imageUrl = filename
            }
        }
    }

    // Validasi rating
    if rating < 1 || rating > 5 {
        http.Error(w, "Rating harus antara 1 sampai 5", http.StatusBadRequest)
        return
    }

    // Validasi restoranId
    if restoranId <= 0 {
        http.Error(w, "Restoran ID tidak valid", http.StatusBadRequest)
        return
    }

    request := web.ReviewCreateRequest{
        RestoranId: restoranId,
        Rating:     rating,
        Comment:    comment,
    }

    response := c.ReviewService.CreateReview(r.Context(), userId, request, imageUrl)
    helper.WriteToResponseBody(w, response)
}

func (c *ReviewControllerImpl) GetReviewsByRestoran(w http.ResponseWriter, r *http.Request, restoranId string) {
    id, _ := strconv.Atoi(restoranId)
    response := c.ReviewService.GetReviewsByRestoran(r.Context(), id)
    helper.WriteToResponseBody(w, response)
}

func (c *ReviewControllerImpl) GetReviewsByUser(w http.ResponseWriter, r *http.Request) {
    authHeader := r.Header.Get("Authorization")
    tokenString := ""
    if strings.HasPrefix(authHeader, "Bearer ") {
        tokenString = strings.TrimPrefix(authHeader, "Bearer ")
    }
    claims, err := helper.ParseJWT(tokenString)
    if err != nil {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }
    userId := int(claims["user_id"].(float64))

    response := c.ReviewService.GetReviewsByUser(r.Context(), userId)
    helper.WriteToResponseBody(w, response)
}