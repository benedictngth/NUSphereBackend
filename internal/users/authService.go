package users

import (
	"context"
	"errors"
	"log"

	"github.com/benedictngth/NUSphereBackend/internal/common"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

// define the AuthService interface with the methods that the service will implement
type AuthService interface {
	Register(ctx context.Context, username, password string) error
	Login(ctx context.Context, username, password string) (string, error)
	GetUsers(ctx context.Context) ([]UserPublic, error)
}

type authService struct {
	jwtSecret string
}

func NewAuthService(jwtSecret string) *authService {
	return &authService{jwtSecret: jwtSecret}
}

//implement the AuthService interface methods and calls respective repository methods

func (s *authService) Register(ctx context.Context, username, password string) error {
	hash, passwordErr := common.HashPassword(password)
	if passwordErr != nil {
		return passwordErr
	}
	nanoid, nanoidErr := gonanoid.New()
	if nanoidErr != nil {
		return nanoidErr
	}

	user := User{
		Username:     username,
		PasswordHash: hash,
		PublicID:     nanoid,
	}

	registerErr := CreateUser(common.GetDB(), ctx, user)
	if registerErr != nil {
		log.Printf("%v", registerErr)
		return registerErr
	}
	// else if registerErr == nil {

	// }
	return nil
}

func (s *authService) Login(ctx context.Context, username, password string) (string, error) {
	user, err := GetUserByUsername(common.GetDB(), ctx, username)

	if err != nil {
		log.Printf("login error: %v", err)
		return "", errors.New(INVALID_USERNAME_ERROR)
	}

	if !common.CheckPasswordHash(password, user.PasswordHash) {
		return "", errors.New(INVALID_PASSWORD_ERROR)
	}

	token, err := common.GenerateJWT(user.PublicID, user.Username, s.jwtSecret)
	if err != nil {
		log.Printf("unable to generate token")
		return "", err
	}
	return token, nil

}

func (s *authService) GetUsers(ctx context.Context) ([]UserPublic, error) {
	return GetUsers(common.GetDB(), ctx)
}
