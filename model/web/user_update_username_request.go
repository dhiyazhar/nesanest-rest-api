package web

type UserUpdateUsernameRequest struct {
    Id       int    `json:"id"`
    Username string `json:"username"`
}