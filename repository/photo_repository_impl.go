package repository

import (
	"context"
	"database/sql"
	"my-gram/helper"
	"my-gram/model/domain"
)

type PhotoRepositoryImpl struct {
}

func NewPhotoRepository() PhotoRepository {
	return &PhotoRepositoryImpl{}
}

func (pr *PhotoRepositoryImpl) CreatePhoto(ctx context.Context, tx *sql.Tx, id string, photo domain.Photo) (*domain.Photo, error) {
	SQL := "insert into photos(title, caption, photo_url, user_id, created_at) values(?, ?, ?, ?, now())"
	result, errRow := tx.ExecContext(ctx, SQL, photo.Title, photo.Caption, photo.PhotoUrl, id)
	helper.PanicIfError(errRow)

	photoId, err := result.LastInsertId()
	helper.PanicIfError(err)

	photo.ID = int(photoId)

	return &photo, nil
}

func (pr *PhotoRepositoryImpl) GetPhotos(tx *sql.Tx) ([]*domain.Photo, error) {
	row, errRow := tx.Query(
	`
	Select p.id, p.title, p.caption, p.photo_url, p.user_id, u.email, u.username
	from photos p join users u 
	on p.user_id = u.id
	`)
	helper.PanicIfError(errRow)

	defer row.Close()
	
	var photos []*domain.Photo

	for row.Next() {
		var photo domain.Photo
		err := row.Scan(
			&photo.ID, &photo.Title, &photo.Caption, &photo.PhotoUrl, &photo.UserID, &photo.User.Email, &photo.User.Username,
		)
		helper.PanicIfError(err)

		photos = append(photos, &photo)
	}

	return photos, nil
}

func (pr *PhotoRepositoryImpl) GetPhotoById(ctx context.Context, tx *sql.Tx, id string) (*domain.Photo, error) {
	var photo domain.Photo

	SQL := "select id, title, caption, photo_url, user_id from photos where id = ?"
	row := tx.QueryRowContext(ctx, SQL, id)
	err := row.Scan(&photo.ID, &photo.Title, &photo.Caption, &photo.PhotoUrl, &photo.UserID)
	helper.PanicIfError(err)

	return &photo, nil
}

func (pr *PhotoRepositoryImpl) UpdatePhoto(ctx context.Context, tx *sql.Tx, id string, photo domain.Photo) (*domain.Photo, error) {
	SQL := "update photos set title = ?, caption = ?, photo_url = ?, updated_at = now() where id = ?"
	result, errRow := tx.ExecContext(ctx, SQL, photo.Title, photo.Caption, photo.PhotoUrl, id)
	helper.PanicIfError(errRow)

	photoId, err := result.LastInsertId()
	helper.PanicIfError(err)

	photo.ID = int(photoId)

	return &photo, nil
}

func (pr *PhotoRepositoryImpl) DeletePhoto(ctx context.Context, tx *sql.Tx, id string) error {
	SQL := "delete from photos where id = ?"
	_, errRow := tx.ExecContext(ctx, SQL, id)
	helper.PanicIfError(errRow)

	return nil
}

