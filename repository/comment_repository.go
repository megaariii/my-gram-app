package repository

import (
	"context"
	"database/sql"
	"my-gram/model/domain"
)

type CommentRepository interface {
	AddComment(ctx context.Context, tx *sql.Tx, id string, comment domain.Comment) (*domain.Comment, error)
	GetAllComment(ctx context.Context, tx *sql.Tx,) ([]*domain.Comment, error)
	GetCommentById(ctx context.Context, tx *sql.Tx, id string) (*domain.Comment, error)
	UpdateComment(ctx context.Context, tx *sql.Tx, id string, comment domain.Comment) (*domain.Comment, error)
	DeleteComment(ctx context.Context, tx *sql.Tx, id string) error
}