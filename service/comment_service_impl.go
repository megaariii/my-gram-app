package service

import (
	"context"
	"database/sql"
	"errors"
	"my-gram/exception"
	"my-gram/helper"
	"my-gram/model/domain"
	"my-gram/repository"
)

type CommentServiceImpl struct {
	CommentRepository repository.CommentRepository
	DB                *sql.DB
}

func NewCommentService(commentRepository repository.CommentRepository, DB *sql.DB) CommentService {
	return &CommentServiceImpl{
		CommentRepository: commentRepository,
		DB:                DB,
	}
}

func (cs *CommentServiceImpl) AddComment(ctx context.Context, id string, comment domain.Comment) (*domain.Comment, error) {
	if comment.Message == "" {
		return nil, errors.New("message cannot be empty")
	}

	tx, err	:= cs.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	newComment, errCreate := cs.CommentRepository.AddComment(ctx, tx, id, comment)
	helper.PanicIfError(errCreate)

	return newComment, nil
}

func (cs *CommentServiceImpl) GetAllComment(ctx context.Context) ([]*domain.Comment, error) {
	tx, err	:= cs.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	allComment, errGetAll := cs.CommentRepository.GetAllComment(ctx, tx)
	helper.PanicIfError(errGetAll)

	return allComment, nil
}

func (cs *CommentServiceImpl) GetCommentById(ctx context.Context, id string) (*domain.Comment, error) {
	tx, err	:= cs.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	getById, errGetById := cs.CommentRepository.GetCommentById(ctx, tx, id)
	if errGetById != nil {
		panic(exception.NewNotFoundError(errGetById.Error()))
	}
	
	return getById, nil

}

func (cs *CommentServiceImpl) UpdateComment(ctx context.Context, id string, input domain.CommentInput) (*domain.Comment, error) {
	if input.Message == "" {
		return nil, errors.New("message cannot be empty")
	}

	tx, err	:= cs.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	var comment domain.Comment

	comment.Message = input.Message

	updateComment, errUpdate := cs.CommentRepository.UpdateComment(ctx, tx ,id, comment)
	if errUpdate != nil {
		panic(exception.NewNotFoundError(errUpdate.Error()))
	}
	

	return updateComment, nil
}

func (cs *CommentServiceImpl) DeleteComment(ctx context.Context, id string) error {
	tx, err	:= cs.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	
	errDelete := cs.CommentRepository.DeleteComment(ctx, tx, id)
	if errDelete != nil {
		panic(exception.NewNotFoundError(errDelete.Error()))
	}

	return nil
}
