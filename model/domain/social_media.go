package domain

import "time"

type SocialMedia struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserID         int       `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	User           User      `json:"user"`
}

type SocialMediaInput struct {
	Name           string `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
}
