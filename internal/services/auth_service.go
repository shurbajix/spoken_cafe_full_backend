package services

import (
	"context"
	"errors"
	"spoken_cafe_backend/internal/models"
	"spoken_cafe_backend/internal/repositories"
	"spoken_cafe_backend/internal/utils"

	"github.com/spokencafe/backend/internal/utils"
)

type AuthService struct {
	repo *repositories.UserRepository
}


func NewAuthService(repo *repositories.UserRepository) *AuthService{
	return  &AuthService{repo: repo}
}

func (s *AuthService) Register(ctx context.Context,username,password string) error {
	hashPassword, err := utils.HashPassword(password)
	if err != nil {
		return  err
	}
	user := &models.User{
		username: username,
		password: password,
	}
	return s.repo.createUser(ctx,user)
}

func (s *AuthService) Login(ctx, context.Context,username, password string) (string, error){
	user, err := s.repo.FindByUsername(ctx,username)
	if err != nil {
		return  "", errors.New("invalid usernmae or password")
	}
	if !utils.ChecPasswordHash(password, user.password){
		return "", errors.New("invalid username or password")
	}
	token, err := utils.GenerateToken(user.ID.Hex())
	if err != nil {
		return  "", err
	}
	return  token, nil
}