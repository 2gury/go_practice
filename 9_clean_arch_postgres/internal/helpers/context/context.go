package contextHelper

import (
	"context"
)

type ContextKey int


const (
	StatusCode ContextKey = 101 + iota
	SessionValue
	UserId
)


func WriteStatusCodeContext(ctx context.Context, code int) {
	statusCode, ok := ctx.Value(StatusCode).(*int)
	if !ok {
		return
	}
	*statusCode = code
}

func GetSessionValue(ctx context.Context) string {
	sessValue, ok := ctx.Value(SessionValue).(string)
	if !ok {
		return ""
	}
	return sessValue
}

func GetUserId(ctx context.Context) uint64 {
	userId, ok := ctx.Value(UserId).(uint64)
	if !ok {
		return 0
	}
	return userId
}
