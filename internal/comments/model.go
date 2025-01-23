package comments

import "time"

type CommentPublic struct {
	//ID is nanoid not primary key
	ID        string    `db:"comments.public_id"`
	Content   string    `db:"content"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	UserID    string    `db:"users.public_id"`
	PostID    string    `db:"posts.public_id"`
}

type NewCommentRequest struct {
	Content string `json:"Content" binding:"required"`
	UserID  string `json:"UserID" binding:"required"`
	//public ID
	PostID string `json:"PostID" binding:"required"`
}

type EditCommentRequest struct {
	Content string `json:"Content" binding:"required"`
}
