package grpc

import (
	"go_practice/9_clean_arch_db/internal/models"
	"google.golang.org/protobuf/types/known/durationpb"
)

func GrpcSessionToModel(sess *Session) *models.Session {
	return &models.Session{
		Value: sess.Value,
		UserId: sess.UserId,
		TimeDuration: sess.TimeDuration.AsDuration(),
	}
}

func ModelSessionToGrpc(sess *models.Session) *Session {
	return &Session{
		Value: sess.Value,
		UserId: sess.UserId,
		TimeDuration: durationpb.New(sess.TimeDuration),
	}
}
