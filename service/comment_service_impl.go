package service

import (
	"context"
	"database/sql"
	"errors"
	"log"
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

func (cs *CommentServiceImpl) AddComment(ctx context.Context, id string, input domain.CommentInput) (*domain.Comment, error) {
	if input.Message == "" {
		return nil, errors.New("message cannot be empty")
	}

	var comment domain.Comment

	tx, err	:= cs.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	newComment, errCreate := cs.CommentRepository.AddComment(ctx, tx, id, comment)

	if errCreate != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	return newComment, nil
}

func (cs *CommentServiceImpl) GetAllComment(ctx context.Context) ([]*domain.Comment, error) {
	tx, err	:= cs.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	allComment, errGetAll := cs.CommentRepository.GetAllComment(ctx, tx)

	if errGetAll != nil {
		log.Fatal(err.Error())
		return nil, errGetAll
	}

	return allComment, nil
}

func (cs *CommentServiceImpl) GetCommentById(ctx context.Context, id string) (*domain.Comment, error) {
	tx, err	:= cs.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	getById, errGetById := cs.CommentRepository.GetCommentById(ctx, tx, id)

	if errGetById != nil {
		log.Fatal(err.Error())
		return nil, errGetById
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
		log.Fatal(err.Error())
		return nil, errUpdate
	}

	return updateComment, nil
}

func (cs *CommentServiceImpl) DeleteComment(ctx context.Context, id string) error {
	tx, err	:= cs.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	
	errDelete := cs.CommentRepository.DeleteComment(ctx, tx, id)

	if errDelete != nil {
		log.Fatal(err.Error())
		return errDelete
	}

	return nil
}
