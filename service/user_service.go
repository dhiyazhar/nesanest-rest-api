package service

import (
	"context"
	"nesanest-rest-api/model/web"
)

type UserService interface {
	Register(ctx context.Context, request web.UserRegisterRequest) web.UserResponse
	Login(ctx context.Context, request web.UserLoginRequest) (web.UserResponse, string)
	UpdateProfile(ctx context.Context, request web.UserUpdateUsernameRequest) web.UserResponse
	UpdatePassword(ctx context.Context, request web.UserUpdatePasswordRequest)
	FindById(ctx context.Context, userId int) web.UserResponse
	Delete(ctx context.Context, userId int)
}
