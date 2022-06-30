package controller

import "net/http"

type SocialMediaController interface {
	CreateSocialMedia(writer http.ResponseWriter, request *http.Request)
	GetAllSocialMedia(writer http.ResponseWriter, request *http.Request)
	GetSocialMediaById(writer http.ResponseWriter, request *http.Request)
	UpdateSocialMedia(writer http.ResponseWriter, request *http.Request)
	DeleteSocialMedia(writer http.ResponseWriter, request *http.Request)
}