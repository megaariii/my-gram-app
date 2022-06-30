package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"my-gram/helper"
	"my-gram/model/domain"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository  {
	return &UserRepositoryImpl{}
}

func (ur *UserRepositoryImpl) Register(ctx context.Context, tx *sql.Tx, user domain.User) (*domain.User, error) {
	pass, errHash := helper.HashPass(user.Password)
	if errHash != nil {
		fmt.Println("Hash Password Error: " + errHash.Error())
		return nil, errHash
	}
	user.Password = pass

	SQL := "insert into users (username, email, password, age, created_at) values ($1, $2, $3, $4, now())"
	_, err := tx.ExecContext(ctx, SQL, user.Username, user.Email, user.Password, user.Age)
	
	if err != nil {
		fmt.Println("Register Repository Error: " + err.Error())
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepositoryImpl) Login(ctx context.Context, tx *sql.Tx, email string) (*domain.User, error) {
	SQL := "select id, username, email, password, age from users where email = $1"
	rows, err := tx.QueryContext(ctx, SQL, email)
	if err != nil {
		fmt.Println("Query Context: " + err.Error())
		return nil, err
	}

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Age)
		if err != nil {
			return nil, err
		}
		return &user, nil
	} else {
		return &user, errors.New("user is not found")
	}
}

func (ur *UserRepositoryImpl) GetUserById(ctx context.Context, tx *sql.Tx, id string) (*domain.User, error) {
	SQL := "select id, username, email, password, age from users where id = $1"
	rows, err := tx.QueryContext(ctx, SQL, id)
	if err != nil {
		fmt.Println("Query Context: " + err.Error())
		return nil, err
	}

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Age)
		if err != nil {
			return nil, err
		}
		return &user, nil
	} else {
		return &user, errors.New("user is not found")
	}
}

func (ur *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user domain.User) (*domain.User, error) {
	SQL := "update users set email = $1, username = $2 where id = $3"
	_, err := tx.ExecContext(ctx, SQL, user.Email, user.Username, user.ID)
	if err != nil {
		fmt.Println("Update Repository Error: " + err.Error())
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id string) error {
	SQL := "delete from users where id = $1"
	_, err := tx.ExecContext(ctx, SQL, id)
	if err != nil {
		return err
	}
	
	return nil
}
