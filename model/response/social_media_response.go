package response

import "time"

type CreateSocialMediaRespone struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserID         int       `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
}

type GetSocialMediaByIdRespone struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserID         int       `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt		time.Time`json:"updated_at"`
}

type UpdateSocialMediaRespone struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserID         int       `json:"user_id"`
	UpdatedAt		time.Time`json:"updated_at"`
}

type DeleteSocialMediaRespone struct {
	Message   string    `json:"message"`
}