package session

import (
	"go_practice/9_clean_arch_db/internal/helpers/errors"
	"go_practice/9_clean_arch_db/internal/models"
)

type SessionUsecase interface {
	Create(usr *models.User) (*models.Session, *errors.Error)
	Delete(sessValue string) *errors.Error
	Get(sessValue string) (*models.Session, *errors.Error)
}
