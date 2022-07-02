package controller

import (
	"my-gram/helper"
	"my-gram/middleware"
	"my-gram/model/domain"
	"my-gram/model/response"
	"my-gram/service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type PhotoControllerImpl struct {
	PhotoService service.PhotoService
}

func NewPhotoController(photoService service.PhotoService) PhotoController {
	return &PhotoControllerImpl{
		PhotoService: photoService,
	}
}

func (pc *PhotoControllerImpl) CreatePhoto(writer http.ResponseWriter, request *http.Request) {
	user := middleware.ForContext(request.Context())
	id := strconv.Itoa(user.ID)

	var photo domain.Photo
	helper.ReadFromRequestBody(request, &photo)

	newPhoto, errCreate := pc.PhotoService.CreatePhoto(request.Context(), id, photo)

	if errCreate != nil {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	createPhotoResponso := response.CreatePhotoRespone {
		ID:  newPhoto.ID,
		Title: newPhoto.Title,
		Caption: newPhoto.Caption,
		PhotoUrl: newPhoto.PhotoUrl,
		UserID: newPhoto.UserID,
		CreatedAt: newPhoto.CreatedAt,
	}

	webResponse := response.WebResponse{
		Code:   201,
		Status: "Created",
		Data:   createPhotoResponso,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (pc *PhotoControllerImpl) GetPhotos(writer http.ResponseWriter, request *http.Request) {
	photos, err := pc.PhotoService.GetPhotos()
	if err != nil {
		writer.Write([]byte(err.Error()))
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   photos,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (pc *PhotoControllerImpl) GetPhotoById(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id := params["id"]

	photoId, errGetById := pc.PhotoService.GetPhotoById(request.Context(), id)

	if errGetById != nil {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	photoById := response.GetPhotoByIdRespone {
		ID: photoId.ID,
		Title: photoId.Title,
		Caption: photoId.Caption,
		PhotoUrl: photoId.PhotoUrl,
		UserID: photoId.UserID,
		CreatedAt: photoId.CreatedAt,
		UpdatedAt: photoId.UpdatedAt,
	}

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   photoById,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (pc *PhotoControllerImpl) UpdatePhoto(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id := params["id"]

	var photo domain.Photo
	helper.ReadFromRequestBody(request, &photo)

	updatedPhoto, errUpdate := pc.PhotoService.UpdatePhoto(request.Context(), id, photo)

	if errUpdate != nil {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	newPhoto := response.UpdatePhotoRespone {
		ID: updatedPhoto.ID,
		Title: updatedPhoto.Title,
		Caption: updatedPhoto.Caption,
		PhotoUrl: updatedPhoto.PhotoUrl,
		UserID: updatedPhoto.UserID,
		UpdatedAt: updatedPhoto.UpdatedAt,
	}

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   newPhoto,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (pc *PhotoControllerImpl) DeletePhoto(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id := params["id"]

	err := pc.PhotoService.DeletePhoto(request.Context(), id)

	if err != nil {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	photoDelete := response.PhotoDelete {
		Message: "Your photo has been successfully deleted",
	}

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   photoDelete,
	}

	helper.WriteToResponseBody(writer, webResponse)
}