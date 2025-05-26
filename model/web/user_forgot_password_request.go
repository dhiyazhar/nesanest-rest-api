package web

type UserForgotPasswordRequest struct {
    Email       string `json:"email"`
    NewPassword string `json:"new_password"`
}