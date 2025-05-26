package main

import (
	"log/slog"
	"nesanest-rest-api/app"
	"nesanest-rest-api/controller"
	"nesanest-rest-api/helper"
	"nesanest-rest-api/middleware"
	"nesanest-rest-api/repository"
	"nesanest-rest-api/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq"
)

func main() {
	db := app.NewDB()
	validate := validator.New()

	restoranRepository := repository.NewRestoranRepository()
	restoranService := service.NewRestoranService(restoranRepository, db, validate)
	restoranController := controller.NewRestoranController(restoranService)

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db)
	userController := controller.NewUserController(userService)

	router := app.NewRouter(restoranController, userController)
	
	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	slog.Info("Starting server", "address", server.Addr)

	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
