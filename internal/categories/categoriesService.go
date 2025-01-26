package categories

import (
	"context"
	"fmt"
	"log"

	"github.com/benedictngth/NUSphereBackend/internal/common"
	"github.com/benedictngth/NUSphereBackend/internal/users"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

type CategoriesService interface {
	CreateChildCategory(ctx context.Context, Name, Description, parentID, created_by string) error
	GetCategories(ctx context.Context) ([]CategoryPublic, error)
}

type categoriesService struct {
}

func NewCategoriesService() *categoriesService {
	return &categoriesService{}
}

func (s *categoriesService) CreateChildCategory(ctx context.Context, Name, Description, parentID, CreatedBy string) error {
	nanoid, err := gonanoid.New()
	if err != nil {
		log.Print(err)
		return err
	}
	category := Category{
		PublicID:    nanoid,
		Name:        Name,
		Description: Description,
		ParentID:    parentID,
	}
	log.Printf("Created by: %s", CreatedBy)
	user, err := users.GetUserByPublicID(common.GetDB(), ctx, CreatedBy)
	if err != nil {
		log.Print(err)
		return fmt.Errorf("unable to get user: %w", err)
	}
	return CreateChildCategory(common.GetDB(), ctx, category, user)

}

func (s *categoriesService) GetCategories(c context.Context) ([]CategoryPublic, error) {
	return GetCategories(common.GetDB(), c)
}

// func (s *categoriesService) DeleteCategoryByPublicID(c context.Context, publicID string) error {
// 	return DeleteCategoryByPublicID(common.GetDB(), c, publicID)
// }
