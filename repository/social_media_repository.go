package repository

import (
	"context"
	"database/sql"
	"my-gram/model/domain"
)

type SocialMediaRepository interface {
	CreateSocialMedia(ctx context.Context, tx *sql.Tx, id string, photo domain.SocialMedia) (*domain.SocialMedia, error)
	GetAllSocialMedia(ctx context.Context, tx *sql.Tx) ([]*domain.SocialMedia, error)
	GetSocialMediaById(ctx context.Context, tx *sql.Tx, id string) (*domain.SocialMedia, error)
	UpdateSocialMedia(ctx context.Context, tx *sql.Tx, id string, photo domain.SocialMedia) (*domain.SocialMedia, error)
	DeleteSocialMedia(ctx context.Context, tx *sql.Tx, id string) error
}