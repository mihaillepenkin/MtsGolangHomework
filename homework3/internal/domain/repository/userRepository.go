package repository

import (
	"homework3/internal/domain/entity"
)

type UserRepositoryInterface interface {
	GetUser(phoneNumber string) (*entity.User, error)
	SaveUser(user *entity.User) error
}