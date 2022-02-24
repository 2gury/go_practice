package usecases

import (
	"context"
	"go_practice/9_clean_arch_db/internal/consts"
	"go_practice/9_clean_arch_db/internal/helpers/errors"
	"go_practice/9_clean_arch_db/internal/models"
	"go_practice/9_clean_arch_db/internal/session"
	"go_practice/9_clean_arch_db/internal/session/delivery/grpc"
	"go_practice/9_clean_arch_db/tools/password"
)

type SessionUsecase struct {
	sessionService grpc.SessionServiceClient
}

func NewSessionUsecase(sessService grpc.SessionServiceClient) session.SessionUsecase {
	return &SessionUsecase{
		sessionService: sessService,
	}
}

func (u *SessionUsecase) Create(userId uint64) (*models.Session, *errors.Error) {
	sess := models.NewSession(userId)
	sess.Value = password.GetMD5Hash(sess.Value)

	_, err := u.sessionService.Create(context.Background(), grpc.ModelSessionToGrpc(sess))
	if err != nil {
		return nil, errors.Get(consts.CodeInternalError)
	}

	return sess, nil
}

func (u *SessionUsecase) Check(sessValue string) (*models.Session, *errors.Error) {
	sess, err := u.sessionService.Check(context.Background(), &grpc.SessionValue{Value: sessValue})
	if err != nil {
		return nil, errors.Get(consts.CodeInternalError)
	}

	return grpc.GrpcSessionToModel(sess), nil
}

func (u *SessionUsecase) Delete(sessValue string) *errors.Error {
	_, err := u.sessionService.Delete(context.Background(), &grpc.SessionValue{Value: sessValue})
	if err != nil {
		return errors.Get(consts.CodeInternalError)
	}

	return nil
}
