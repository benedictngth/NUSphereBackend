package users

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func UpdateContextUserModel(c *gin.Context, my_user_id, my_username string) {
	var myUserModel User
	c.Set("user_id", my_user_id)
	c.Set("username", my_username)
	c.Set("user_model", myUserModel)
}

func AuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		//get the token from the cookie Authorisation
		tokenString, err := c.Cookie("Authorisation")
		if err != nil || tokenString == "" {
			log.Print("token string: ", tokenString)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "no auth token in auth cookie, error: " + err.Error()})
			return
		}
		//remove the Bearer prefix
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		//parse the token and validate signature and expiration
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("unexpected signing method: %v", token.Header["alg"])})
				return nil, nil
			}
			return []byte(secret), nil
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		}

		// Check the expiration
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token expired"})
			return
		}
		//atach the user_id and username to the context
		userID := (claims["user_id"]).(string)
		username := (claims["username"]).(string)
		UpdateContextUserModel(c, userID, username)

		c.Next()
	}
}
