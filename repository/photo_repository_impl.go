package repository

import (
	"context"
	"database/sql"
	"fmt"
	"my-gram/model/domain"
)

type PhotoRepositoryImpl struct {
}

func NewPhotoRepository() PhotoRepository {
	return &PhotoRepositoryImpl{}
}

func (pr *PhotoRepositoryImpl) CreatePhoto(ctx context.Context, tx *sql.Tx, id string, photo domain.Photo) (*domain.Photo, error) {
	SQL := "insert into photos(title, caption, photo_url, user_id, created_at) values($1, $2, $3, $4, now())"
	_, errRow := tx.ExecContext(ctx, SQL, photo.Title, photo.Caption, photo.PhotoUrl, id)

	if errRow != nil {
		fmt.Println("Create Photo Repository Error: " + errRow.Error())
		return nil, errRow
	}

	return &photo, nil
}

func (pr *PhotoRepositoryImpl) GetPhotos(tx *sql.Tx) ([]*domain.Photo, error) {
	row, errRow := tx.Query(
	`
	Select p.id, p.title, p.caption, p.photo_url, p.user_id, p.created_at, u.email, u.username
	from photos p join users u 
	on p.user_id = u.id
	`)

	if errRow != nil {
		fmt.Println("Query Get Photo Error: " + errRow.Error())
		return nil, errRow
	}

	defer row.Close()
	
	var photos []*domain.Photo

	for row.Next() {
		var photo domain.Photo
		err := row.Scan(
			&photo.ID, &photo.Title, &photo.Caption, &photo.PhotoUrl, &photo.UserID, &photo.CreatedAt, &photo.User.Email, &photo.User.Username,
		)

		if err != nil {
			fmt.Println("Get All Photo", err.Error())
			return nil, err
		}

		photos = append(photos, &photo)
	}

	return photos, nil
}

func (pr *PhotoRepositoryImpl) GetPhotoById(ctx context.Context, tx *sql.Tx, id string) (*domain.Photo, error) {
	var photo domain.Photo

	SQL := "select id, title, caption, photo_url, user_id from photos where id = $1"
	row := tx.QueryRowContext(ctx, SQL, id)
	err := row.Scan(&photo.ID, &photo.Title, &photo.Caption, &photo.PhotoUrl, &photo.UserID)
	if err != nil {
		fmt.Println("Get Photo By ID", err.Error())
			return nil, err
	}

	return &photo, nil
}

func (pr *PhotoRepositoryImpl) UpdatePhoto(ctx context.Context, tx *sql.Tx, id string, photo domain.Photo) (*domain.Photo, error) {
	SQL := "update photos set title = $1, caption = $2, photo_url = $3, updated_at = now() where id = $4"
	_, errRow := tx.ExecContext(ctx, SQL, photo.Title, photo.Caption, photo.PhotoUrl, id)

	if errRow != nil {
		fmt.Println("Query Update Photo Error", errRow.Error())
		return nil, errRow
	}

	return &photo, nil

}

func (pr *PhotoRepositoryImpl) DeletePhoto(ctx context.Context, tx *sql.Tx, id string) error {
	SQL := "delete from photos where id = $1"

	_, errRow := tx.ExecContext(ctx, SQL, id)

	if errRow != nil {
		fmt.Println("Query Delete Photo Error", errRow.Error())
		return errRow
	}

	return nil
}

