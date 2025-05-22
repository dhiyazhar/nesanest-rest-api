package controller

import (
	"log/slog"
	"nesanest-rest-api/helper"
	"nesanest-rest-api/model/web"
	"nesanest-rest-api/service"
	"net/http"
	"strconv"
)

type RestoranControllerImpl struct {
	RestoranService service.RestoranService
}

func NewRestoranController(restoranService service.RestoranService) RestoranController {
	return &RestoranControllerImpl{
		RestoranService: restoranService,
	}
}

func (controller *RestoranControllerImpl) Create(writer http.ResponseWriter, request *http.Request) {
	restoranCreateRequest := web.RestoranCreateRequest{}
	helper.ReadFromRequestBody(request, &restoranCreateRequest)
	slog.Info("New HTTP request", slog.Any("HTTP", request))
	slog.Info("Created new RestoranRequest", restoranCreateRequest.Name, restoranCreateRequest.Description)

	restoranResponse := controller.RestoranService.Create(request.Context(), restoranCreateRequest)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   restoranResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *RestoranControllerImpl) Update(writer http.ResponseWriter, request *http.Request, id string) {
	restoranUpdateRequest := web.RestoranUpdateRequest{}
	helper.ReadFromRequestBody(request, &restoranUpdateRequest)

	idInt, err := strconv.Atoi(id)
	helper.PanicIfError(err)

	restoranUpdateRequest.Id = idInt

	restoranResponse := controller.RestoranService.Update(request.Context(), restoranUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   restoranResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *RestoranControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, id string) {
	idInt, err := strconv.Atoi(id)
	helper.PanicIfError(err)

	controller.RestoranService.Delete(request.Context(), idInt)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *RestoranControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, id string) {
	idInt, err := strconv.Atoi(id)
	helper.PanicIfError(err)

	restoranResponse := controller.RestoranService.FindById(request.Context(), idInt)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   restoranResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *RestoranControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request) {
	restoranResponses := controller.RestoranService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   restoranResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
