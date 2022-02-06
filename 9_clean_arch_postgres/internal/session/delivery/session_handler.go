package delivery

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go_practice/9_clean_arch_db/internal/consts"
	contextHelper "go_practice/9_clean_arch_db/internal/helpers/context"
	"go_practice/9_clean_arch_db/internal/helpers/errors"
	"go_practice/9_clean_arch_db/internal/models"
	"go_practice/9_clean_arch_db/internal/session"
	"go_practice/9_clean_arch_db/internal/user"
	cookieHelper "go_practice/9_clean_arch_db/tools/cookie"
	"go_practice/9_clean_arch_db/tools/request_reader"
	"go_practice/9_clean_arch_db/tools/response"
	"net/http"
)

type SessionHandler struct {
	sessionUse session.SessionUsecase
	userUse user.UserUsecase
}

func NewSessionHandler(sessUse session.SessionUsecase, userUse user.UserUsecase) *SessionHandler {
	return &SessionHandler{
		sessionUse: sessUse,
		userUse: userUse,
	}
}

func (h *SessionHandler) Configure(m *mux.Router) {
	m.HandleFunc("/api/v1/session", h.Login()).Methods("PUT")
	m.HandleFunc("/api/v1/session", h.Logout()).Methods("DELETE")
}

func (h *SessionHandler) Login() http.HandlerFunc {
	type Request struct {
		Email            string `json:"email" valid:"email,required"`
		Password         string `json:"password" valid:"stringlength(6|32),required"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		ctx := r.Context()
		req := &Request{}
		json.NewDecoder(r.Body).Decode(req)
		err := request_reader.ValidateStruct(req)
		if err != nil {
			w.WriteHeader(err.HttpCode)
			contextHelper.WriteStatusCodeContext(ctx, err.HttpCode)
			json.NewEncoder(w).Encode(response.Response{Error: err})
			return
		}
		usr := &models.User{
			Email: req.Email,
			Password: req.Password,
		}
		dbUsr, err := h.userUse.GetByEmail(usr.Email)
		if err != nil {
			w.WriteHeader(err.HttpCode)
			contextHelper.WriteStatusCodeContext(ctx, err.HttpCode)
			json.NewEncoder(w).Encode(response.Response{Error: err})
			return
		}
		err = h.userUse.ComparePasswordAndHash(dbUsr, usr.Password)
		if err != nil {
			w.WriteHeader(err.HttpCode)
			contextHelper.WriteStatusCodeContext(ctx, err.HttpCode)
			json.NewEncoder(w).Encode(response.Response{Error: err})
			return
		}
		sess, err := h.sessionUse.Create(dbUsr)
		if err != nil {
			w.WriteHeader(err.HttpCode)
			contextHelper.WriteStatusCodeContext(ctx, err.HttpCode)
			json.NewEncoder(w).Encode(response.Response{Error: err})
			return
		}
		cookie := cookieHelper.CreateCookie(sess)
		cookieHelper.SetCookie(w, cookie)
		w.WriteHeader(http.StatusOK)
		contextHelper.WriteStatusCodeContext(ctx, http.StatusOK)
		json.NewEncoder(w).Encode(response.Response{Body: &response.Body{
				"status": "OK",
			},
		})
	}
}

func (h *SessionHandler) Logout() http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		cookie, err := cookieHelper.GetCookie(r, consts.SessionName)
		if err != nil {
			err := errors.Get(consts.CodeStatusUnauthorized)
			w.WriteHeader(err.HttpCode)
			contextHelper.WriteStatusCodeContext(ctx, err.HttpCode)
			json.NewEncoder(w).Encode(response.Response{Error: err})
			return
		}
		cutomErr := h.sessionUse.Delete(cookie.Value)
		if cutomErr != nil {
			w.WriteHeader(cutomErr.HttpCode)
			contextHelper.WriteStatusCodeContext(ctx, cutomErr.HttpCode)
			json.NewEncoder(w).Encode(response.Response{Error: cutomErr})
			return
		}
		cookieHelper.DeleteCookie(w, r, consts.SessionName)
		w.WriteHeader(http.StatusOK)
		contextHelper.WriteStatusCodeContext(ctx, http.StatusOK)
		json.NewEncoder(w).Encode(response.Response{Body: &response.Body{
				"status": "OK",
			},
		})
	}
}
