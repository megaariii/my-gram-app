package service

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"my-gram/helper"
	"my-gram/model/domain"
	"my-gram/repository"
	"time"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB *sql.DB
}

func NewUserService(userRepository repository.UserRepository, DB *sql.DB) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB: DB,
	}
}

func (us *UserServiceImpl) Register(ctx context.Context, user domain.User) (*domain.User, error) {
	if _, ok := helper.ValidMailAddress(user.Email); !ok {
		return nil, errors.New("email is not valid")
	}
	if user.Email == "" {
		return nil, errors.New("email can't be empty or must be filled")
	}
	if user.Username == "" {
		return nil, errors.New("username can't be empty or must be filled")
	}
	if user.Password == "" {
		return nil, errors.New("password can't be empty or must be filled")
	}
	if len(user.Password) < 6 {
		return nil, errors.New("password must greater than 6 characters")
	}
	if user.Age == 0 {
		return nil, errors.New("age can't be empty or must be filled")
	}
	if user.Age <= 8 {
		return nil, errors.New("age must greater than 8")
	}
	
	tx, err	:= us.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	newUser, err := us.UserRepository.Register(ctx, tx, user)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	
	return newUser, nil
}

func (us *UserServiceImpl) Login(ctx context.Context, login domain.UserLogin) (*domain.User, error) { 
	email := login.Email
	password := login.Password

	tx, err	:= us.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	
	user, err := us.UserRepository.Login(ctx, tx, email)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	
	if !helper.ComparePass(password, user.Password) {
		return nil, errors.New("password must be correct")
	}

	return user, nil
}

func (us *UserServiceImpl) GetUserById(ctx context.Context, id string) (*domain.User, error) { 
	tx, err	:= us.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := us.UserRepository.GetUserById(ctx, tx, id)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	return user, nil
}

func (us *UserServiceImpl) Update(ctx context.Context, id string, user domain.UserLogin) (*domain.User, error) { 
	tx, err	:= us.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	userId, errGetId := us.UserRepository.GetUserById(ctx, tx, id)
	if errGetId != nil {
		log.Fatalln(errGetId.Error())
		return nil, errGetId
	}

	userId.Email = user.Email
	userId.Username = user.Username
	userId.UpdatedAt = time.Now()

	updatedUser, err := us.UserRepository.Update(ctx, tx, *userId)
	if err != nil {
		log.Fatalln(err.Error())
		return nil, err
	}

	return updatedUser, nil
}

func (us *UserServiceImpl) Delete(ctx context.Context, id string) error { 
	tx, err	:= us.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	errDel := us.UserRepository.Delete(ctx, tx, id)
	if errDel != nil {
		log.Fatal(errDel.Error())
		return errDel
	}

	return nil
}