package web

type ReviewCreateRequest struct {
    RestoranId int    `json:"restoran_id"`
    Rating     int    `json:"rating"`
    Comment    string `json:"comment"`
}