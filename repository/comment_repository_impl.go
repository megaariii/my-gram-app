package repository

import (
	"context"
	"database/sql"
	"my-gram/helper"
	"my-gram/model/domain"
	"time"
)

type CommentRepositoryImpl struct {
}

func NewCommentRepository() CommentRepository {
	return &CommentRepositoryImpl{}
}

func (cr *CommentRepositoryImpl) AddComment(ctx context.Context, tx *sql.Tx, id string, comment domain.Comment) (*domain.Comment, error) {
	SQL := "INSERT INTO comments(message, photo_id, user_id, created_at) VALUES (?, ?, ?, ?)"
	result, errInsert := tx.ExecContext(ctx, SQL, comment.Message, comment.PhotoID, id, time.Now())
	helper.PanicIfError(errInsert)

	commentId, err := result.LastInsertId()
	helper.PanicIfError(err)

	comment.ID = int(commentId)

	return &comment, nil
}

func (cr *CommentRepositoryImpl) GetAllComment(ctx context.Context, tx *sql.Tx) ([]*domain.Comment, error) {
	SQL := `SELECT c.id, c.user_id, c.photo_id, c.message, 
	p.id, p.title, p.caption, p.photo_url, p.user_id, u.id, u.username, u.email
	FROM comments c
	LEFT JOIN photos p on c.photo_id = p.id
	LEFT JOIN users u on c.user_id = u.id`

	row, errRow := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(errRow)

	defer row.Close()

	var comments []*domain.Comment

	for row.Next() {
		var comment domain.Comment

		err := row.Scan(
			&comment.ID, &comment.UserID, &comment.PhotoID, &comment.Message, 
			&comment.Photo.ID, &comment.Photo.Title, &comment.Photo.Caption,
			&comment.Photo.PhotoUrl, &comment.Photo.UserID, &comment.User.ID, &comment.User.Username,
			&comment.User.Email,
		)

		helper.PanicIfError(err)
		comments = append(comments, &comment)
	}

	return comments, nil
}

func (cr *CommentRepositoryImpl) GetCommentById(ctx context.Context, tx *sql.Tx, id string) (*domain.Comment, error) {
	var comment *domain.Comment

	SQL := `SELECT id, message, photo_id, user_id FROM photos where id = ?`
	row, errRow := tx.QueryContext(ctx, SQL, id)
	helper.PanicIfError(errRow)

	defer row.Close()

	err := row.Scan(&comment.ID, &comment.Message, &comment.PhotoID, &comment.UserID)
	helper.PanicIfError(err)

	return comment, nil
}


func (cr *CommentRepositoryImpl) UpdateComment(ctx context.Context, tx *sql.Tx, id string, comment domain.Comment) (*domain.Comment, error) {
	SQL := `UPDATE comments SET message = ?, updated_at = now() WHERE id = ?`
	result, errRow := tx.ExecContext(ctx, SQL, comment.Message, id)
	helper.PanicIfError(errRow)

	commentId, err := result.LastInsertId()
	helper.PanicIfError(err)

	comment.ID = int(commentId)

	return &comment, nil
}

func (cr *CommentRepositoryImpl) DeleteComment(ctx context.Context, tx *sql.Tx, id string) error {
	sqlQuery := `DELETE FROM comments WHERE id = ?`
	_, errRow := tx.ExecContext(ctx, sqlQuery, id)
	helper.PanicIfError(errRow)

	return nil
}
