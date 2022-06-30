package controller

import (
	"net/http"
)

type CommentController interface {
	AddComment(writer http.ResponseWriter, request *http.Request)
	GetAllComment(writer http.ResponseWriter, request *http.Request)
	GetCommentById(writer http.ResponseWriter, request *http.Request)
	UpdateComment(writer http.ResponseWriter, request *http.Request)
	DeleteComment(writer http.ResponseWriter, request *http.Request)
}