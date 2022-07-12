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

type SocialMediaControllerImpl struct {
	SocialMediaService service.SocialMediaService
}

func NewSocialMediaController(socialMediaService service.SocialMediaService) SocialMediaController {
	return &SocialMediaControllerImpl{
		SocialMediaService: socialMediaService,
	}
}

func (smc *SocialMediaControllerImpl) CreateSocialMedia(writer http.ResponseWriter, request *http.Request) {
	user := middleware.ForContext(request.Context())
	id := strconv.Itoa(user.ID)

	var input domain.SocialMediaInput
	helper.ReadFromRequestBody(request, &input)

	newSm, errCreate := smc.SocialMediaService.CreateSocialMedia(request.Context(), id, input)
	helper.PanicIfError(errCreate)

	newSocialMedia := response.CreateSocialMediaRespone {
		ID: newSm.ID,
		Name: newSm.Name,
		SocialMediaUrl: newSm.SocialMediaUrl,
		UserID: newSm.UserID,
		CreatedAt: newSm.CreatedAt,
	}

	webResponse := response.WebResponse{
		Code:   http.StatusCreated,
		Status: "Created",
		Data:   newSocialMedia,
	}

	helper.WriteToResponseBody(writer, http.StatusCreated, webResponse)
}

func (smc *SocialMediaControllerImpl) GetAllSocialMedia(writer http.ResponseWriter, request *http.Request) {
	socialMedias, errSocialMedias := smc.SocialMediaService.GetAllSocialMedia(request.Context())
	helper.PanicIfError(errSocialMedias)

	var socialMediasResponse []map[string]interface{}

	for _, val := range socialMedias {
		formatter := map[string]interface{}{
			"id":               val.ID,
			"name":             val.Name,
			"social_media_url": val.SocialMediaUrl,
			"user_id":          val.UserID,
			"created_at":       val.CreatedAt,
			"updated_at":       val.UpdatedAt,
			"user": map[string]interface{}{
				"id":       val.User.ID,
				"username": val.User.Username,
			},
		}
		socialMediasResponse = append(socialMediasResponse, formatter)
	}
	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   socialMediasResponse,
	}

	helper.WriteToResponseBody(writer, http.StatusOK, webResponse)
}

func (smc *SocialMediaControllerImpl) GetSocialMediaById(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id := params["id"]

	getById, errById := smc.SocialMediaService.GetSocialMediaById(request.Context(), id)
	helper.PanicIfError(errById)

	socialMediaById := response.GetSocialMediaByIdRespone {
		ID: getById.ID,
		Name: getById.Name,
		SocialMediaUrl: getById.SocialMediaUrl,
		UserID: getById.UserID,
		CreatedAt: getById.CreatedAt,
		UpdatedAt: getById.UpdatedAt,
	}

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   socialMediaById,
	}

	helper.WriteToResponseBody(writer, http.StatusOK, webResponse)
}

func (smc *SocialMediaControllerImpl) UpdateSocialMedia(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id := params["id"]

	var input domain.SocialMediaInput
	helper.ReadFromRequestBody(request, &input)

	updateSm, errUpdateSosmed := smc.SocialMediaService.UpdateSocialMedia(request.Context(), id, input)
	helper.PanicIfError(errUpdateSosmed)

	newSosialMedia := response.UpdateSocialMediaRespone {
		ID: updateSm.ID,
		Name: updateSm.Name,
		SocialMediaUrl: updateSm.SocialMediaUrl,
		UserID: updateSm.UserID,
		UpdatedAt: updateSm.UpdatedAt,
	}

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   newSosialMedia,
	}

	helper.WriteToResponseBody(writer, http.StatusOK, webResponse)
}

func (smc *SocialMediaControllerImpl) DeleteSocialMedia(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id := params["id"]

	errDelete := smc.SocialMediaService.DeleteSocialMedia(request.Context(), id)
	helper.PanicIfError(errDelete)

	socialMediaDelete := response.DeleteSocialMediaRespone {
		Message: "Your Social Media has been successfully deleted",
	}

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   socialMediaDelete,
	}

	helper.WriteToResponseBody(writer, http.StatusOK, webResponse)
}
