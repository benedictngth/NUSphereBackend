package posts

import (
	// "errors"
	"NUSphereBackend/internal/common"
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Posts(router *gin.RouterGroup, postsService PostsService) {
	router.POST("/create", CreatePostHandler(postsService))
	router.GET("", GetPostsHandler(postsService))
	router.GET("/:id", GetPostByIDHandler(postsService))
	router.PUT("/edit/:id", EditPostByPublicIDHandler(postsService))
	router.DELETE("/delete/:id", DeletePostByPublicIDHandler(postsService))
}

func CreatePostHandler(postsService PostsService) gin.HandlerFunc {
	return func(c *gin.Context) {
		time.Sleep(1 * time.Second)
		var req NewPostRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": common.INVALID_INPUT})
			return
		}
		//insert post into db
		err := postsService.CreatePost(context.Background(), req.Title, req.Content, req.UserID, req.CategoryID)
		if err != nil {
			c.Error(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": common.INVALID_INPUT})
		}
		c.JSON(http.StatusOK, gin.H{"message": CREATE_POST})
	}
}

func GetPostsHandler(postsService PostsService) gin.HandlerFunc {
	return func(c *gin.Context) {
		time.Sleep(1 * time.Second)
		category := c.Query("category")
		log.Printf("category: %s", category)
		if category == "" {
			//get all posts
			posts, err := postsService.GetPosts(context.Background())
			if err != nil {
				c.Error(err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": NO_POSTS})
				return
			}
			c.JSON(http.StatusOK, posts)
		} else {
			//get posts by category
			posts, err := postsService.GetPostsByCategory(context.Background(), category)
			if err != nil {
				c.Error(err)
				switch err.Error() {
				case NO_CATEGORY:
					c.JSON(http.StatusBadRequest, gin.H{"error": NO_CATEGORY})
				case NO_POSTS:
					c.JSON(http.StatusInternalServerError, gin.H{"error": NO_POSTS})
				}
			} else {
				c.JSON(http.StatusOK, posts)
			}
		}
	}
}

func GetPostByIDHandler(postsService PostsService) gin.HandlerFunc {
	return func(c *gin.Context) {
		time.Sleep(1 * time.Second)
		publicID := c.Param("id")
		//get category by public id
		post, err := postsService.GetPostPublicByPublicID(context.Background(), publicID)
		if err != nil {
			c.Error(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": NO_POST})
			return
		}
		c.JSON(http.StatusOK, post)
	}
}

func EditPostByPublicIDHandler(postsService PostsService) gin.HandlerFunc {
	return func(c *gin.Context) {
		time.Sleep(1 * time.Second)
		publicID := c.Param("id")
		var req EditPostRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": common.INVALID_INPUT})
			return
		}
		//edit post by public id
		err := postsService.EditPostByPublicID(context.Background(), publicID, req.Title, req.Content, req.CategoryID)
		if err != nil {
			c.Error(err)
			switch err.Error() {
			case NO_POSTS_MUTATION:
				c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to edit post"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"message": EDIT_POST})
		}
	}
}

func DeletePostByPublicIDHandler(postsService PostsService) gin.HandlerFunc {
	return func(c *gin.Context) {
		time.Sleep(1 * time.Second)
		publicID := c.Param("id")
		//delete post by public id
		err := postsService.DeletePostByPublicID(context.Background(), publicID)
		if err != nil {
			c.Error(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": NO_POSTS_MUTATION})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": DELETE_POST})
	}
}
