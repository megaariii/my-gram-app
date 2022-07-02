package service

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"my-gram/exception"
	"my-gram/helper"
	"my-gram/model/domain"
	"my-gram/repository"
)

type PhotoServiceImpl struct {
	PhotoRepository repository.PhotoRepository
	CommentRepository repository.CommentRepository
	DB              *sql.DB
}

func NewPhotoService(photoRepository repository.PhotoRepository, commentRepository repository.CommentRepository, DB *sql.DB) PhotoService {
	return &PhotoServiceImpl{
		PhotoRepository: photoRepository,
		CommentRepository: commentRepository,
		DB:              DB,
	}
}

func (ps *PhotoServiceImpl) CreatePhoto(ctx context.Context, id string, photo domain.Photo) (*domain.Photo, error) {
	if photo.Title == "" {
		return nil, errors.New("photo title cannot be empty")
	}
	if photo.PhotoUrl == "" {
		return nil, errors.New("photo url cannot be empty")
	}

	tx, err	:= ps.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	newPhoto, errCreate := ps.PhotoRepository.CreatePhoto(ctx, tx, id, photo)
	helper.PanicIfError(errCreate)

	return newPhoto, nil
}

func (ps *PhotoServiceImpl) GetPhotos() ([]*domain.Photo, error) {
	tx, err	:= ps.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	allPhotos, errGet := ps.PhotoRepository.GetPhotos(tx)

	if errGet != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	return allPhotos, nil
}

func (ps *PhotoServiceImpl) GetPhotoById(ctx context.Context, id string) (*domain.Photo, error) {
	tx, err	:= ps.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	photoById, errGetById := ps.PhotoRepository.GetPhotoById(ctx, tx, id)
	if errGetById != nil {
		panic(exception.NewNotFoundError(errGetById.Error()))
	}

	return photoById, nil
}

func (ps *PhotoServiceImpl) UpdatePhoto(ctx context.Context, id string, photo domain.Photo) (*domain.Photo, error) {
	if photo.Title == "" {
		return nil, errors.New("photo title cannot be empty")
	}
	if photo.PhotoUrl == "" {
		return nil, errors.New("photo url cannot be empty")
	}
	
	tx, err	:= ps.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	updatedPhoto, errUpdate := ps.PhotoRepository.UpdatePhoto(ctx, tx, id, photo)
	if errUpdate != nil {
		panic(exception.NewNotFoundError(errUpdate.Error()))
	}


	return updatedPhoto, nil
}

func (ps *PhotoServiceImpl) DeletePhoto(ctx context.Context, id string) error {
	tx, err	:= ps.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	
	errDelete := ps.PhotoRepository.DeletePhoto(ctx, tx, id)
	if errDelete != nil {
		panic(exception.NewNotFoundError(errDelete.Error()))
	}


	return nil
}
