package main

import (
	"context"
	"go_practice/17_protobuf/session"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"math/rand"
	"sync"
)

type SessionManager struct {
	mu *sync.Mutex
	sessions map[string]*session.Session
	session.UnimplementedAuthServer
}

func NewSessionManager() *SessionManager {
	return &SessionManager{
		mu: &sync.Mutex{},
		sessions: map[string]*session.Session{},
	}
}

func (sm *SessionManager) Create(ctx context.Context, sess *session.Session) (*session.SessionId, error) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	grpc.SendHeader(ctx, metadata.Pairs(
		"header-from-server", "kek",
		),
	)
	grpc.SetTrailer(ctx, metadata.Pairs(
		"trailer-from-server", "kuk",
		),
	)
	sessId := &session.SessionId{Id: RandStringRunes(10)}
	sm.sessions[sessId.Id] = sess
	return sessId, nil
}

func (sm *SessionManager) Check(ctx context.Context, id *session.SessionId) (*session.Session, error) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sess, ok := sm.sessions[id.Id]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "session not found")
	}
	return sess, nil
}

func (sm *SessionManager) Delete(ctx context.Context, id *session.SessionId) (*session.Nothing, error) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	delete(sm.sessions, id.Id)
	return &session.Nothing{Dummy: true}, nil
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}