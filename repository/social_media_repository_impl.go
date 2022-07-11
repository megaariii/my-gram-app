package repository

import (
	"context"
	"database/sql"
	"my-gram/helper"
	"my-gram/model/domain"
)

type SocialMediaRepositoryImpl struct {
}

func NewSocialMediaRepository() SocialMediaRepository {
	return &SocialMediaRepositoryImpl{}
}

func (smr *SocialMediaRepositoryImpl) CreateSocialMedia(ctx context.Context, tx *sql.Tx, id string, socialMedia domain.SocialMedia) (*domain.SocialMedia, error) {
	SQL := "insert into social_media(name, social_media_url, user_id, created_at) values(?, ?, ?, now())"
	result, err := tx.ExecContext(ctx, SQL, socialMedia.Name, socialMedia.SocialMediaUrl, id)
	helper.PanicIfError(err)

	socmedId, err := result.LastInsertId()
	helper.PanicIfError(err)

	socialMedia.ID = int(socmedId)

	return &socialMedia, nil
}

func (smr *SocialMediaRepositoryImpl) GetAllSocialMedia(ctx context.Context, tx *sql.Tx) ([]*domain.SocialMedia, error) {
	var socialMedias []*domain.SocialMedia
	row, err := tx.QueryContext(ctx,
		`select sm.id, sm.name, sm.social_media_url, sm.user_id, us.id, us.username
		from social_media sm
		join users us
		on sm.user_id = us.id;`)
	helper.PanicIfError(err)

	defer row.Close()

	for row.Next() {
		var socialMedia domain.SocialMedia

		err := row.Scan(
			&socialMedia.ID, &socialMedia.Name, &socialMedia.SocialMediaUrl, &socialMedia.UserID, &socialMedia.User.ID,
			&socialMedia.User.Username,
		)
		helper.PanicIfError(err)

		socialMedias = append(socialMedias, &socialMedia)
	}

	return socialMedias, nil
}

func (smr *SocialMediaRepositoryImpl) GetSocialMediaById(ctx context.Context, tx *sql.Tx, id string) (*domain.SocialMedia, error) {
	var sm domain.SocialMedia
	SQL := "select id, name, social_media_url, user_id, created_at from social_media where id = ?"
	row := tx.QueryRowContext(ctx, SQL, id)
	err := row.Scan(&sm.ID, &sm.Name, &sm.SocialMediaUrl, &sm.UserID, &sm.CreatedAt)
	helper.PanicIfError(err)

	return &sm, nil
}

func (smr *SocialMediaRepositoryImpl) UpdateSocialMedia(ctx context.Context, tx *sql.Tx, id string, socialMedia domain.SocialMedia) (*domain.SocialMedia, error) {
	SQL := "update social_media set name = ?, social_media_url = ?, updated_at = now() where id = ?"
	result, err := tx.ExecContext(ctx, SQL, socialMedia.Name, socialMedia.SocialMediaUrl, id)
	helper.PanicIfError(err)

	socmedId, err := result.LastInsertId()
	helper.PanicIfError(err)

	socialMedia.ID = int(socmedId)

	return &socialMedia, nil
}

func (smr *SocialMediaRepositoryImpl) DeleteSocialMedia(ctx context.Context, tx *sql.Tx, id string) error {
	SQL := "delete from social_media where id = ?"
	_, err := tx.ExecContext(ctx, SQL, id)
	helper.PanicIfError(err)

	return nil
}
