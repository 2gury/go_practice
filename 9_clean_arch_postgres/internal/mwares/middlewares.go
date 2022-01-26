package mwares

import (
	"context"
	"encoding/json"
	"fmt"
	"go_practice/9_clean_arch_db/internal/consts"
	contextHelper "go_practice/9_clean_arch_db/internal/helpers/context"
	"go_practice/9_clean_arch_db/internal/helpers/errors"
	"go_practice/9_clean_arch_db/tools/logger"
	"go_practice/9_clean_arch_db/tools/response"
	"net/http"
	"time"
)

func PanicCoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				ctx := r.Context()
				customErr := errors.Get(consts.CodeInternalError)
				w.WriteHeader(customErr.HttpCode)
				contextHelper.WriteStatusCodeContext(ctx, customErr.HttpCode)
				json.NewEncoder(w).Encode(response.Response{Error: customErr})
				logger.Warn(fmt.Sprintf("%s %d %s %s %s", r.RemoteAddr, customErr.HttpCode, r.Method, r.URL.Path, err))
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func AccessLogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		var code int
		ctx = context.WithValue(ctx,
			contextHelper.StatusCode, &code,
		)

		start := time.Now()
		next.ServeHTTP(w, r.WithContext(ctx))
		logger.Info(fmt.Sprintf("%s %d %s %s %s", r.RemoteAddr, code, r.Method, r.URL.Path, time.Since(start)))
	})
}
