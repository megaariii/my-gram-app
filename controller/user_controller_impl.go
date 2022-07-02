package controller

import (
	"my-gram/helper"
	"my-gram/middleware"
	"my-gram/model/domain"
	"my-gram/model/response"
	"my-gram/service"
	"net/http"
	"strconv"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl {
		UserService: userService,
	}
}

func (uc *UserControllerImpl) Register(writer http.ResponseWriter, request *http.Request) {
	var user domain.User
	helper.ReadFromRequestBody(request, &user)

	newRegister, errRegister := uc.UserService.Register(request.Context(), user)
	if errRegister != nil {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	registerRespone := response.RegisterRespone {
		User_id:  newRegister.ID,
		Email:    newRegister.Email,
		Username: newRegister.Username,
		Age:      newRegister.Age,
	}

	webResponse := response.WebResponse {
		Code: 201,
		Status: "Created",
		Data: registerRespone,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (uc *UserControllerImpl) Login(writer http.ResponseWriter, request *http.Request) {
	var login domain.UserLogin
	helper.ReadFromRequestBody(request, &login)

	errValidate := helper.CheckEmpty(login.Email, login.Password)
	if errValidate != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	user, errLogin := uc.UserService.Login(request.Context(), login)
	if errLogin != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	id := strconv.Itoa(user.ID)
	
	token, errToken := helper.GenerateToken(id)
	if errToken != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	userToken := response.UserToken {
		Token: token,
	}

	webResponse := response.WebResponse {
		Code: 200,
		Status: "OK",
		Data: userToken,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (uc *UserControllerImpl) GetUserById(writer http.ResponseWriter, request *http.Request) {
	user := middleware.ForContext(request.Context())
	id := strconv.Itoa(user.ID)

	userId, err := uc.UserService.GetUserById(request.Context(), id)
	if err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	userById := response.GetUserById {
		ID: userId.ID,
		Username: userId.Username,
		Email: userId.Email,
		Age: userId.Age,
	}
	
	webResponse := response.WebResponse {
		Code: 200,
		Status: "OK",
		Data: userById,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (uc *UserControllerImpl) Update(writer http.ResponseWriter, request *http.Request) {
	user := middleware.ForContext(request.Context())

	var login domain.UserLogin
	helper.ReadFromRequestBody(request, &login)

	id := strconv.Itoa(user.ID)

	userUpdate, errUpdate := uc.UserService.Update(request.Context(), id, login)
	if errUpdate != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusInternalServerError)
	}

	newUserUpdate := response.UserUpdate {
		ID: 		userUpdate.ID,
		Username: 	userUpdate.Username,
		Email: 		userUpdate.Email,
		Age:		userUpdate.Age,
		UpdatedAt: 	userUpdate.UpdatedAt,
	}

	webResponse := response.WebResponse {
		Code: 200,
		Status: "OK",
		Data: newUserUpdate,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (uc *UserControllerImpl) Delete(writer http.ResponseWriter, request *http.Request) {
	user := middleware.ForContext(request.Context())

	id := strconv.Itoa(user.ID)

	err := uc.UserService.Delete(request.Context(), id)

	if err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	userDelete := response.UserDelete {
		Message: "Your account has been successfully deleted",
	}

	webResponse := response.WebResponse {
		Code: 200,
		Status: "OK",
		Data: userDelete,
	}

	helper.WriteToResponseBody(writer, webResponse)
}