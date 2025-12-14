package auth

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type RegisterService struct {
	repo Repository
}

func NewRegisterService(repo Repository) *RegisterService {
	return &RegisterService{repo: repo}
}

func (s *RegisterService) Register(ctx context.Context, email string, password string) (User, error) {
	if email == "" || password == "" || len(password) > 72 { // TODO: add more validations to this
		return User{}, errors.New("validation failed")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, err
	}
	user := User{
		Email:    email,
		Password: string(hash),
	}
	return s.repo.Create(ctx, user)

}
