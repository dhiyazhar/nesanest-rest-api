package domain

type Review struct {
    Id         int
    UserId     int
    RestoranId int
    Rating     int
    Comment    string
    ImageUrl   string
    CreatedAt  string
}