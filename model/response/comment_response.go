package response

import "time"

type CreateCommentRespone struct {
	ID        int       `json:"id"`
	Message   string    `json:"message"`
	PhotoID   int    	`json:"photo_id"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type GetCommentByIdRespone struct {
	ID        int       `json:"id"`
	Message   string    `json:"message"`
	PhotoID   int    	`json:"photo_id"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time	`json:"updated_at"`
}

type UpdateCommentRespone struct {
	ID        int       `json:"id"`
	Message   string    `json:"message"`
	PhotoID   int    	`json:"photo_id"`
	UserID    int       `json:"user_id"`
	UpdatedAt time.Time	`json:"updated_at"`
}

type DeleteCommentRespone struct {
	Message   string    `json:"message"`
}