package user

import "go_practice/9_clean_arch_db/internal/models"

type UserRepository interface {
	SelectById(id uint64) (*models.User, error)
	Insert(usr *models.User) (uint64, error)
	SelectByEmail(email string) (*models.User, error)
	UpdatePassword(usr *models.User) error
	DeleteById(id uint64) error
}
