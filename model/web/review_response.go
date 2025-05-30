package web

type ReviewResponse struct {
    Id         int    `json:"id"`
    UserId     int    `json:"user_id"`
    RestoranId int    `json:"restoran_id"`
    Rating     int    `json:"rating"`
    Comment    string `json:"comment"`
    ImageUrl   string `json:"image_url"`
    CreatedAt  string `json:"created_at"`
}