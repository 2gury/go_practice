package delivery

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go_practice/9_clean_arch_db/internal/consts"
	contextHelper "go_practice/9_clean_arch_db/internal/helpers/context"
	"go_practice/9_clean_arch_db/internal/mwares"
	"go_practice/9_clean_arch_db/internal/session"
	"go_practice/9_clean_arch_db/internal/user"
	cookieHelper "go_practice/9_clean_arch_db/tools/cookie"
	"go_practice/9_clean_arch_db/tools/request_reader"
	"go_practice/9_clean_arch_db/tools/response"
	"net/http"
)

type SessionHandler struct {
	sessionUse session.SessionUsecase
	userUse    user.UserUsecase
}

func NewSessionHandler(sessUse session.SessionUsecase, userUse user.UserUsecase) *SessionHandler {
	return &SessionHandler{
		sessionUse: sessUse,
		userUse:    userUse,
	}
}

func (h *SessionHandler) Configure(m *mux.Router, mwManager *mwares.MiddlewareManager) {
	m.HandleFunc("/api/v1/session", h.Login()).Methods("PUT", "OPTIONS")

	customMux := m.PathPrefix("/api/v1").Subrouter()
	customMux.Use(mwManager.CheckAuth)
	customMux.Path("/session").HandlerFunc(h.Logout()).Methods("DELETE", "OPTIONS")
}

func (h *SessionHandler) Login() http.HandlerFunc {
	type Request struct {
		Email    string `json:"email" valid:"email,required"`
		Password string `json:"password" valid:"stringlength(6|32),required"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		ctx := r.Context()
		req := &Request{}

		json.NewDecoder(r.Body).Decode(req)
		if err := request_reader.ValidateStruct(req); err != nil {
			response.WriteErrorResponse(w, ctx, err)
			return
		}
		dbUsr, err := h.userUse.GetByEmail(req.Email)
		if err != nil {
			response.WriteErrorResponse(w, ctx, err)
			return
		}
		if err = h.userUse.ComparePasswordAndHash(dbUsr, req.Password); err != nil {
			response.WriteErrorResponse(w, ctx, err)
			return
		}
		sess, err := h.sessionUse.Create(dbUsr.Id)
		if err != nil {
			response.WriteErrorResponse(w, ctx, err)
			return
		}
		cookie := cookieHelper.CreateCookie(sess)
		cookieHelper.SetCookie(w, cookie)

		response.WriteStatusCode(w, ctx, http.StatusOK)
		json.NewEncoder(w).Encode(response.Response{Body: &response.Body{
			"status": "OK",
		},
		})
	}
}

func (h *SessionHandler) Logout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		sessValue, err := contextHelper.GetSessionValue(ctx)
		if err != nil {
			response.WriteErrorResponse(w, ctx, err)
			return
		}
		if err := h.sessionUse.Delete(sessValue); err != nil {
			response.WriteErrorResponse(w, ctx, err)
			return
		}
		cookieHelper.DeleteCookie(w, r, consts.SessionName)

		response.WriteStatusCode(w, ctx, http.StatusOK)
		json.NewEncoder(w).Encode(response.Response{Body: &response.Body{
			"status": "OK",
		},
		})
	}
}
