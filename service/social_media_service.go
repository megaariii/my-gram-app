package service

import (
	"context"
	"my-gram/model/domain"
)

type SocialMediaService interface {
	CreateSocialMedia(ctx context.Context, id string, sm domain.SocialMediaInput) (*domain.SocialMedia, error)
	GetAllSocialMedia(ctx context.Context) ([]*domain.SocialMedia, error)
	GetSocialMediaById(ctx context.Context, id string) (*domain.SocialMedia, error)
	UpdateSocialMedia(ctx context.Context, id string, sm domain.SocialMediaInput) (*domain.SocialMedia, error)
	DeleteSocialMedia(ctx context.Context, id string) error
}