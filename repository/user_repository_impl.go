package repository

import (
	"context"
	"database/sql"
	"errors"
	"my-gram/app"
	"my-gram/helper"
	"my-gram/model/domain"
	"time"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository  {
	return &UserRepositoryImpl{}
}

func (ur *UserRepositoryImpl) Register(ctx context.Context, tx *sql.Tx, user domain.User) (*domain.User, error) {
	pass, errHash := helper.HashPass(user.Password)
	helper.PanicIfError(errHash)
	user.Password = pass

	SQL := "INSERT INTO users (username, email, password, age, created_at) VALUES ($1, $2, $3, $4, now()) RETURNING id"
	_, err := tx.ExecContext(ctx, SQL, user.Username, user.Email, user.Password, user.Age)
	helper.PanicIfError(err)

	return &user, nil
}

func (ur *UserRepositoryImpl) Login(ctx context.Context, tx *sql.Tx, email string) (*domain.User, error) {
	SQL := "SELECT id, username, email, password, age FROM users WHERE email = $1"
	rows, err := tx.QueryContext(ctx, SQL, email)
	helper.PanicIfError(err)

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Age)
		helper.PanicIfError(err)
		return &user, nil
	} else {
		return &user, errors.New("user is not found")
	}
}

func (ur *UserRepositoryImpl) GetUserById(ctx context.Context, tx *sql.Tx, id string) (*domain.User, error) {
	SQL := "SELECT id, username, email, password, age FROM users WHERE id = $1"
	rows, err := tx.QueryContext(ctx, SQL, id)
	helper.PanicIfError(err)

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Age)
		helper.PanicIfError(err)
		return &user, nil
	} else {
		return &user, errors.New("user is not found")
	}
}

func (ur *UserRepositoryImpl) Update(ctx context.Context, user domain.User) (*domain.User, error) {
	SQL := `UPDATE users SET email = $1, username = $2, updated_at = $3 WHERE id = $4`
	_, errRow := app.Db.Exec(SQL, user.Email, user.Username, time.Now(), user.ID)
	helper.PanicIfError(errRow)

	return &user, nil
}

func (ur *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id string) error {
	SQL := "DELETE FROM users WHERE id = $1"
	_, err := tx.ExecContext(ctx, SQL, id)
	helper.PanicIfError(err)
	
	return nil
}
