package mwares

import (
	"context"
	"encoding/json"
	"fmt"
	"go_practice/9_clean_arch_db/internal/consts"
	contextHelper "go_practice/9_clean_arch_db/internal/helpers/context"
	"go_practice/9_clean_arch_db/internal/helpers/errors"
	"go_practice/9_clean_arch_db/internal/session"
	cookieHelper "go_practice/9_clean_arch_db/tools/cookie"
	"go_practice/9_clean_arch_db/tools/logger"
	"go_practice/9_clean_arch_db/tools/response"
	"net/http"
	"time"
)

type MiddlewareManager struct {
	sessionUse session.SessionUsecase
}

func NewMiddlewareManager(sessUse session.SessionUsecase) *MiddlewareManager {
	return &MiddlewareManager{
		sessionUse: sessUse,
	}
}

func (mwm *MiddlewareManager) PanicCoverMiddleware(next http.Handler) http.Handler {
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

func (mwm *MiddlewareManager) AccessLogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info(fmt.Sprintf("%s %s %s", r.RemoteAddr, r.Method, r.URL.Path))
		ctx := r.Context()
		var code int
		ctx = context.WithValue(ctx,
			contextHelper.StatusCode, &code,
		)
		start := time.Now()
		next.ServeHTTP(w, r.WithContext(ctx))
		logger.Info(fmt.Sprintf("Status: %d Work time: %s", code, time.Since(start)))
	})
}

func (mwm *MiddlewareManager) CheckAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		sessCookie, err := cookieHelper.GetCookie(r, consts.SessionName)
		if err != nil {
			err := errors.Get(consts.CodeStatusUnauthorized)
			w.WriteHeader(err.HttpCode)
			contextHelper.WriteStatusCodeContext(ctx, err.HttpCode)
			json.NewEncoder(w).Encode(response.Response{Error: err})
			return
		}
		sess, customErr := mwm.sessionUse.Check(sessCookie.Value)
		if customErr != nil {
			w.WriteHeader(customErr.HttpCode)
			contextHelper.WriteStatusCodeContext(ctx, customErr.HttpCode)
			json.NewEncoder(w).Encode(response.Response{Error: customErr})
			return
		}
		ctx = context.WithValue(ctx,
			contextHelper.SessionValue, sess.Value,
		)
		ctx = context.WithValue(ctx,
			contextHelper.UserId, sess.UserId,
		)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
