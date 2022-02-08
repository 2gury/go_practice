package session

import "go_practice/9_clean_arch_db/internal/models"

type SessionRepository interface {
	Create(usrId uint64) (*models.Session, error)
	Get(sessValue string) (*models.Session, error)
	Delete(sessValue string) (error)
}
