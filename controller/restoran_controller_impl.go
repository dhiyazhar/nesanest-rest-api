package controller

import (
	"log/slog"
	"nesanest-rest-api/helper"
	"nesanest-rest-api/model/web"
	"nesanest-rest-api/service"
	"net/http"
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
	err := request.ParseMultipartForm(10 << 20)
	helper.PanicIfError(err)

	restoranCreateRequest := web.RestoranCreateRequest{
		Name:        request.FormValue("name"),
		Description: request.FormValue("description"),
		Address:     request.FormValue("address"),
	}

	uploadDir := "./static/restoran_images"
	filePath, err := helper.SaveUploadedFile(request, "image", uploadDir)
	if err != nil {
		if err == http.ErrMissingFile {
			helper.PanicIfError(err)
		} else {
			slog.Error("failed to save uploaded file", "error", err)
			return
		}
	} else {
		restoranCreateRequest.ImageUrl = filePath
	}

	slog.Info("Created new RestoranRequest", restoranCreateRequest.Name, restoranCreateRequest.Description)

	restoranResponse := controller.RestoranService.Create(request.Context(), restoranCreateRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusCreated,
		Status: "OK",
		Data:   restoranResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *RestoranControllerImpl) Update(writer http.ResponseWriter, request *http.Request) {
	id := request.Context().Value("restoran_id").(int)

	restoranUpdateRequest := web.RestoranUpdateRequest{}
	helper.ReadFromRequestBody(request, &restoranUpdateRequest)

	restoranUpdateRequest.Id = id

	restoranResponse := controller.RestoranService.Update(request.Context(), restoranUpdateRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   restoranResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *RestoranControllerImpl) Delete(writer http.ResponseWriter, request *http.Request) {
	id := request.Context().Value("restoran_id").(int)

	controller.RestoranService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *RestoranControllerImpl) FindById(writer http.ResponseWriter, request *http.Request) {
	id := request.Context().Value("restoran_id").(int)

	restoranResponse := controller.RestoranService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   restoranResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *RestoranControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request) {
	restoranResponses := controller.RestoranService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   restoranResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
