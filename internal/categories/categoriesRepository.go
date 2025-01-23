package categories

import (
	"NUSphereBackend/internal/common"
	"NUSphereBackend/internal/users"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type CategoriesRepository interface {
	CreateChildCategory(ctx context.Context, category CategoryPublic) error
	GetCategories(ctx context.Context) ([]CategoryPublic, error)
	// GetParentCategories(ctx context.Context) ([]CategoryPublic, error)
	GetCategoryByPublicID(ctx context.Context, publicID string) (Category, error)
	DeleteCategoryByPublicID(ctx context.Context, publicID string) error
}

func CreateChildCategory(pg *common.Postgres, ctx context.Context, category Category, user users.User) error {
	query := `INSERT INTO categories (name, description, public_id, parent_id, created_by) VALUES (@name, @description, @publicID, @parentID, @createdBy)`
	args := pgx.NamedArgs{
		"name":        category.Name,
		"description": category.Description,
		"publicID":    category.PublicID,
		"parentID":    category.ParentID,
		"createdBy":   user.ID,
	}

	_, dbErr := pg.DB.Exec(ctx, query, args)
	if dbErr != nil {
		return fmt.Errorf("unable to insert row: %w", dbErr)
	}
	return nil
}

// API use
func GetCategories(pg *common.Postgres, ctx context.Context) ([]CategoryPublic, error) {
	query := "SELECT categories.public_id, categories.name, categories.description, categories.parent_id, users.public_id as \"users.public_id\" FROM categories " +
		"INNER JOIN users ON categories.created_by = users.id "
		// "WHERE parent_id = $1"
	rows, err := pg.DB.Query(ctx, query)

	if err != nil {
		return nil, fmt.Errorf("unable to query categories: %w", err)
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[CategoryPublic])
}

// internal use
func GetCategoryByPublicID(pg *common.Postgres, ctx context.Context, publicID string) (Category, error) {
	query := "SELECT * FROM categories WHERE categories.public_id = $1"
	row, err := pg.DB.Query(ctx, query, publicID)
	if err != nil {
		return Category{}, fmt.Errorf("unable to get category: %w", err)
	}
	defer row.Close()

	return pgx.CollectOneRow(row, pgx.RowToStructByName[Category])
}

// // API use
// func GetParentCategories(pg *common.Postgres, ctx context.Context) ([]ParentCategoryPublic, error) {
// 	query := "SELECT categories.public_id, categories.name, categories.description, users.public_id as \"users.public_id\" FROM categories " +
// 		"INNER JOIN users ON categories.created_by = users.id " +
// 		"WHERE parent_id = $1"
// 	rows, err := pg.DB.Query(ctx, query, "PARENT")
// 	if err != nil {
// 		return nil, fmt.Errorf("unable to query categories: %w", err)
// 	}
// 	defer rows.Close()

// 	return pgx.CollectRows(rows, pgx.RowToStructByName[ParentCategoryPublic])
// }

// func DeleteCategoryByPublicID(pg *common.Postgres, ctx context.Context, publicID string) error {
// 	query := "DELETE FROM categories WHERE public_id = $1"
// 	_, err := pg.DB.Exec(ctx, query, publicID)
// 	if err != nil {
// 		return fmt.Errorf("unable to delete category: %w", err)
// 	}
// 	return nil
// }

func EditCategoryByPublicID(pg *common.Postgres, ctx context.Context, publicID string, category Category) error {
	query := "UPDATE categories SET name = @name, description = @description WHERE public_id = @publicID"
	args := pgx.NamedArgs{
		"name":        category.Name,
		"description": category.Description,
		"publicID":    publicID,
	}

	_, err := pg.DB.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to update category: %w", err)
	}
	return nil
}
