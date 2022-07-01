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

type SocialMediaControllerImpl struct {
	SocialMediaService service.SocialMediaService
}

func NewSocialMediaController(socialMediaService service.SocialMediaService) SocialMediaController {
	return &SocialMediaControllerImpl{
		SocialMediaService: socialMediaService,
	}
}

func (smc *SocialMediaControllerImpl) CreateSocialMedia(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	user := middleware.ForContext(ctx)
	id := strconv.Itoa(user.ID)

	var input domain.SocialMediaInput
	errDecode := json.NewDecoder(request.Body).Decode(&input)

	if errDecode != nil {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	newSm, errCreate := smc.SocialMediaService.CreateSocialMedia(ctx, id, input)

	if errCreate != nil {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	newSocialMedia := response.CreateSocialMediaRespone {
		ID: newSm.ID,
		Name: newSm.Name,
		SocialMediaUrl: newSm.SocialMediaUrl,
		UserID: newSm.UserID,
		CreatedAt: newSm.CreatedAt,
	}

	response, _ := json.Marshal(newSocialMedia)
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	writer.Write(response)
}

func (smc *SocialMediaControllerImpl) GetAllSocialMedia(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	socialMedias, errSocialMedias := smc.SocialMediaService.GetAllSocialMedia(ctx)

	if errSocialMedias != nil {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

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

	response, _ := json.Marshal(socialMediasResponse)
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(response)
}

func (smc *SocialMediaControllerImpl) GetSocialMediaById(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	params := mux.Vars(request)
	id := params["id"]

	getById, errById := smc.SocialMediaService.GetSocialMediaById(ctx, id)

	if errById != nil {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	socialMediaById := response.GetSocialMediaByIdRespone {
		ID: getById.ID,
		Name: getById.Name,
		SocialMediaUrl: getById.SocialMediaUrl,
		UserID: getById.UserID,
		CreatedAt: getById.CreatedAt,
		UpdatedAt: getById.UpdatedAt,
	}

	response, _ := json.Marshal(socialMediaById)
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(response)
}

func (smc *SocialMediaControllerImpl) UpdateSocialMedia(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	params := mux.Vars(request)
	id := params["id"]

	var input domain.SocialMediaInput
	errDecode := json.NewDecoder(request.Body).Decode(&input)

	if errDecode != nil {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	updateSm, errUpdateSosmed := smc.SocialMediaService.UpdateSocialMedia(ctx, id, input)

	if errUpdateSosmed != nil {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	newSosialMedia := response.UpdateSocialMediaRespone {
		ID: updateSm.ID,
		Name: updateSm.Name,
		SocialMediaUrl: updateSm.SocialMediaUrl,
		UserID: updateSm.UserID,
		UpdatedAt: updateSm.UpdatedAt,
	}

	response, _ := json.Marshal(newSosialMedia)
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(response)
}

func (smc *SocialMediaControllerImpl) DeleteSocialMedia(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	params := mux.Vars(request)
	id := params["id"]

	errDelete := smc.SocialMediaService.DeleteSocialMedia(ctx, id)

	if errDelete != nil {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	socialMediaDelete := response.DeleteSocialMediaRespone {
		Message: "Your Social Media has been successfully deleted",
	}

	response, _ := json.Marshal(socialMediaDelete)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(response)
}
