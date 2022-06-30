package response

import "time"

type CreatePhotoRespone struct {
	ID        int		`json:"id"`
	Title     string 	`json:"tittle"`
	Caption   string 	`json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserID    int		`json:"user_id"`
	CreatedAt time.Time	`json:"created_at"`
}

type GetPhotosRespone struct {
	ID        int		`json:"id"`
	Title     string 	`json:"tittle"`
	Caption   string 	`json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserID    int		`json:"user_id"`
	CreatedAt time.Time	`json:"created_at"`
	UpdatedAt time.Time	`json:"updated_at"`
	User	struct {}	`json:"user"`
}

type GetPhotoByIdRespone struct {
	ID        int		`json:"id"`
	Title     string 	`json:"tittle"`
	Caption   string 	`json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserID    int		`json:"user_id"`
	CreatedAt time.Time	`json:"created_at"`
	UpdatedAt time.Time	`json:"updated_at"`
}

type UpdatePhotoRespone struct {
	ID        int		`json:"id"`
	Title     string 	`json:"tittle"`
	Caption   string 	`json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserID    int		`json:"user_id"`
	UpdatedAt time.Time	`json:"updated_at"`
}

type PhotoDelete struct {
	Message string `json:"message"`
}