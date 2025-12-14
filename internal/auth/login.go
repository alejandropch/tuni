package auth

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	repo Repository
}

func NewLoginService(repo Repository) *LoginService {
	return &LoginService{repo: repo}
}

func (s *LoginService) Login(ctx context.Context, email string, password string) (User, error) {
	user, err := s.repo.FindByEmail(ctx, email)
	if err != nil {
		return User{}, errors.New("Invalid Credentials")
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return User{}, errors.New("Invalid Credentials")
	}

	return user, nil

}
