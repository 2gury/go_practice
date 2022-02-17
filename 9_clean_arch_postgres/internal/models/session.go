package models

import (
	uuid "github.com/satori/go.uuid"
	"go_practice/9_clean_arch_db/internal/consts"
	"time"
)

type Session struct {
	Value        string        `json:"value"`
	UserId       uint64        `json:"user_id"`
	TimeDuration time.Duration `json:"time_duration"`
}

func NewSession(userId uint64) *Session {
	value := uuid.NewV4().String()
	expires := consts.ExpiresDuration
	return &Session{
		Value:        value,
		UserId:       userId,
		TimeDuration: expires,
	}
}

func (sess *Session) GetTime() int {
	sess.TimeDuration.Seconds()
	return int(sess.TimeDuration.Seconds())
}
