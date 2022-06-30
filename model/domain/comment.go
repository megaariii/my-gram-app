package domain

import "time"

type Comment struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	PhotoID   int       `json:"photo_id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      User      `json:"user"`
	Photo     Photo     `json:"photo"`
}

type CommentInput struct {
	PhotoID int    `json:"photo_id"`
	Message string `json:"message"`
}