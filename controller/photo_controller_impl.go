package controller

import (
	"my-gram/exception"
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
		panic(exception.NewBadRequestError(errCreate.Error()))
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
		Code:   http.StatusCreated,
		Status: "Created",
		Data:   createPhotoResponso,
	}

	helper.WriteToResponseBody(writer, http.StatusCreated, webResponse)
}

func (pc *PhotoControllerImpl) GetPhotos(writer http.ResponseWriter, request *http.Request) {
	photos, err := pc.PhotoService.GetPhotos()
	helper.PanicIfError(err)

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   photos,
	}

	helper.WriteToResponseBody(writer, http.StatusOK, webResponse)
}

func (pc *PhotoControllerImpl) GetPhotoById(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id := params["id"]

	photoId, errGetById := pc.PhotoService.GetPhotoById(request.Context(), id)
	if errGetById != nil {
		panic(exception.NewBadRequestError(errGetById.Error()))
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
		Code:   http.StatusOK,
		Status: "OK",
		Data:   photoById,
	}

	helper.WriteToResponseBody(writer, http.StatusOK, webResponse)
}

func (pc *PhotoControllerImpl) UpdatePhoto(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id := params["id"]

	var photo domain.Photo
	helper.ReadFromRequestBody(request, &photo)

	updatedPhoto, errUpdate := pc.PhotoService.UpdatePhoto(request.Context(), id, photo)
	if errUpdate != nil {
		panic(exception.NewBadRequestError(errUpdate.Error()))
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
		Code:   http.StatusOK,
		Status: "OK",
		Data:   newPhoto,
	}

	helper.WriteToResponseBody(writer, http.StatusOK, webResponse)
}

func (pc *PhotoControllerImpl) DeletePhoto(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id := params["id"]

	err := pc.PhotoService.DeletePhoto(request.Context(), id)
	if err != nil {
		panic(exception.NewBadRequestError(err.Error()))
	}

	photoDelete := response.PhotoDelete {
		Message: "Your photo has been successfully deleted",
	}

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   photoDelete,
	}

	helper.WriteToResponseBody(writer, http.StatusOK, webResponse)
}