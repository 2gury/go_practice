package contextHelper

import (
	"context"
)

type ContextKey int

const StatusCode ContextKey = 1

func WriteStatusCodeContext(ctx context.Context, code int) {
	statusCode, ok := ctx.Value(StatusCode).(*int)
	if !ok {
		return
	}
	*statusCode = code
}
