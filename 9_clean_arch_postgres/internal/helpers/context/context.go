package contextHelper

import (
	"context"
)

type key int
const StatusCode key = 1

func WriteStatusCodeContext(ctx context.Context, code int) {
	statusCode, ok := ctx.Value(StatusCode).(*int)
	if !ok {
		return
	}
	*statusCode = code
}
