package repository

import (
	"context"
	"database/sql"
	"fmt"
	"my-gram/model/domain"
	"time"
)

type CommentRepositoryImpl struct {
}

func NewCommentRepository() CommentRepository {
	return &CommentRepositoryImpl{}
}

func (cr *CommentRepositoryImpl) AddComment(ctx context.Context, tx *sql.Tx, id string, comment domain.Comment) (*domain.Comment, error) {
	SQL := "INSERT INTO comments(message, photo_id, user_id, created_at) VALUES ($1, $2, $3, $4)"
	_, errInsert := tx.ExecContext(ctx, SQL, comment.Message, comment.PhotoID, id, time.Now())

	if errInsert != nil {
		fmt.Println("Query Add Comment Repository Error: " + errInsert.Error())
		return nil, errInsert
	}

	return &comment, nil
}

func (cr *CommentRepositoryImpl) GetAllComment(ctx context.Context, tx *sql.Tx) ([]*domain.Comment, error) {
	SQL := `SELECT c.id, c.user_id, c.photo_id, c.message, c.created_at, c.updated_at, 
	p.id, p.title, p.caption, p.photo_url, p.user_id, u.id, u.username, u.email
	FROM comments c
	LEFT JOIN photos p on c.photo_id = p.id
	LEFT JOIN users u on c.user_id = u.id`

	row, errRow := tx.QueryContext(ctx, SQL)

	if errRow != nil {
		fmt.Println("Query Get All Comment Error", errRow.Error())
		return nil, errRow
	}

	defer row.Close()

	var comments []*domain.Comment

	for row.Next() {
		var comment domain.Comment
		var timeAt sql.NullTime

		err := row.Scan(
			&comment.ID, &comment.UserID, &comment.PhotoID, &comment.Message, &comment.CreatedAt,
			&timeAt, &comment.Photo.ID, &comment.Photo.Title, &comment.Photo.Caption,
			&comment.Photo.PhotoUrl, &comment.Photo.UserID, &comment.User.ID, &comment.User.Username,
			&comment.User.Email,
		)

		if err != nil {
			fmt.Println("Err Ccan Get All", err.Error())
			return nil, err
		}

		comments = append(comments, &comment)
	}

	return comments, nil
}

func (cr *CommentRepositoryImpl) GetCommentById(ctx context.Context, tx *sql.Tx, id string) (*domain.Comment, error) {
	var comment *domain.Comment

	SQL := `SELECT id, message, photo_id, user_id FROM photos where id = $1`
	row, errRow := tx.QueryContext(ctx, SQL, id)

	if errRow != nil {
		fmt.Println("Query Get Comment By Id Error", errRow)
		return nil, errRow
	}

	defer row.Close()

	err := row.Scan(&comment.ID, &comment.Message, &comment.PhotoID, &comment.UserID)

	if err != nil {
		return nil, err
	}

	return comment, nil
}


func (cr *CommentRepositoryImpl) UpdateComment(ctx context.Context, tx *sql.Tx, id string, comment domain.Comment) (*domain.Comment, error) {
	SQL := `UPDATE comments SET message = $1, updated_at = now() WHERE id = $2`

	_, errRow := tx.ExecContext(ctx, SQL, comment.Message, id)

	if errRow != nil {
		fmt.Println("Query Update Comment Error", errRow)
		return nil, errRow
	}

	return &comment, nil
}

func (cr *CommentRepositoryImpl) DeleteComment(ctx context.Context, tx *sql.Tx, id string) error {
	sqlQuery := `DELETE FROM comments WHERE id = $1`
	_, errRow := tx.ExecContext(ctx, sqlQuery, id)

	if errRow != nil {
		fmt.Println("Query Delete Comment Error", errRow)
		return errRow
	}

	return nil
}
