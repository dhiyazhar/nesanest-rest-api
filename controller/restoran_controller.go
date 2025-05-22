package controller

import (
	"net/http"
)

type RestoranController interface {
	Create(writer http.ResponseWriter, request *http.Request)
	Update(writer http.ResponseWriter, request *http.Request, id string)
	Delete(writer http.ResponseWriter, request *http.Request, id string)
	FindById(writer http.ResponseWriter, request *http.Request, id string)
	FindAll(writer http.ResponseWriter, request *http.Request)
}
