package web

type UserUpdatePasswordRequest struct {
    Id          int    `json:"id"`
    OldPassword string `json:"old_password"`
    NewPassword string `json:"new_password"`
}