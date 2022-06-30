package repository

import (
	"context"
	"database/sql"
	"fmt"
	"my-gram/model/domain"
)

type SocialMediaRepositoryImpl struct {
}

func NewSocialMediaRepository() SocialMediaRepository {
	return &SocialMediaRepositoryImpl{}
}

func (smr *SocialMediaRepositoryImpl) CreateSocialMedia(ctx context.Context, tx *sql.Tx, id string, socialMedia domain.SocialMedia) (*domain.SocialMedia, error) {
	SQL := "insert into social_medias(name, social_media_url, user_id, created_at) values($1, $2, $3, now())"

	_, err := tx.ExecContext(ctx, SQL, socialMedia.Name, socialMedia.SocialMediaUrl, id)

	if err != nil {
		fmt.Println("Query Add Sosmed Repository Error", err.Error())
		return nil, err
	}

	return &socialMedia, nil
}

func (smr *SocialMediaRepositoryImpl) GetAllSocialMedia(ctx context.Context, tx *sql.Tx) ([]*domain.SocialMedia, error) {
	var socialMedias []*domain.SocialMedia

	row, err := tx.QueryContext(ctx,
		`select sm.id, sm.name, sm.social_media_url, sm.user_id, sm.created_at, sm.updated_at, us.id, us.username
		from social_medias sm
		join users us
		on sm.user_id = us.id;`)

	if err != nil {
		fmt.Println("Query Get All Sosmed Error", err.Error())
		return nil, err
	}

	defer row.Close()

	for row.Next() {
		var socialMedia domain.SocialMedia
		var timeAt sql.NullTime

		err := row.Scan(
			&socialMedia.ID, &socialMedia.Name, &socialMedia.SocialMediaUrl, &socialMedia.UserID,
			&socialMedia.CreatedAt, &timeAt, &socialMedia.User.ID,
			&socialMedia.User.Username,
		)

		if err != nil {
			fmt.Println("errGetAll", err.Error())
			return nil, err
		}

		socialMedias = append(socialMedias, &socialMedia)
	}

	return socialMedias, nil
}

func (smr *SocialMediaRepositoryImpl) GetSocialMediaById(ctx context.Context, tx *sql.Tx, id string) (*domain.SocialMedia, error) {
	SQL := "select id, name, social_media_url, user_id, created_at, updated_at from social_medias where id=$1"
	row := tx.QueryRowContext(ctx, SQL, id)

	var sm domain.SocialMedia

	err := row.Scan(&sm.ID, &sm.Name, &sm.SocialMediaUrl, &sm.UserID, &sm.CreatedAt, &sm.UpdatedAt)

	if err != nil {
		fmt.Println("Query Get Sosmed By Id Error", err.Error())
		return nil, err
	}

	return &sm, nil
}

func (smr *SocialMediaRepositoryImpl) UpdateSocialMedia(ctx context.Context, tx *sql.Tx, id string, socialMedia domain.SocialMedia) (*domain.SocialMedia, error) {
	SQL := "update social_medias set name=$1, social_media_url=$2, updated_at=now() where id=$3"

	_, err := tx.ExecContext(ctx, SQL, socialMedia.Name, socialMedia.SocialMediaUrl, id)

	if err != nil {
		fmt.Println("Query Update Sosmed Error", err.Error())
		return nil, err
	}

	return &socialMedia, nil
}

func (smr *SocialMediaRepositoryImpl) DeleteSocialMedia(ctx context.Context, tx *sql.Tx, id string) error {
	SQL := "delete from social_medias where id = $1"

	_, err := tx.ExecContext(ctx, SQL, id)

	if err != nil {
		fmt.Println("Query Delete Sosmed Error", err.Error())
		return err
	}

	return nil
}
