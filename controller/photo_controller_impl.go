package controller

import (
	"encoding/json"
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
	ctx := request.Context()
	user := middleware.ForContext(ctx)
	id := strconv.Itoa(user.ID)

	var photo domain.Photo
	errDecode := json.NewDecoder(request.Body).Decode(&photo)

	if errDecode != nil {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	newPhoto, errCreate := pc.PhotoService.CreatePhoto(ctx, id, photo)

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

	response, _ := json.Marshal(createPhotoResponso)
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	writer.Write(response)
}

// GetPhotos implements PhotoController
func (pc *PhotoControllerImpl) GetPhotos(writer http.ResponseWriter, request *http.Request) {
	photos, err := pc.PhotoService.GetPhotos()
	if err != nil {
		writer.Write([]byte(err.Error()))
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, _ := json.Marshal(photos)
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(response)
}

// GetPhotoById implements PhotoController
func (pc *PhotoControllerImpl) GetPhotoById(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	params := mux.Vars(request)
	id := params["id"]

	photoId, errGetById := pc.PhotoService.GetPhotoById(ctx, id)

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

	response, _ := json.Marshal(photoById)
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(response)
}
// UpdatePhoto implements PhotoController
func (pc *PhotoControllerImpl) UpdatePhoto(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	params := mux.Vars(request)
	id := params["id"]

	var photo domain.Photo
	errDecode := json.NewDecoder(request.Body).Decode(&photo)

	if errDecode != nil {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	updatedPhoto, errUpdate := pc.PhotoService.UpdatePhoto(ctx, id, photo)

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

	response, _ := json.Marshal(newPhoto)
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(response)
}

// DeletePhoto implements PhotoController
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

	response, _ := json.Marshal(photoDelete)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(response)
}