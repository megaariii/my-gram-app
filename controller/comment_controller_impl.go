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

type CommentControllerImpl struct {
	CommentService service.CommentService
}

func NewCommentController(commentService service.CommentService) CommentController {
	return &CommentControllerImpl{
		CommentService: commentService,
	}
}

func (cc *CommentControllerImpl) AddComment(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	user := middleware.ForContext(ctx)
	id := strconv.Itoa(user.ID)

	var input domain.CommentInput
	errDecode := json.NewDecoder(request.Body).Decode(&input)

	if errDecode != nil {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	newComment, errCreate := cc.CommentService.AddComment(ctx, id, input)

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

	response, _ := json.Marshal(createComment)
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	writer.Write(response)
}

func (cc *CommentControllerImpl) GetAllComment(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	commentAll, errGetAll := cc.CommentService.GetAllComment(ctx)
	if errGetAll != nil {
		writer.Write([]byte(errGetAll.Error()))
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, _ := json.Marshal(commentAll)
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(response)
}

func (cc *CommentControllerImpl) GetCommentById(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	params := mux.Vars(request)
	id := params["id"]

	getById, errById := cc.CommentService.GetCommentById(ctx, id)
	
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

	response, _ := json.Marshal(getCommentById)
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(response)
}

func (cc *CommentControllerImpl) UpdateComment(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	params := mux.Vars(request)
	id := params["id"]

	var input domain.CommentInput
	errDecode := json.NewDecoder(request.Body).Decode(&input)

	if errDecode != nil {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	updatedComment, errUpdateComment := cc.CommentService.UpdateComment(ctx, id, input)

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

	response, _ := json.Marshal(newComment)
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(response)
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

	response, _ := json.Marshal(commentDelete)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(response)
}