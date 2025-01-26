package posts

import (
	"context"
	"log"

	"github.com/benedictngth/NUSphereBackend/internal/common"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

type PostsService interface {
	CreatePost(ctx context.Context, Title, Content, UserID, CategoryID string) error
	GetPosts(ctx context.Context) ([]PostPublic, error)
	//used by api endpoint to get post by public id without database id
	GetPostPublicByPublicID(ctx context.Context, publicID string) (PostPublic, error)
	GetPostByPublicID(ctx context.Context, publicID string) (Post, error)
	EditPostByPublicID(ctx context.Context, publicID, title, content, categoryID string) error
	DeletePostByPublicID(ctx context.Context, publicID string) error
	GetPostsByCategory(ctx context.Context, categoryID string) ([]PostPublic, error)
}

type postsService struct {
}

func NewPostsService() *postsService {
	return &postsService{}
}

func (s *postsService) CreatePost(c context.Context, Title, Content, UserID, CategoryID string) error {
	//generate a new post id
	nanoid, err := gonanoid.New()
	if err != nil {
		log.Print(err)
		return err
	}
	post := PostPublic{
		ID:         nanoid,
		Title:      Title,
		Content:    Content,
		UserID:     UserID,
		CategoryID: CategoryID,
	}
	return CreatePost(common.GetDB(), c, post)

}

func (s *postsService) GetPosts(c context.Context) ([]PostPublic, error) {
	return GetPosts(common.GetDB(), c)
}

func (s *postsService) GetPostPublicByPublicID(c context.Context, publicID string) (PostPublic, error) {
	return GetPostPublicByPublicID(common.GetDB(), c, publicID)
}

func (s *postsService) GetPostByPublicID(c context.Context, publicID string) (Post, error) {
	return GetPostByPublicID(common.GetDB(), c, publicID)
}

func (s *postsService) EditPostByPublicID(c context.Context, publicID, title, content, categoryID string) error {
	//get post with publicID
	post, err := GetPostPublicByPublicID(common.GetDB(), c, publicID)
	if err != nil {
		return err
	}
	//update title and content
	post.Title = title
	post.Content = content
	post.CategoryID = categoryID
	return EditPostPublicByPublicID(common.GetDB(), c, publicID, post)
}

func (s *postsService) DeletePostByPublicID(c context.Context, publicID string) error {
	return DeletePostByPublicID(common.GetDB(), c, publicID)
}

func (s *postsService) GetPostsByCategory(c context.Context, categoryID string) ([]PostPublic, error) {
	return GetPostsByCategory(common.GetDB(), c, categoryID)
}
