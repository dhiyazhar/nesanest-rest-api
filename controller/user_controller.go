package controller

import (
	"net/http"
)

type UserController interface {
	Register(writer http.ResponseWriter, request *http.Request)
	Login(writer http.ResponseWriter, request *http.Request)
	UpdateProfile(writer http.ResponseWriter, request *http.Request)
	UpdatePassword(writer http.ResponseWriter, request *http.Request)
	Delete(writer http.ResponseWriter, request *http.Request, id string)
	ForgotPassword(w http.ResponseWriter, r *http.Request)
	FindById(writer http.ResponseWriter, request *http.Request, id string)}
