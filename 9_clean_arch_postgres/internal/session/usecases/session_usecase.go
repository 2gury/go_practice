package usecases

import (
	"context"
	"go_practice/9_clean_arch_db/internal/helpers/errors"
	"go_practice/9_clean_arch_db/internal/models"
	"go_practice/9_clean_arch_db/internal/session"
	"go_practice/9_clean_arch_db/internal/session/delivery/grpc"
)

type SessionUsecase struct {
	sessionSvc grpc.SessionServiceClient
}

func NewSessionUsecase(svc grpc.SessionServiceClient) session.SessionUsecase {
	return &SessionUsecase{
		sessionSvc: svc,
	}
}

func (u *SessionUsecase) Create(userId uint64) (*models.Session, *errors.Error) {
	sess, err := u.sessionSvc.Create(context.Background(), &grpc.SessionUserIdValue{Value: userId})
	if err != nil {
		return nil, errors.GetCustomError(err)
	}

	return grpc.GrpcSessionToModel(sess), nil
}

func (u *SessionUsecase) Check(sessValue string) (*models.Session, *errors.Error) {
	sess, err := u.sessionSvc.Check(context.Background(), &grpc.SessionValue{Value: sessValue})
	if err != nil {
		return nil, errors.GetCustomError(err)
	}

	return grpc.GrpcSessionToModel(sess), nil
}

func (u *SessionUsecase) Delete(sessValue string) *errors.Error {
	_, err := u.sessionSvc.Delete(context.Background(), &grpc.SessionValue{Value: sessValue})
	if err != nil {
		return errors.GetCustomError(err)
	}

	return nil
}
