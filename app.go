package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/benedictngth/NUSphereBackend/internal/categories"
	"github.com/benedictngth/NUSphereBackend/internal/comments"
	"github.com/benedictngth/NUSphereBackend/internal/common"
	"github.com/benedictngth/NUSphereBackend/internal/config"
	"github.com/benedictngth/NUSphereBackend/internal/posts"
	"github.com/benedictngth/NUSphereBackend/internal/users"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()
	fmt.Println(cfg)

	pool, err := common.NewPG(context.Background(), cfg.DataBaseURL)

	if err != nil {
		log.Fatalf("unable to create connection pool: %v\n", err)
	}

	defer pool.Close()

	// err = pool.DB.Ping(context.Background())
	// if err != nil {
	// 	log.Fatalf("unable to ping database: %v\n", err)
	// }
	// fmt.Println("Ping successful; DB succesfully connected!")
	err = common.RunMigrations(cfg.DataBaseURL)
	if err != nil {
		log.Fatalf("unable to run migrations: %v\n", err)
	}

	fmt.Println("Migration successful")

	//create repository and services structs
	authService := users.NewAuthService(cfg.JWTSecret)
	postService := posts.NewPostsService()
	categoriesService := categories.NewCategoriesService()
	commentsService := comments.NewCommentsService()

	r := gin.Default()
	// Attach log request body and recovery middleware
	r.Use(common.LogRequestBodyMiddleware())
	r.Use(gin.Recovery())

	//configure cors middleware
	r.Use(cors.New(cors.Config{
		//allow origins for npm run dev, npm run service and the deployed frontend
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:4173", "https://www.nusphere.benedictngth.dev"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorisation"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	//routes
	v1 := r.Group("/api")
	users.Users(v1.Group("/users"), authService)

	posts.Posts(v1.Group("/posts"), postService)
	//protected routes with JWT cookie middleware
	v1.Use(users.AuthMiddleware(cfg.JWTSecret))
	users.AuthUsers(v1.Group("/users"), authService)
	users.Profile(v1.Group("/users"))

	comments.Comments(v1.Group("/comments"), commentsService)
	categories.Categories(v1.Group("/categories"), categoriesService)

	port := cfg.Port
	if port == "" {
		port = "8000"
	}
	log.Printf("server is running on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("unable to start server: %v\n", err)
	}
}
