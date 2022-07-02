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

type CommentControllerImpl struct {
	CommentService service.CommentService
}

func NewCommentController(commentService service.CommentService) CommentController {
	return &CommentControllerImpl{
		CommentService: commentService,
	}
}

func (cc *CommentControllerImpl) AddComment(writer http.ResponseWriter, request *http.Request) {
	user := middleware.ForContext(request.Context())
	id := strconv.Itoa(user.ID)

	var input domain.Comment
	helper.ReadFromRequestBody(request, &input)

	newComment, errCreate := cc.CommentService.AddComment(request.Context(), id, input)

	if errCreate != nil {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	createComment := response.CreateCommentRespone {
		ID: newComment.ID,
		Message: newComment.Message,
		PhotoID: newComment.PhotoID,
		UserID: newComment.UserID,
		CreatedAt: newComment.CreatedAt,
	}

	webResponse := response.WebResponse{
		Code:   201,
		Status: "Created",
		Data:   createComment,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (cc *CommentControllerImpl) GetAllComment(writer http.ResponseWriter, request *http.Request) {
	commentAll, errGetAll := cc.CommentService.GetAllComment(request.Context())
	if errGetAll != nil {
		writer.Write([]byte(errGetAll.Error()))
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   commentAll,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
func (cc *CommentControllerImpl) GetCommentById(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id := params["id"]

	getById, errById := cc.CommentService.GetCommentById(request.Context(), id)
	
	if errById != nil {
		writer.Write([]byte(errById.Error()))
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	getCommentById := response.GetCommentByIdRespone {
		ID: getById.ID,
		Message: getById.Message,
		PhotoID: getById.PhotoID,
		UserID: getById.UserID,
		CreatedAt: getById.CreatedAt,
		UpdatedAt: getById.UpdatedAt,
	}

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   getCommentById,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (cc *CommentControllerImpl) UpdateComment(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id := params["id"]

	var input domain.CommentInput
	helper.ReadFromRequestBody(request, &input)

	updatedComment, errUpdateComment := cc.CommentService.UpdateComment(request.Context(), id, input)

	if errUpdateComment != nil {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	newComment := response.UpdateCommentRespone {
		ID: updatedComment.ID,
		Message: updatedComment.Message,
		PhotoID: updatedComment.PhotoID,
		UserID: updatedComment.UserID,
		UpdatedAt: updatedComment.UpdatedAt,
	}

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   newComment,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (cc *CommentControllerImpl) DeleteComment(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id := params["id"]

	errDelete := cc.CommentService.DeleteComment(request.Context(), id)

	if errDelete != nil {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	commentDelete := response.DeleteCommentRespone {
		Message: "Your comment has been successfully deleted",
	}

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   commentDelete,
	}

	helper.WriteToResponseBody(writer, webResponse)
}