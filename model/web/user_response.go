package web

type UserResponse struct {
    Id         int    `json:"id"`
    Username   string `json:"username"`
    Email      string `json:"email"`
    ProfileImg string `json:"profile_img"`
}