package grpc

import (
	"context"
	"go_practice/9_clean_arch_db/internal/consts"
	"go_practice/9_clean_arch_db/internal/helpers/errors"
	"go_practice/9_clean_arch_db/internal/models"
	"go_practice/9_clean_arch_db/internal/session"
	"go_practice/9_clean_arch_db/tools/password"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type SessionService struct {
	sessionRep session.SessionRepository
	UnimplementedSessionServiceServer
}

func NewSessionService(rep session.SessionRepository) *SessionService {
	return &SessionService{
		sessionRep: rep,
	}
}

func (sm *SessionService) Create(ctx context.Context, userId *SessionUserIdValue) (*Session, error) {
	sess := models.NewSession(userId.Value)
	sess.Value = password.GetMD5Hash(sess.Value)

	err := sm.sessionRep.Create(sess)
	if err != nil {
		return nil, errors.GetErrorFromGrpc(consts.CodeInternalError, err)
	}

	return ModelSessionToGrpc(sess), nil
}

func (sm *SessionService) Check(ctx context.Context, sessValue *SessionValue) (*Session, error) {
	sess, err := sm.sessionRep.Get(sessValue.Value)
	if err != nil {
		return nil, errors.GetErrorFromGrpc(consts.CodeInternalError, err)
	}

	return ModelSessionToGrpc(sess), nil
}

func (sm *SessionService) Delete(ctx context.Context, sessValue *SessionValue) (*emptypb.Empty, error) {
	err := sm.sessionRep.Delete(sessValue.Value)
	if err != nil {
		return &emptypb.Empty{}, errors.GetErrorFromGrpc(consts.CodeInternalError, err)
	}

	return &emptypb.Empty{}, nil
}

