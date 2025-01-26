package posts

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5"

	// "github.com/jackc/pgx/v5/pgxpool"
	"github.com/benedictngth/NUSphereBackend/internal/categories"
	"github.com/benedictngth/NUSphereBackend/internal/common"
	"github.com/benedictngth/NUSphereBackend/internal/users"
)

type PostRepository interface {
	CreatePost(ctx context.Context, post PostPublic) error
	GetPosts(ctx context.Context) ([]PostPublic, error)
	GetPostPublicByPublicID(ctx context.Context, publicID string) (PostPublic, error)
	GetPostByPublicID(ctx context.Context, publicID string) (Post, error)
	EditPostByPublicID(ctx context.Context, publicID string, post PostPublic) error
	DeletePostByPublicID(ctx context.Context, publicID string) error
	GetPostsByCategory(ctx context.Context, categoryID string) ([]PostPublic, error)
}

func CreatePost(pg *common.Postgres, ctx context.Context, post PostPublic) error {
	//get user and category by public id with ID fields
	user, userErr := users.GetUserByPublicID(pg, ctx, post.UserID)
	category, categoryErr := categories.GetCategoryByPublicID(pg, ctx, post.CategoryID)
	if userErr != nil {
		return fmt.Errorf("unable to get user: %w", userErr)
	}
	if categoryErr != nil {
		return fmt.Errorf("unable to get category: %w", categoryErr)
	}
	query := `INSERT INTO posts (title, content, public_id, user_id, category_id) 
	VALUES (@title, @content, @PublicID, @userID, @categoryID)`
	args := pgx.NamedArgs{
		"title":      post.Title,
		"content":    post.Content,
		"PublicID":   post.ID,     //public nanoid !database id
		"userID":     user.ID,     //database id
		"categoryID": category.ID, //database id
	}
	_, dbErr := pg.DB.Exec(ctx, query, args)
	if dbErr != nil {
		return fmt.Errorf("unable to insert row: %w", dbErr)
	}
	return nil
}

func GetPosts(pg *common.Postgres, ctx context.Context) ([]PostPublic, error) {
	query := "SELECT posts.public_id as \"posts.public_id\", posts.title, posts.content, posts.created_at, posts.updated_at, users.public_id as \"users.public_id\", categories.public_id as \"categories.public_id\" " +
		"FROM posts " +
		"INNER JOIN users ON posts.user_id = users.id " +
		"INNER JOIN categories ON posts.category_id = categories.id"
	rows, err := pg.DB.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("unable to query posts: %w", err)
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[PostPublic])
}

// used for API endpoint WITHOUT database id
func GetPostPublicByPublicID(pg *common.Postgres, ctx context.Context, publicID string) (PostPublic, error) {
	log.Print(publicID)
	query := "SELECT posts.public_id as \"posts.public_id\", posts.title, posts.content, posts.created_at, posts.updated_at, users.public_id as \"users.public_id\", categories.public_id as \"categories.public_id\" FROM posts " +
		"INNER JOIN users ON posts.user_id = users.id " +
		"INNER JOIN categories ON posts.category_id = categories.id " +
		"WHERE posts.public_id = $1"
	rows, err := pg.DB.Query(ctx, query, publicID)
	if err != nil {
		return PostPublic{}, fmt.Errorf("unable to query posts: %w", err)
	}
	defer rows.Close()

	return pgx.CollectOneRow(rows, pgx.RowToStructByName[PostPublic])
}

// get ALL fields for backend usage with database id
func GetPostByPublicID(pg *common.Postgres, ctx context.Context, publicID string) (Post, error) {
	query := "SELECT * FROM posts WHERE public_id = $1"
	row, err := pg.DB.Query(ctx, query, publicID)
	if err != nil {
		return Post{}, fmt.Errorf("unable to query post: %w", err)
	}
	defer row.Close()

	return pgx.CollectOneRow(row, pgx.RowToStructByName[Post])
}

func EditPostPublicByPublicID(pg *common.Postgres, ctx context.Context, publicID string, post PostPublic) error {
	category, categoryErr := categories.GetCategoryByPublicID(pg, ctx, post.CategoryID)
	if categoryErr != nil {
		return fmt.Errorf("unable to get category: %w", categoryErr)
	}
	query := `UPDATE posts SET title = @title, content = @content, updated_at = @updatedAt, category_id = @categoryID WHERE public_id = @postPublicID`
	args := pgx.NamedArgs{
		"title":        post.Title,
		"content":      post.Content,
		"postPublicID": publicID,
		"categoryID":   category.ID,
		"updatedAt":    time.Now().Format(time.RFC3339),
	}
	result, err := pg.DB.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to update row: %w", err)
	}
	// check if any rows were affected
	if result.RowsAffected() == 0 {
		return fmt.Errorf(NO_POSTS_MUTATION)
	}
	return nil
}

func DeletePostByPublicID(pg *common.Postgres, ctx context.Context, publicID string) error {
	query := `DELETE FROM posts WHERE public_id = $1`
	result, err := pg.DB.Exec(ctx, query, publicID)
	if err != nil {
		return fmt.Errorf("unable to delete row: %w", err)
	}
	if result.RowsAffected() == 0 {
		return fmt.Errorf(NO_POSTS_MUTATION)
	}
	return nil
}

func GetPostsByCategory(pg *common.Postgres, ctx context.Context, categoryID string) ([]PostPublic, error) {
	query := "SELECT posts.public_id as \"posts.public_id\", posts.title, posts.content, posts.created_at, posts.updated_at, users.public_id as \"users.public_id\", categories.public_id as \"categories.public_id\" " +
		"FROM posts " +
		"INNER JOIN users ON posts.user_id = users.id " +
		"INNER JOIN categories ON posts.category_id = categories.id " +
		"WHERE categories.public_id = $1"

	result, err := pg.DB.Query(ctx, query, categoryID)
	if err != nil {
		return nil, fmt.Errorf("unable to query posts by category: %w", err)
	}
	defer result.Close()

	return pgx.CollectRows(result, pgx.RowToStructByName[PostPublic])
}
