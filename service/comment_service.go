package service

import (
	"context"
	"my-gram/model/domain"
)

type CommentService interface {
	AddComment(ctx context.Context, id string, input domain.Comment) (*domain.Comment, error)
	GetAllComment(ctx context.Context) ([]*domain.Comment, error)
	GetCommentById(ctx context.Context, id string) (*domain.Comment, error)
	UpdateComment(ctx context.Context, id string, input domain.CommentInput) (*domain.Comment, error)
	DeleteComment(ctx context.Context, id string) error
}