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
		panic(exception.NewBadRequestError(errRegister.Error()))
	}

	registerRespone := response.RegisterRespone {
		User_id:  newRegister.ID,
		Email:    newRegister.Email,
		Username: newRegister.Username,
		Age:      newRegister.Age,
	}

	webResponse := response.WebResponse {
		Code: http.StatusCreated,
		Status: "Created",
		Data: registerRespone,
	}

	helper.WriteToResponseBody(writer, http.StatusCreated, webResponse)
}

func (uc *UserControllerImpl) Login(writer http.ResponseWriter, request *http.Request) {
	var login domain.UserLogin
	helper.ReadFromRequestBody(request, &login)

	errValidate := helper.CheckEmpty(login.Email, login.Password)
	if errValidate != nil {
		panic(exception.NewBadRequestError(errValidate.Error()))
	}

	user, errLogin := uc.UserService.Login(request.Context(), login)
	if errLogin != nil {
		panic(exception.NewBadRequestError(errLogin.Error()))
	}

	id := strconv.Itoa(user.ID)
	
	token, errToken := helper.GenerateToken(id)
	if errToken != nil {
		panic(exception.NewBadRequestError(errToken.Error()))
	}

	userToken := response.UserToken {
		Token: token,
	}

	webResponse := response.WebResponse {
		Code: http.StatusOK,
		Status: "OK",
		Data: userToken,
	}

	helper.WriteToResponseBody(writer, http.StatusOK, webResponse)
}

func (uc *UserControllerImpl) GetUserById(writer http.ResponseWriter, request *http.Request) {
	user := middleware.ForContext(request.Context())
	id := strconv.Itoa(user.ID)

	userId, err := uc.UserService.GetUserById(request.Context(), id)
	if err != nil {
		panic(exception.NewBadRequestError(err.Error()))
	}

	userById := response.GetUserById {
		ID: userId.ID,
		Username: userId.Username,
		Email: userId.Email,
		Age: userId.Age,
	}
	
	webResponse := response.WebResponse {
		Code: http.StatusOK,
		Status: "OK",
		Data: userById,
	}

	helper.WriteToResponseBody(writer, http.StatusOK, webResponse)
}

func (uc *UserControllerImpl) Update(writer http.ResponseWriter, request *http.Request) {
	user := middleware.ForContext(request.Context())

	var login domain.UserLogin
	helper.ReadFromRequestBody(request, &login)

	id := strconv.Itoa(user.ID)

	userUpdate, errUpdate := uc.UserService.Update(request.Context(), id, login)
	if errUpdate != nil {
		panic(exception.NewBadRequestError(errUpdate.Error()))
	}

	newUserUpdate := response.UserUpdate {
		ID: 		userUpdate.ID,
		Username: 	userUpdate.Username,
		Email: 		userUpdate.Email,
		Age:		userUpdate.Age,
		UpdatedAt: 	userUpdate.UpdatedAt,
	}

	webResponse := response.WebResponse {
		Code: http.StatusOK,
		Status: "OK",
		Data: newUserUpdate,
	}

	helper.WriteToResponseBody(writer, http.StatusOK, webResponse)
}

func (uc *UserControllerImpl) Delete(writer http.ResponseWriter, request *http.Request) {
	user := middleware.ForContext(request.Context())

	id := strconv.Itoa(user.ID)

	err := uc.UserService.Delete(request.Context(), id)
	if err != nil {
		panic(exception.NewBadRequestError(err.Error()))
	}

	userDelete := response.UserDelete {
		Message: "Your account has been successfully deleted",
	}

	webResponse := response.WebResponse {
		Code: http.StatusOK,
		Status: "OK",
		Data: userDelete,
	}

	helper.WriteToResponseBody(writer, http.StatusOK, webResponse)
}