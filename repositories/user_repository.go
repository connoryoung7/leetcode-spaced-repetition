package repositories

import (
	"context"

	"leetcode-spaced-repetition/models"
)

type UserRepository interface {
	CreateUser(ctx context.Context, email string, passwordHash string) error
	GetUserByEmail(ctx context.Context, email string) (models.User, error)
	GetPasswordHashByEmail(ctx context.Context, email string) (*string, error)
}
