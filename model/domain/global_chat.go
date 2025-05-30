package domain

type GlobalChat struct {
    Id        int
    Message   string
    CreatedAt string
    AnonId    string // random id per user/session, atau "Anonymous"
}