package controller

import (
	"net/http"
)

type PhotoController interface {
	CreatePhoto(writer http.ResponseWriter, request *http.Request)
	GetPhotos(writer http.ResponseWriter, request *http.Request)
	GetPhotoById(writer http.ResponseWriter, request *http.Request)
	UpdatePhoto(writer http.ResponseWriter, request *http.Request)
	DeletePhoto(writer http.ResponseWriter, request *http.Request)
}