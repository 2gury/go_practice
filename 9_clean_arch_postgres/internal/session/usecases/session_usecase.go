package usecases

import (
	"go_practice/9_clean_arch_db/internal/consts"
	"go_practice/9_clean_arch_db/internal/helpers/errors"
	"go_practice/9_clean_arch_db/internal/models"
	"go_practice/9_clean_arch_db/internal/session"
)

type SessionUsecase struct {
	sessionRep session.SessionRepository
}

func NewSessionUsecase(rep session.SessionRepository) session.SessionUsecase {
	return &SessionUsecase{
		sessionRep: rep,
	}
}

func (u *SessionUsecase) Create(usr *models.User) (*models.Session, *errors.Error) {
	sess, err := u.sessionRep.Create(usr.Id)
	if err != nil {
		return nil, errors.Get(consts.CodeInternalError)
	}
	return sess, nil
}

func (u *SessionUsecase) Check(sessValue string) (*models.Session, *errors.Error) {
	sess, err := u.sessionRep.Check(sessValue)
	if err != nil {
		return nil, errors.Get(consts.CodeInternalError)
	}
	return sess, nil
}

func (u *SessionUsecase) Delete(sessValue string) *errors.Error {
	err := u.sessionRep.Delete(sessValue)
	if err != nil {
		return errors.Get(consts.CodeInternalError)
	}
	return nil
}
