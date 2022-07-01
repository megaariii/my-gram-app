package controller

import (
	"encoding/json"
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
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&user)
	if err != nil {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

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

	// webResponse := response.WebResponse {
	// 	Code: http.StatusCreated,
	// 	Status: "Success to Created User",
	// 	Data: registerRespone,
	// }

	response, _ := json.Marshal(registerRespone)
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	writer.Write(response)
}

func (uc *UserControllerImpl) Login(writer http.ResponseWriter, request *http.Request) {
	var login domain.UserLogin
	err := json.NewDecoder(request.Body).Decode(&login)
	if err != nil {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

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

	response, _ := json.Marshal(userToken)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(response)
}

func (uc *UserControllerImpl) GetUserById(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	user := middleware.ForContext(ctx)
	id := strconv.Itoa(user.ID)

	userId, err := uc.UserService.GetUserById(ctx, id)
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
	
	response, _ := json.Marshal(userById)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(response)
}

func (uc *UserControllerImpl) Update(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	user := middleware.ForContext(ctx)

	var login domain.User
	err := json.NewDecoder(request.Body).Decode(&login)
	if err != nil {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	id := strconv.Itoa(user.ID)

	userUpdate, errUpdate := uc.UserService.Update(ctx, id, login)
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

	response, _ := json.Marshal(newUserUpdate)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(response)
}

func (uc *UserControllerImpl) Delete(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	user := middleware.ForContext(ctx)

	id := strconv.Itoa(user.ID)

	err := uc.UserService.Delete(ctx, id)

	if err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	userDelete := response.UserDelete {
		Message: "Your account has been successfully deleted",
	}

	response, _ := json.Marshal(userDelete)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(response)
}