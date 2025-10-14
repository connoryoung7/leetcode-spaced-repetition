package services

import (
	"context"
	"leetcode-spaced-repetition/repositories"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepository repositories.UserRepository
}

type newUser struct {
	email    string `validate:"email"`
	password string `v`
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (a AuthService) Login(ctx context.Context, email string, password string) (bool, error) {
	passwordHash, err := a.userRepository.GetPasswordHashByEmail(ctx, email)
	if err != nil {
		return false, err
	}

	if passwordHash == nil {
		return false, nil
	}

	result := a.comparePasswords(*passwordHash, []byte(password))
	return result, nil
}

func (a AuthService) Logout() {

}

func (a AuthService) RegisterUser(ctx context.Context, email string, password string) error {
	hash, err := a.hashAndSaltPassword(password)
	if err != nil {
		return err
	}

	return a.userRepository.CreateUser(ctx, email, hash)
}

func (a AuthService) hashAndSaltPassword(plainPassword string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.MinCost)
	if err != nil {
		return "", nil
	}

	return string(hash), nil
}

func (a AuthService) comparePasswords(hashedPassword string, plainPassword []byte) bool {
	byteHash := []byte(hashedPassword)

	if err := bcrypt.CompareHashAndPassword(byteHash, plainPassword); err != nil {
		return false
	}

	return true
}
