package usecases

import (
	"go_practice/9_clean_arch_db/internal/consts"
	"go_practice/9_clean_arch_db/internal/helpers/errors"
	"go_practice/9_clean_arch_db/internal/models"
	"go_practice/9_clean_arch_db/internal/session"
	"go_practice/9_clean_arch_db/tools/password"
)

type SessionUsecase struct {
	sessionRep session.SessionRepository
}

func NewSessionUsecase(rep session.SessionRepository) session.SessionUsecase {
	return &SessionUsecase{
		sessionRep: rep,
	}
}

func (u *SessionUsecase) Create(userId uint64) (*models.Session, *errors.Error) {
	sess := models.NewSession(userId)
	sess.Value = password.GetMD5Hash(sess.Value)

	err := u.sessionRep.Create(sess)
	if err != nil {
		return nil, errors.Get(consts.CodeInternalError)
	}

	return sess, nil
}

func (u *SessionUsecase) Check(sessValue string) (*models.Session, *errors.Error) {
	sess, err := u.sessionRep.Get(sessValue)
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
