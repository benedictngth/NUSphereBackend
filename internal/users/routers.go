package users

//contains the handlers for the user routes and routers

import (
	// "errors"
	"NUSphereBackend/internal/common"
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Users(router *gin.RouterGroup, authService AuthService) {
	router.POST("/register", RegisterHandler(authService))
	router.POST("/login", LoginHandler(authService))
	router.GET("/auth", GetAuthHandler(authService))
	router.GET("", GetUsersHandler(authService))
}

func AuthUsers(router *gin.RouterGroup, authService AuthService) {
	router.GET("/authUser", GetAuthUserHandler(authService))
	router.POST("/logout", LogoutHandler(authService))
}

// used to get the authenticated user with assumption that cookies already set
func GetAuthUserHandler(authService AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, userIdExists := c.Get("user_id")
		username, userNameExists := c.Get("username")
		if !userIdExists || !userNameExists {
			c.JSON(http.StatusInternalServerError, gin.H{"error": AUTH_USER_NOT_FOUND})
			return
		}
		c.JSON(http.StatusOK, gin.H{"id": userId, "username": username})
	}
}

func GetAuthHandler(authService AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := c.Cookie("Authorisation")
		log.Print(err)
		if err != nil {
			c.JSON(http.StatusUnauthorized, AUTH_FAILURE)
			return
		}
		c.JSON(http.StatusOK, AUTH_SUCCESS)
	}
}

func LogoutHandler(authService AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.SetCookie("Authorisation", "", -1, "/", "localhost", false, true)
		c.JSON(http.StatusOK, gin.H{"message": LOGOUT_SUCCESS})
	}
}

func Profile(router *gin.RouterGroup) {
	router.GET("/profile", func(c *gin.Context) {
		userID, userErr := c.Get("user_id")
		if !userErr {
			c.JSON(404, gin.H{"error": NO_USER_PROFILE})
			return
		}
		c.JSON(http.StatusOK, gin.H{"user": userID})
	})
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}

// handler for registering a new user in the system
func RegisterHandler(authService AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req RegisterRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			log.Printf("error: %v", err)
			switch err.Error() {
			case NOT_MIN_LENGTH_ERROR:
				c.JSON(http.StatusBadRequest, gin.H{"error": DUPLICATE_USER})
				return
			default:
				c.JSON(http.StatusBadRequest, gin.H{"error": common.INVALID_INPUT, "details": err.Error()})
				return
			}
		}
		err := authService.Register(context.Background(), req.Username, req.Password)
		if err != nil {
			c.Error(err)
			log.Printf("err.Error() is: %v", err.Error())
			switch err.Error() {
			case DUPLICATE_USER_ERROR:
				c.JSON(http.StatusConflict, gin.H{"error": DUPLICATE_USER})
				return
			case NOT_MIN_LENGTH_ERROR:
				c.JSON(http.StatusBadRequest, gin.H{"error": NOT_MIN_LENGTH})
				return
			default:
				c.JSON(http.StatusInternalServerError, gin.H{"error": REGISTRATION_FAILED, "details": err.Error()})
			}
		}
		c.JSON(http.StatusOK, gin.H{"message": REGISTRATION_SUCCESS})
	}
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// handler for logging in a user calls authService.Login
func LoginHandler(authService AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req LoginRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.Error(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
			return
		}

		token, err := authService.Login(context.Background(), req.Username, req.Password)
		if err != nil {
			c.Error(err)
			switch err.Error() {
			case INVALID_USERNAME_ERROR:
				c.JSON(http.StatusUnauthorized, gin.H{"error": INVALID_USERNAME})
			case INVALID_PASSWORD_ERROR:
				c.JSON(http.StatusUnauthorized, gin.H{"error": INVALID_PASSWORD})
				return
			}
		}
		log.Print("generated token string: ", token)
		//persist token in cookie for 24 hours
		c.SetCookie("Authorisation", fmt.Sprintf("Bearer %v", token), 60*60*24, "/", "nuspherebackend.benedictngth.dev", false, true)
		c.JSON(http.StatusOK, gin.H{"username": req.Username})
	}
}

func GetUsersHandler(authService AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := authService.GetUsers(context.Background())
		if err != nil {
			c.Error(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": USERS_NOT_FOUND})
			return
		}
		c.JSON(http.StatusOK, users)
	}
}
