package delivery

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go_practice/9_clean_arch_db/internal/consts"
	contextHelper "go_practice/9_clean_arch_db/internal/helpers/context"
	"go_practice/9_clean_arch_db/internal/helpers/errors"
	"go_practice/9_clean_arch_db/internal/models"
	"go_practice/9_clean_arch_db/internal/mwares"
	"go_practice/9_clean_arch_db/internal/session"
	"go_practice/9_clean_arch_db/internal/user"
	cookieHelper "go_practice/9_clean_arch_db/tools/cookie"
	"go_practice/9_clean_arch_db/tools/request_reader"
	"go_practice/9_clean_arch_db/tools/response"
	"net/http"
	"strconv"
)

type UserHandler struct {
	userUse    user.UserUsecase
	sessionUse session.SessionUsecase
}

func NewUserHandler(usrUse user.UserUsecase, sessUse session.SessionUsecase) *UserHandler {
	return &UserHandler{
		userUse:    usrUse,
		sessionUse: sessUse,
	}
}

func (h *UserHandler) Configure(m *mux.Router, mwManager *mwares.MiddlewareManager) {
	m.HandleFunc("/api/v1/user/{id:[0-9]+}", h.GetUserById()).Methods("GET")
	m.HandleFunc("/api/v1/user/register", h.RegisterUser()).Methods("PUT")

	customMux := m.PathPrefix("/api/v1").Subrouter()
	customMux.Use(mwManager.CheckAuth)
	customMux.Path("/user/password").HandlerFunc(h.ChangePassword()).Methods("POST")
	customMux.Path("/user/profile").HandlerFunc(h.DeleteUserById()).Methods("DELETE")
}

func (h *UserHandler) GetUserById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		userId, _ := mux.Vars(r)["id"]
		intUserId, parseErr := strconv.ParseUint(userId, 10, 64)

		if parseErr != nil {
			err := errors.Get(consts.CodeValidateError)
			response.WriteErrorResponse(w, ctx, err)
			return
		}
		usr, err := h.userUse.GetById(intUserId)
		if err != nil {
			response.WriteErrorResponse(w, ctx, err)
			return
		}

		response.WriteStatusCode(w, ctx, http.StatusOK)
		json.NewEncoder(w).Encode(response.Response{
			Body: &response.Body{
				"user": usr,
			},
		})
	}
}

func (u *UserHandler) RegisterUser() http.HandlerFunc {
	type Request struct {
		Email            string `json:"email" valid:"email,required"`
		Password         string `json:"password" valid:"stringlength(6|32),required"`
		RepeatedPassword string `json:"repeated_password" valid:"stringlength(6|32),required"`
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
		if err := u.userUse.ComparePasswords(req.Password, req.RepeatedPassword); err != nil {
			response.WriteErrorResponse(w, ctx, err)
			return
		}
		user := &models.User{
			Email:    req.Email,
			Password: req.Password,
		}
		lastId, err := u.userUse.Create(user)
		if err != nil {
			response.WriteErrorResponse(w, ctx, err)
			return
		}

		response.WriteStatusCode(w, ctx, http.StatusOK)
		json.NewEncoder(w).Encode(response.Response{
			Body: &response.Body{
				"id": lastId,
			},
		})
	}
}

func (u *UserHandler) ChangePassword() http.HandlerFunc {
	type Request struct {
		OldPassword string `json:"old_password" valid:"stringlength(6|32),required"`
		NewPassword string `json:"new_password" valid:"stringlength(6|32),required"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		ctx := r.Context()
		req := &Request{}

		userId, err := contextHelper.GetUserId(ctx)
		if err != nil {
			response.WriteErrorResponse(w, ctx, err)
			return
		}
		json.NewDecoder(r.Body).Decode(req)
		if err := request_reader.ValidateStruct(req); err != nil {
			response.WriteErrorResponse(w, ctx, err)
			return
		}
		usr, err := u.userUse.GetById(userId)
		if err != nil {
			response.WriteErrorResponse(w, ctx, err)
			return
		}
		if err = u.userUse.ComparePasswordAndHash(usr, req.OldPassword); err != nil {
			response.WriteErrorResponse(w, ctx, err)
			return
		}
		usr.Password = req.NewPassword
		if err = u.userUse.UpdateUserPassword(usr); err != nil {
			response.WriteErrorResponse(w, ctx, err)
			return
		}

		response.WriteStatusCode(w, ctx, http.StatusOK)
		json.NewEncoder(w).Encode(response.Response{
			Body: &response.Body{
				"status": "OK",
			},
		})
	}
}

func (u *UserHandler) DeleteUserById() http.HandlerFunc {
	type Request struct {
		Password string `json:"password" valid:"stringlength(6|32),required"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		ctx := r.Context()
		req := &Request{}

		userId, err := contextHelper.GetUserId(ctx)
		if err != nil {
			response.WriteErrorResponse(w, ctx, err)
			return
		}
		json.NewDecoder(r.Body).Decode(req)
		if err := request_reader.ValidateStruct(req); err != nil {
			response.WriteErrorResponse(w, ctx, err)
			return
		}
		usr, err := u.userUse.GetById(userId)
		if err != nil {
			response.WriteErrorResponse(w, ctx, err)
			return
		}
		if err = u.userUse.ComparePasswordAndHash(usr, req.Password); err != nil {
			response.WriteErrorResponse(w, ctx, err)
			return
		}
		if err = u.userUse.DeleteUserById(usr.Id); err != nil {
			response.WriteErrorResponse(w, ctx, err)
			return
		}
		cookieHelper.DeleteCookie(w, r, consts.SessionName)

		response.WriteStatusCode(w, ctx, http.StatusOK)
		json.NewEncoder(w).Encode(response.Response{
			Body: &response.Body{
				"status": "OK",
			},
		})
	}
}
