package user

import (
	"go_practice/9_clean_arch_db/internal/helpers/errors"
	"go_practice/9_clean_arch_db/internal/models"
)

type UserUsecase interface {
	GetById(id uint64) (*models.User, *errors.Error)
	GetByEmail(email string) (*models.User, *errors.Error)
	Create(*models.User) (uint64, *errors.Error)
	ComparePasswords(password, repeatedPassword string) *errors.Error
	ComparePasswordAndHash(usr *models.User, password string) (*errors.Error)
	UpdateUserPassword(usr *models.User) *errors.Error
	DeleteUserById(id uint64) *errors.Error
}