package comments

import (
	"context"
	"log"

	"github.com/benedictngth/NUSphereBackend/internal/common"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

type CommentsService interface {
	CreateComment(ctx context.Context, Content, UserID, PostID string) error
	GetCommentsByPostID(ctx context.Context, PostID string) ([]CommentPublic, error)
	GetCommentByPublicID(ctx context.Context, publicID string) (CommentPublic, error)
	EditCommentByPublicID(ctx context.Context, publicID, content string) error
	DeleteCommentByPublicID(ctx context.Context, publicID string) error
}

type commentsService struct {
}

func NewCommentsService() *commentsService {
	return &commentsService{}
}

func (s *commentsService) CreateComment(c context.Context, Content, UserID, PostID string) error {
	//generate a new random comment id
	nanoid, err := gonanoid.New()
	if err != nil {
		log.Print(err)
		return err
	}
	comment := CommentPublic{
		ID:      nanoid,
		Content: Content,
		UserID:  UserID,
		PostID:  PostID,
	}
	return CreateComment(common.GetDB(), c, comment)
}

func (s *commentsService) GetCommentsByPostID(c context.Context, postID string) ([]CommentPublic, error) {
	return GetCommentsByPostID(common.GetDB(), c, postID)
}

func (s *commentsService) GetCommentByPublicID(c context.Context, publicID string) (CommentPublic, error) {
	return GetCommentByPublicID(common.GetDB(), c, publicID)
}

func (s *commentsService) EditCommentByPublicID(c context.Context, publicID, content string) error {
	return EditCommentByPublicID(common.GetDB(), c, publicID, content)
}

func (s *commentsService) DeleteCommentByPublicID(c context.Context, publicID string) error {
	return DeleteCommentByPublicID(common.GetDB(), c, publicID)
}
