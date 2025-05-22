package app

import (
	"nesanest-rest-api/controller"
	"nesanest-rest-api/exception"
	"nesanest-rest-api/helper"
	"net/http"
	"path"
	"strings"
)

type Router struct {
	restoranController controller.RestoranController
}

func NewRouter(restoranController controller.RestoranController) http.Handler {
	return &Router{
		restoranController: restoranController,
	}
}

func (router *Router) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			exception.ErrorHandler(writer, request, err)
		}
	}()

	cleanPath := path.Clean(request.URL.Path)

	if cleanPath == "/" {
		message := "NesaNest Rest API - 2025"
		helper.WriteToResponseBody(writer, message)

		return
	}

	if cleanPath == "/api/v1/restoran" {
		switch request.Method {
		case http.MethodGet:
			router.restoranController.FindAll(writer, request)
		case http.MethodPost:
			router.restoranController.Create(writer, request)
		default:
			writer.WriteHeader(http.StatusMethodNotAllowed)
		}

		return
	}

	if strings.HasPrefix(cleanPath, "/api/v1/restoran") {
		id := strings.TrimPrefix(cleanPath, "/api/v1/restoran/")
		if id == "" {
			http.Error(writer, "missing restoran ID", http.StatusBadRequest)

			return
		}

		switch request.Method {
		case http.MethodGet:
			router.restoranController.FindById(writer, request, id)
		case http.MethodPut:
			router.restoranController.Update(writer, request, id)
		case http.MethodDelete:
			router.restoranController.Delete(writer, request, id)
		default:
			writer.WriteHeader(http.StatusMethodNotAllowed)
		}

		return
	}

	http.NotFound(writer, request)
}
