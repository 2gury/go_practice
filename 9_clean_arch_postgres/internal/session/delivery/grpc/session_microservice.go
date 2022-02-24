package grpc

import (
	"context"
	"go_practice/9_clean_arch_db/internal/consts"
	"go_practice/9_clean_arch_db/internal/session"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (sm *SessionService) Create(ctx context.Context, sess *Session) (*emptypb.Empty, error) {
	err := sm.sessionRep.Create(GrpcSessionToModel(sess))
	if err != nil {
		return &emptypb.Empty{}, status.Error(codes.Code(consts.CodeInternalError), err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (sm *SessionService) Check(ctx context.Context, sessValue *SessionValue) (*Session, error) {
	sess, err := sm.sessionRep.Get(sessValue.Value)
	if err != nil {
		return nil, status.Error(codes.Code(consts.CodeInternalError), err.Error())
	}

	return ModelSessionToGrpc(sess), nil
}

func (sm *SessionService) Delete(ctx context.Context, sessValue *SessionValue) (*emptypb.Empty, error) {
	err := sm.sessionRep.Delete(sessValue.Value)
	if err != nil {
		return &emptypb.Empty{}, status.Error(codes.Code(consts.CodeInternalError), err.Error())
	}

	return &emptypb.Empty{}, nil
}

