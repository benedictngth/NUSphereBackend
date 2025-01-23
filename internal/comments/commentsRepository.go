package comments

import (
	"context"
	"fmt"
	"time"

	"log"

	"github.com/jackc/pgx/v5"

	"NUSphereBackend/internal/common"
	"NUSphereBackend/internal/posts"
	"NUSphereBackend/internal/users"
)

type CommentRepository interface {
	CreateComment(ctx context.Context, comment CommentPublic) error
	GetCommentsByPostID(ctx context.Context) ([]CommentPublic, error)
	GetCommentByPublicID(ctx context.Context, publicID string) (CommentPublic, error)
	EditCommentByPublicID(ctx context.Context, publicID string, comment CommentPublic) error
	DeleteCommentByPublicID(ctx context.Context, publicID string) error
}

func CreateComment(pg *common.Postgres, ctx context.Context, comment CommentPublic) error {
	//get user ID and post ID by public id
	user, userErr := users.GetUserByPublicID(pg, ctx, comment.UserID)
	post, postErr := posts.GetPostByPublicID(pg, ctx, comment.PostID)
	// log.Printf("user: %v", user)
	// log.Printf("post: %v", post)
	if userErr != nil {
		return fmt.Errorf("unable to get user: %w", userErr)
	}
	if postErr != nil {
		return fmt.Errorf("unable to get post: %w", postErr)
	}

	query := `INSERT INTO comments (content, public_id, user_id, post_id) 
	VALUES (@content, @PublicID, @userID, @postID)`
	args := pgx.NamedArgs{
		"content":  comment.Content,
		"PublicID": comment.ID, //public nanoid !database id
		"userID":   user.ID,    //database id
		"postID":   post.ID,    //database id
	}
	// log.Printf("args: %v", args)
	_, dbErr := pg.DB.Exec(ctx, query, args)
	if dbErr != nil {
		return fmt.Errorf("unable to insert row: %w", dbErr)
	}
	return nil
}

func GetCommentsByPostID(pg *common.Postgres, ctx context.Context, postID string) ([]CommentPublic, error) {
	query := "SELECT comments.public_id as \"comments.public_id\", comments.content, comments.created_at, comments.updated_at, users.public_id as \"users.public_id\", posts.public_id as \"posts.public_id\" " +
		"FROM comments " +
		"INNER JOIN users ON comments.user_id = users.id " +
		"INNER JOIN posts ON comments.post_id = posts.id " +
		"WHERE posts.public_id = $1"
	rows, err := pg.DB.Query(ctx, query, postID)
	if err != nil {
		return nil, fmt.Errorf("unable to query comments: %w", err)
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[CommentPublic])
}

func GetCommentByPublicID(pg *common.Postgres, ctx context.Context, publicID string) (CommentPublic, error) {
	query := "SELECT comments.public_id as \"comments.public_id\", comments.content, comments.created_at, comments.updated_at, users.public_id as \"users.public_id\", posts.public_id as \"posts.public_id\" " +
		"FROM comments " +
		"INNER JOIN users ON comments.user_id = users.id " +
		"INNER JOIN posts ON comments.post_id = posts.id " +
		"WHERE comments.public_id = $1"
	row, err := pg.DB.Query(ctx, query, publicID)
	if err != nil {
		return CommentPublic{}, fmt.Errorf("unable to query comment: %w", err)
	}
	defer row.Close()

	return pgx.CollectOneRow(row, pgx.RowToStructByName[CommentPublic])
}

func EditCommentByPublicID(pg *common.Postgres, ctx context.Context, publicID, content string) error {
	query := `UPDATE comments SET content=@content, updated_at = @updatedAt WHERE public_id = @publicID`
	log.Printf("publicID: %v", publicID)
	args := pgx.NamedArgs{
		"content":   content,
		"updatedAt": time.Now().Format(time.RFC3339),
		"publicID":  publicID,
	}
	result, dbErr := pg.DB.Exec(ctx, query, args)
	if dbErr != nil {
		return fmt.Errorf("unable to update row: %w", dbErr)
	}
	//check if any rows were affected
	if result.RowsAffected() == 0 {
		return fmt.Errorf("no rows affected")
	}

	return nil
}

func DeleteCommentByPublicID(pg *common.Postgres, ctx context.Context, publicID string) error {
	query := `DELETE FROM comments WHERE public_id = $1`
	result, dbErr := pg.DB.Exec(ctx, query, publicID)

	if dbErr != nil {
		return fmt.Errorf("unable to delete row: %w", dbErr)
	}
	// log.Printf("Rows affected: %d", result.RowsAffected())
	if result.RowsAffected() == 0 {
		return fmt.Errorf("no rows affected")
	}
	return nil
}
