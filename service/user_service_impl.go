package service

import (
	"context"
	"database/sql"
	"errors"
	"nesanest-rest-api/exception"
	"nesanest-rest-api/helper"
	"nesanest-rest-api/model/domain"
	"nesanest-rest-api/model/web"
	"nesanest-rest-api/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
}

func NewUserService(userRepository repository.UserRepository, db *sql.DB) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             db,
	}
}

func (service *UserServiceImpl) Register(ctx context.Context, request web.UserRegisterRequest) web.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	helper.PanicIfError(err)

	user := domain.User{
		Username:   request.Username,
		Email:      request.Email,
		Password:   string(hashedPassword),
		ProfileImg: "",
	}

	user = service.UserRepository.Save(ctx, tx, user)
	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) Login(ctx context.Context, request web.UserLoginRequest) (web.UserResponse, string) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindByEmail(ctx, tx, request.Email)
	if err != nil {
		panic(errors.New("email tidak ditemukan"))
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		panic(errors.New("password salah"))
	}

	token, err := helper.GenerateJWT(user.Id, user.Email)
	helper.PanicIfError(err)

	return helper.ToUserResponse(user), token
}

func (service *UserServiceImpl) UpdateProfile(ctx context.Context, request web.UserUpdateUsernameRequest) web.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError("User tidak ditemukan"))
	}

	user.Username = request.Username
	user = service.UserRepository.UpdateUsername(ctx, tx, user)
	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) UpdatePassword(ctx context.Context, request web.UserUpdatePasswordRequest) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError("User tidak ditemukan"))
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.OldPassword))
	if err != nil {
		panic(errors.New("password lama salah"))
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.NewPassword), bcrypt.DefaultCost)
	helper.PanicIfError(err)

	user.Password = string(hashedPassword)
	service.UserRepository.UpdatePassword(ctx, tx, user)
}

func (service *UserServiceImpl) FindById(ctx context.Context, userId int) web.UserResponse {
    tx, err := service.DB.Begin()
    helper.PanicIfError(err)
    defer helper.CommitOrRollback(tx)

    user, err := service.UserRepository.FindById(ctx, tx, userId)
    if err != nil {
        panic(exception.NewNotFoundError("user tidak ditemukan"))
    }
    return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) Delete(ctx context.Context, userId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, userId)
	if err != nil {
		panic(exception.NewNotFoundError("user tidak ditemukan"))
	}

	service.UserRepository.Delete(ctx, tx, user)
}

func (service *UserServiceImpl) ForgotPassword(ctx context.Context, request web.UserForgotPasswordRequest) {
    tx, err := service.DB.Begin()
    helper.PanicIfError(err)
    defer helper.CommitOrRollback(tx)

    user, err := service.UserRepository.FindByEmail(ctx, tx, request.Email)
    if err != nil {
        panic(errors.New("email tidak ditemukan"))
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.NewPassword), bcrypt.DefaultCost)
    helper.PanicIfError(err)

    user.Password = string(hashedPassword)
    service.UserRepository.UpdatePassword(ctx, tx, user)
}