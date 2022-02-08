package delivery

import (
	"encoding/json"
	"github.com/asaskevich/govalidator"
	"github.com/gorilla/mux"
	"go_practice/9_clean_arch_db/internal/consts"
	contextHelper "go_practice/9_clean_arch_db/internal/helpers/context"
	"go_practice/9_clean_arch_db/internal/helpers/errors"
	"go_practice/9_clean_arch_db/internal/models"
	"go_practice/9_clean_arch_db/internal/session"
	"go_practice/9_clean_arch_db/internal/user"
	cookieHelper "go_practice/9_clean_arch_db/tools/cookie"
	"go_practice/9_clean_arch_db/tools/logger"
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

func (h *UserHandler) Configure(m *mux.Router) {
	m.HandleFunc("/api/v1/user/{id:[0-9]+}", h.GetUserById()).Methods("GET")
	m.HandleFunc("/api/v1/user/register", h.RegisterUser()).Methods("PUT")
	m.HandleFunc("/api/v1/user/password", h.ChangePassword()).Methods("POST")
	m.HandleFunc("/api/v1/user/profile", h.DeleteUserById()).Methods("DELETE")
}

func (h *UserHandler) GetUserById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		userId, _ := mux.Vars(r)["id"]
		intUserId, parseErr := strconv.ParseUint(userId, 10, 64)
		if parseErr != nil {
			err := errors.Get(consts.CodeValidateError)
			w.WriteHeader(err.HttpCode)
			contextHelper.WriteStatusCodeContext(ctx, err.HttpCode)
			json.NewEncoder(w).Encode(response.Response{Error: err})
			return
		}
		usr, err := h.userUse.GetById(intUserId)
		if err != nil {
			w.WriteHeader(err.HttpCode)
			contextHelper.WriteStatusCodeContext(ctx, err.HttpCode)
			json.NewEncoder(w).Encode(response.Response{Error: err})
			return
		}
		w.WriteHeader(http.StatusOK)
		contextHelper.WriteStatusCodeContext(ctx, http.StatusOK)
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
		err := request_reader.ValidateStruct(req)
		if err != nil {
			w.WriteHeader(err.HttpCode)
			contextHelper.WriteStatusCodeContext(ctx, err.HttpCode)
			json.NewEncoder(w).Encode(response.Response{Error: err})
			return
		}
		err = u.userUse.ComparePasswords(req.Password, req.RepeatedPassword)
		if err != nil {
			w.WriteHeader(err.HttpCode)
			contextHelper.WriteStatusCodeContext(ctx, err.HttpCode)
			json.NewEncoder(w).Encode(response.Response{Error: err})
			return
		}
		user := &models.User{
			Email:    req.Email,
			Password: req.Password,
		}
		lastId, err := u.userUse.Create(user)
		if err != nil {
			w.WriteHeader(err.HttpCode)
			contextHelper.WriteStatusCodeContext(ctx, err.HttpCode)
			json.NewEncoder(w).Encode(response.Response{Error: err})
			return
		}
		w.WriteHeader(http.StatusOK)
		contextHelper.WriteStatusCodeContext(ctx, http.StatusOK)
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
		json.NewDecoder(r.Body).Decode(req)
		customErr := request_reader.ValidateStruct(req)
		if customErr != nil {
			w.WriteHeader(customErr.HttpCode)
			contextHelper.WriteStatusCodeContext(ctx, customErr.HttpCode)
			json.NewEncoder(w).Encode(response.Response{Error: customErr})
			return
		}
		sessCookie, err := cookieHelper.GetCookie(r, consts.SessionName)
		if err != nil {
			err := errors.Get(consts.CodeStatusUnauthorized)
			w.WriteHeader(err.HttpCode)
			contextHelper.WriteStatusCodeContext(ctx, err.HttpCode)
			json.NewEncoder(w).Encode(response.Response{Error: err})
			return
		}
		if ok := govalidator.IsMD5(sessCookie.Value); !ok {
			err := errors.Get(consts.CodeStatusUnauthorized)
			w.WriteHeader(err.HttpCode)
			contextHelper.WriteStatusCodeContext(ctx, err.HttpCode)
			json.NewEncoder(w).Encode(response.Response{Error: err})
			return
		}
		sess, customErr := u.sessionUse.Check(sessCookie.Value)
		if customErr != nil {
			w.WriteHeader(customErr.HttpCode)
			contextHelper.WriteStatusCodeContext(ctx, customErr.HttpCode)
			json.NewEncoder(w).Encode(response.Response{Error: customErr})
			return
		}
		usr, customErr := u.userUse.GetById(sess.UserId)
		if customErr != nil {
			w.WriteHeader(customErr.HttpCode)
			contextHelper.WriteStatusCodeContext(ctx, customErr.HttpCode)
			json.NewEncoder(w).Encode(response.Response{Error: customErr})
			return
		}
		customErr = u.userUse.ComparePasswordAndHash(usr, req.OldPassword)
		logger.Info(req.OldPassword)
		if customErr != nil {
			w.WriteHeader(customErr.HttpCode)
			contextHelper.WriteStatusCodeContext(ctx, customErr.HttpCode)
			json.NewEncoder(w).Encode(response.Response{Error: customErr})
			return
		}
		usr.Password = req.NewPassword
		customErr = u.userUse.UpdateUserPassword(usr)
		logger.Info(usr.Password)
		if customErr != nil {
			w.WriteHeader(customErr.HttpCode)
			contextHelper.WriteStatusCodeContext(ctx, customErr.HttpCode)
			json.NewEncoder(w).Encode(response.Response{Error: customErr})
			return
		}
		w.WriteHeader(http.StatusOK)
		contextHelper.WriteStatusCodeContext(ctx, http.StatusOK)
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
		Confirmation bool `json:"confirmation"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		ctx := r.Context()
		req := &Request{}
		json.NewDecoder(r.Body).Decode(req)
		customErr := request_reader.ValidateStruct(req)
		if customErr != nil {
			w.WriteHeader(customErr.HttpCode)
			contextHelper.WriteStatusCodeContext(ctx, customErr.HttpCode)
			json.NewEncoder(w).Encode(response.Response{Error: customErr})
			return
		}
		if !req.Confirmation {
			err := errors.Get(consts.CodeUserNotConfirmation)
			w.WriteHeader(err.HttpCode)
			contextHelper.WriteStatusCodeContext(ctx, err.HttpCode)
			json.NewEncoder(w).Encode(response.Response{Error: err})
			return
		}
		sessCookie, err := cookieHelper.GetCookie(r, consts.SessionName)
		if err != nil {
			err := errors.Get(consts.CodeStatusUnauthorized)
			w.WriteHeader(err.HttpCode)
			contextHelper.WriteStatusCodeContext(ctx, err.HttpCode)
			json.NewEncoder(w).Encode(response.Response{Error: err})
			return
		}
		if ok := govalidator.IsMD5(sessCookie.Value); !ok {
			err := errors.Get(consts.CodeStatusUnauthorized)
			w.WriteHeader(err.HttpCode)
			contextHelper.WriteStatusCodeContext(ctx, err.HttpCode)
			json.NewEncoder(w).Encode(response.Response{Error: err})
			return
		}
		sess, customErr := u.sessionUse.Check(sessCookie.Value)
		if customErr != nil {
			w.WriteHeader(customErr.HttpCode)
			contextHelper.WriteStatusCodeContext(ctx, customErr.HttpCode)
			json.NewEncoder(w).Encode(response.Response{Error: customErr})
			return
		}
		usr, customErr := u.userUse.GetById(sess.UserId)
		if customErr != nil {
			w.WriteHeader(customErr.HttpCode)
			contextHelper.WriteStatusCodeContext(ctx, customErr.HttpCode)
			json.NewEncoder(w).Encode(response.Response{Error: customErr})
			return
		}
		customErr = u.userUse.ComparePasswordAndHash(usr, req.Password)
		if customErr != nil {
			w.WriteHeader(customErr.HttpCode)
			contextHelper.WriteStatusCodeContext(ctx, customErr.HttpCode)
			json.NewEncoder(w).Encode(response.Response{Error: customErr})
			return
		}
		customErr = u.userUse.DeleteUserById(usr.Id)
		if customErr != nil {
			w.WriteHeader(customErr.HttpCode)
			contextHelper.WriteStatusCodeContext(ctx, customErr.HttpCode)
			json.NewEncoder(w).Encode(response.Response{Error: customErr})
			return
		}
		w.WriteHeader(http.StatusOK)
		contextHelper.WriteStatusCodeContext(ctx, http.StatusOK)
		json.NewEncoder(w).Encode(response.Response{
			Body: &response.Body{
				"status": "OK",
			},
		})
	}
}
