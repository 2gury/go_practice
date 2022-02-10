package response

import (
	"context"
	"encoding/json"
	contextHelper "go_practice/9_clean_arch_db/internal/helpers/context"
	"go_practice/9_clean_arch_db/internal/helpers/errors"
	"net/http"
)

type Body map[string]interface{}

type Response struct {
	Code  int           `json:"code,omitempty"`
	Error *errors.Error `json:"error,omitempty"`
	Body  *Body         `json:"body,omitempty"`
}

func WriteStatusCode(w http.ResponseWriter ,ctx context.Context, statusCode int) {
	w.WriteHeader(statusCode)
	contextHelper.WriteStatusCodeContext(ctx, statusCode)
}

func WriteErrorResponse(w http.ResponseWriter ,ctx context.Context, err *errors.Error) {
	WriteStatusCode(w, ctx, err.HttpCode)
	json.NewEncoder(w).Encode(Response{Error: err})
}
