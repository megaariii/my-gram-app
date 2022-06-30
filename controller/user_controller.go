package controller

import (
	"net/http"
)

type UserController interface {
	Register(writer http.ResponseWriter, request *http.Request)
	Login(writer http.ResponseWriter, request *http.Request)
	GetUserById(writer http.ResponseWriter, request *http.Request)
	Update(writer http.ResponseWriter, request *http.Request)
	Delete(writer http.ResponseWriter, request *http.Request)
}