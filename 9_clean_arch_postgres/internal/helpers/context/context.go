package contextHelper

import (
	"context"
	"go_practice/9_clean_arch_db/internal/consts"
	"go_practice/9_clean_arch_db/internal/helpers/errors"
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

func GetSessionValue(ctx context.Context) (string, *errors.Error) {
	sessValue, ok := ctx.Value(SessionValue).(string)
	if !ok {
		return "", errors.Get(consts.CodeBadRequest)
	}
	return sessValue, nil
}

func GetUserId(ctx context.Context) (uint64, *errors.Error) {
	userId, ok := ctx.Value(UserId).(uint64)
	if !ok {
		return 0, errors.Get(consts.CodeBadRequest)
	}
	return userId, nil
}
