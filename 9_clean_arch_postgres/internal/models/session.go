package models

import (
	uuid "github.com/satori/go.uuid"
	"go_practice/9_clean_arch_db/internal/consts"
	"time"
)

type Session struct {
	Id        uint64
	UserId    uint64
	Value     string
	ExpiresAt time.Time
}

func NewSession(userId uint64) *Session {
	value := uuid.NewV4().String()
	expires := consts.ExpiresDuration
	return &Session{
		UserId: userId,
		Value: value,
		ExpiresAt: time.Now().Add(expires),
	}
}
