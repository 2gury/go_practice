package delivery

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"go_practice/9_clean_arch_db/internal/consts"
	contextHelper "go_practice/9_clean_arch_db/internal/helpers/context"
	"go_practice/9_clean_arch_db/internal/helpers/errors"
	"go_practice/9_clean_arch_db/internal/models"
	mock_user "go_practice/9_clean_arch_db/internal/user/mocks"
	"go_practice/9_clean_arch_db/tools/converter"
	"go_practice/9_clean_arch_db/tools/response"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestUserHandler_GetUserById(t *testing.T) {
	type mockBehaviour func(userUse *mock_user.MockUserUsecase, userId uint64)
	t.Parallel()

	user := &models.User{
		Id:    1,
		Email: "testmail@kek.ru",
	}

	testTable := []struct {
		name          string
		mockBehaviour mockBehaviour
		inPath        string
		inParams      map[string]string
		outUser       *models.User
		expStatusCode int
		expRespBody   response.Response
	}{
		{
			name: "OK",
			mockBehaviour: func(userUse *mock_user.MockUserUsecase, userId uint64) {
				userUse.
					EXPECT().
					GetById(userId).
					Return(user, nil)
			},
			inPath: "/api/v1/user/1",
			inParams: map[string]string{
				"id": "1",
			},
			expStatusCode: http.StatusOK,
			expRespBody: response.Response{
				Body: &response.Body{
					"user": user,
				},
			},
		},
		{
			name: "Error: CodeUserDoesNotExist",
			mockBehaviour: func(userUse *mock_user.MockUserUsecase, userId uint64) {
				userUse.
					EXPECT().
					GetById(userId).
					Return(nil, errors.Get(consts.CodeUserDoesNotExist))
			},
			inPath: "/api/v1/user/1000",
			inParams: map[string]string{
				"id": "1000",
			},
			expStatusCode: errors.Get(consts.CodeUserDoesNotExist).HttpCode,
			expRespBody: response.Response{
				Error: errors.Get(consts.CodeUserDoesNotExist),
			},
		},
		{
			name: "Error: CodeInternalError",
			mockBehaviour: func(userUse *mock_user.MockUserUsecase, userId uint64) {
				userUse.
					EXPECT().
					GetById(userId).
					Return(nil, errors.Get(consts.CodeInternalError))
			},
			inPath: "/api/v1/user/1",
			inParams: map[string]string{
				"id": "1",
			},
			expStatusCode: errors.Get(consts.CodeInternalError).HttpCode,
			expRespBody: response.Response{
				Error: errors.Get(consts.CodeInternalError),
			},
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mx := mux.NewRouter()
			r := httptest.NewRequest("GET", testCase.inPath, nil)
			r = mux.SetURLVars(r, testCase.inParams)
			w := httptest.NewRecorder()
			userUse := mock_user.NewMockUserUsecase(ctrl)

			userId, _ := mux.Vars(r)["id"]
			intUserId, _ := strconv.ParseUint(userId, 10, 64)
			testCase.mockBehaviour(userUse, intUserId)
			userHandler := NewUserHandler(userUse, nil)
			userHandler.Configure(mx, nil)

			userHandler.GetUserById()(w, r)
			expResBody, err := converter.AnyToBytesBuffer(testCase.expRespBody)
			if err != nil {
				t.Error(err.Error())
				return
			}
			bytes := converter.ReadBytes(w.Body)

			assert.Equal(t, testCase.expStatusCode, w.Code)
			assert.JSONEq(t, expResBody.String(), string(bytes))
		})
	}
}

func TestUserHandler_RegisterUser(t *testing.T) {
	type Request struct {
		Email            string `json:"email" valid:"email,required"`
		Password         string `json:"password" valid:"stringlength(6|32),required"`
		RepeatedPassword string `json:"repeated_password" valid:"stringlength(6|32),required"`
	}

	type mockBehaviourComparePasswords func(userUse *mock_user.MockUserUsecase, pass string, repPass string)
	type mockBehaviourCreate func(userUse *mock_user.MockUserUsecase, lastId uint64)

	t.Parallel()

	testTable := []struct {
		name                          string
		mockBehaviourComparePasswords mockBehaviourComparePasswords
		mockBehaviourCreate           mockBehaviourCreate
		inRequest                     *Request
		outLastId                     uint64
		expStatusCode                 int
		expRespBody                   response.Response
	}{
		{
			name: "OK",
			mockBehaviourComparePasswords: func(userUse *mock_user.MockUserUsecase, pass string, repPass string) {
				userUse.
					EXPECT().
					ComparePasswords(pass, repPass).
					Return(nil)
			},
			mockBehaviourCreate: func(userUse *mock_user.MockUserUsecase, lastId uint64) {
				userUse.
					EXPECT().
					Create(gomock.Any()).
					Return(lastId, nil)
			},
			inRequest: &Request{
				Email:            "testmail@kek.ru",
				Password:         "password",
				RepeatedPassword: "password",
			},
			outLastId:     1,
			expStatusCode: http.StatusOK,
			expRespBody: response.Response{
				Body: &response.Body{
					"id": 1,
				},
			},
		},
		{
			name: "Error: CodeUserPasswordsDoNotMatch",
			mockBehaviourComparePasswords: func(userUse *mock_user.MockUserUsecase, pass string, repPass string) {
				userUse.
					EXPECT().
					ComparePasswords(pass, repPass).
					Return(errors.Get(consts.CodeUserPasswordsDoNotMatch))
			},
			mockBehaviourCreate: func(userUse *mock_user.MockUserUsecase, lastId uint64) {},
			inRequest: &Request{
				Email:            "testmail@kek.ru",
				Password:         "password",
				RepeatedPassword: "pasword]",
			},
			outLastId:     1,
			expStatusCode: errors.Get(consts.CodeUserPasswordsDoNotMatch).HttpCode,
			expRespBody: response.Response{
				Error: errors.Get(consts.CodeUserPasswordsDoNotMatch),
			},
		},
		{
			name: "Error: CodeUserPasswordsDoNotMatch",
			mockBehaviourComparePasswords: func(userUse *mock_user.MockUserUsecase, pass string, repPass string) {
				userUse.
					EXPECT().
					ComparePasswords(pass, repPass).
					Return(nil)
			},
			mockBehaviourCreate: func(userUse *mock_user.MockUserUsecase, lastId uint64) {
				userUse.
					EXPECT().
					Create(gomock.Any()).
					Return(lastId, errors.Get(consts.CodeInternalError))
			},
			inRequest: &Request{
				Email:            "testmail@kek.ru",
				Password:         "password",
				RepeatedPassword: "password",
			},
			outLastId:     0,
			expStatusCode: errors.Get(consts.CodeInternalError).HttpCode,
			expRespBody: response.Response{
				Error: errors.Get(consts.CodeInternalError),
			},
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mx := mux.NewRouter()
			r := httptest.NewRequest("GET", "/api/v1/user/register", converter.AnyBytesToString(testCase.inRequest))
			w := httptest.NewRecorder()
			userUse := mock_user.NewMockUserUsecase(ctrl)

			testCase.mockBehaviourComparePasswords(userUse, testCase.inRequest.Password, testCase.inRequest.RepeatedPassword)
			testCase.mockBehaviourCreate(userUse, testCase.outLastId)
			userHandler := NewUserHandler(userUse, nil)
			userHandler.Configure(mx, nil)

			userHandler.RegisterUser()(w, r)
			expResBody, err := converter.AnyToBytesBuffer(testCase.expRespBody)
			if err != nil {
				t.Error(err.Error())
				return
			}
			bytes := converter.ReadBytes(w.Body)

			assert.Equal(t, testCase.expStatusCode, w.Code)
			assert.JSONEq(t, expResBody.String(), string(bytes))
		})
	}
}

func TestUserHandler_ChangePassword(t *testing.T) {
	type Request struct {
		OldPassword string `json:"old_password" valid:"stringlength(6|32),required"`
		NewPassword string `json:"new_password" valid:"stringlength(6|32),required"`
	}
	type mockBehaviourGetById func(userUse *mock_user.MockUserUsecase, userId uint64)
	type mockBehaviourComparePasswordAndHash func(userUse *mock_user.MockUserUsecase, oldPass string)
	type mockBehaviourUpdateUserPassword func(userUse *mock_user.MockUserUsecase)

	t.Parallel()

	testTable := []struct {
		name                                string
		mockBehaviourGetById                mockBehaviourGetById
		mockBehaviourComparePasswordAndHash mockBehaviourComparePasswordAndHash
		mockBehaviourUpdateUserPassword     mockBehaviourUpdateUserPassword
		inRequest                           *Request
		inUserId                            uint64
		expStatusCode                       int
		expRespBody                         response.Response
	}{
		{
			name: "OK",
			mockBehaviourGetById: func(userUse *mock_user.MockUserUsecase, userId uint64) {
				userUse.
					EXPECT().
					GetById(userId).
					Return(&models.User{}, nil)
			},
			mockBehaviourComparePasswordAndHash: func(userUse *mock_user.MockUserUsecase, oldPass string) {
				userUse.
					EXPECT().
					ComparePasswordAndHash(gomock.Any(), oldPass).
					Return(nil)
			},
			mockBehaviourUpdateUserPassword: func(userUse *mock_user.MockUserUsecase) {
				userUse.
					EXPECT().
					UpdateUserPassword(gomock.Any()).
					Return(nil)
			},
			inRequest: &Request{
				OldPassword: "password",
				NewPassword: "password",
			},
			inUserId:      1,
			expStatusCode: http.StatusOK,
			expRespBody: response.Response{
				Body: &response.Body{
					"status": "OK",
				},
			},
		},
		{
			name: "Error: CodeWrongPasswords",
			mockBehaviourGetById: func(userUse *mock_user.MockUserUsecase, userId uint64) {
				userUse.
					EXPECT().
					GetById(userId).
					Return(&models.User{}, nil)
			},
			mockBehaviourComparePasswordAndHash: func(userUse *mock_user.MockUserUsecase, oldPass string) {
				userUse.
					EXPECT().
					ComparePasswordAndHash(gomock.Any(), oldPass).
					Return(errors.Get(consts.CodeWrongPasswords))
			},
			mockBehaviourUpdateUserPassword: func(userUse *mock_user.MockUserUsecase) {},
			inRequest: &Request{
				OldPassword: "password",
				NewPassword: "password",
			},
			inUserId:      1,
			expStatusCode: errors.Get(consts.CodeWrongPasswords).HttpCode,
			expRespBody: response.Response{
				Error: errors.Get(consts.CodeWrongPasswords),
			},
		},
		{
			name: "Error: CodeUserDoesNotExist",
			mockBehaviourGetById: func(userUse *mock_user.MockUserUsecase, userId uint64) {
				userUse.
					EXPECT().
					GetById(userId).
					Return(nil, errors.Get(consts.CodeUserDoesNotExist))
			},
			mockBehaviourComparePasswordAndHash: func(userUse *mock_user.MockUserUsecase, oldPass string) {},
			mockBehaviourUpdateUserPassword:     func(userUse *mock_user.MockUserUsecase) {},
			inRequest: &Request{
				OldPassword: "password",
				NewPassword: "password",
			},
			inUserId:      1,
			expStatusCode: errors.Get(consts.CodeUserDoesNotExist).HttpCode,
			expRespBody: response.Response{
				Error: errors.Get(consts.CodeUserDoesNotExist),
			},
		},
		{
			name: "Error: CodeUserDoesNotExist",
			mockBehaviourGetById: func(userUse *mock_user.MockUserUsecase, userId uint64) {
				userUse.
					EXPECT().
					GetById(userId).
					Return(&models.User{}, nil)
			},
			mockBehaviourComparePasswordAndHash: func(userUse *mock_user.MockUserUsecase, oldPass string) {
				userUse.
					EXPECT().
					ComparePasswordAndHash(gomock.Any(), oldPass).
					Return(nil)
			},
			mockBehaviourUpdateUserPassword: func(userUse *mock_user.MockUserUsecase) {
				userUse.
					EXPECT().
					UpdateUserPassword(gomock.Any()).
					Return(errors.Get(consts.CodeInternalError))
			},
			inRequest: &Request{
				OldPassword: "password",
				NewPassword: "password",
			},
			inUserId:      1,
			expStatusCode: errors.Get(consts.CodeInternalError).HttpCode,
			expRespBody: response.Response{
				Error: errors.Get(consts.CodeInternalError),
			},
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mx := mux.NewRouter()
			r := httptest.NewRequest("POST", "/api/v1/user/password", converter.AnyBytesToString(testCase.inRequest))
			w := httptest.NewRecorder()
			ctx := r.Context()
			ctx = context.WithValue(ctx,
				contextHelper.UserId, testCase.inUserId,
			)
			userUse := mock_user.NewMockUserUsecase(ctrl)

			testCase.mockBehaviourGetById(userUse, testCase.inUserId)
			testCase.mockBehaviourComparePasswordAndHash(userUse, testCase.inRequest.OldPassword)
			testCase.mockBehaviourUpdateUserPassword(userUse)
			userHandler := NewUserHandler(userUse, nil)
			userHandler.Configure(mx, nil)

			userHandler.ChangePassword()(w, r.WithContext(ctx))
			expResBody, err := converter.AnyToBytesBuffer(testCase.expRespBody)
			if err != nil {
				t.Error(err.Error())
				return
			}
			bytes := converter.ReadBytes(w.Body)

			assert.Equal(t, testCase.expStatusCode, w.Code)
			assert.JSONEq(t, expResBody.String(), string(bytes))
		})
	}
}

func TestUserHandler_DeleteUserById(t *testing.T) {
	type Request struct {
		Password string `json:"password" valid:"stringlength(6|32),required"`
	}
	type mockBehaviourGetById func(userUse *mock_user.MockUserUsecase, userId uint64)
	type mockBehaviourComparePasswordAndHash func(userUse *mock_user.MockUserUsecase, oldPass string)
	type mockBehaviourDeleteUserById func(userUse *mock_user.MockUserUsecase, userId uint64)

	t.Parallel()

	testTable := []struct {
		name                                string
		mockBehaviourGetById                mockBehaviourGetById
		mockBehaviourComparePasswordAndHash mockBehaviourComparePasswordAndHash
		mockBehaviourDeleteUserById         mockBehaviourDeleteUserById
		inUserId                            uint64
		inRequest                           *Request
		expStatusCode                       int
		expRespBody                         response.Response
	}{
		{
			name: "OK",
			mockBehaviourGetById: func(userUse *mock_user.MockUserUsecase, userId uint64) {
				userUse.
					EXPECT().
					GetById(userId).
					Return(&models.User{}, nil)
			},
			mockBehaviourComparePasswordAndHash: func(userUse *mock_user.MockUserUsecase, pass string) {
				userUse.
					EXPECT().
					ComparePasswordAndHash(gomock.Any(), pass).
					Return(nil)
			},
			mockBehaviourDeleteUserById: func(userUse *mock_user.MockUserUsecase, userId uint64) {
				userUse.
					EXPECT().
					DeleteUserById(gomock.Any()).
					Return(nil)
			},
			inUserId: 1,
			inRequest: &Request{
				Password: "password",
			},
			expStatusCode: http.StatusOK,
			expRespBody: response.Response{
				Body: &response.Body{
					"status": "OK",
				},
			},
		},
		{
			name: "Error: CodeUserDoesNotExist",
			mockBehaviourGetById: func(userUse *mock_user.MockUserUsecase, userId uint64) {
				userUse.
					EXPECT().
					GetById(userId).
					Return(nil, errors.Get(consts.CodeUserDoesNotExist))
			},
			mockBehaviourComparePasswordAndHash: func(userUse *mock_user.MockUserUsecase, pass string) {},
			mockBehaviourDeleteUserById:         func(userUse *mock_user.MockUserUsecase, userId uint64) {},
			inUserId:                            1,
			inRequest: &Request{
				Password: "password",
			},
			expStatusCode: errors.Get(consts.CodeUserDoesNotExist).HttpCode,
			expRespBody: response.Response{
				Error: errors.Get(consts.CodeUserDoesNotExist),
			},
		},
		{
			name: "Error: CodeWrongPasswords",
			mockBehaviourGetById: func(userUse *mock_user.MockUserUsecase, userId uint64) {
				userUse.
					EXPECT().
					GetById(userId).
					Return(&models.User{}, nil)
			},
			mockBehaviourComparePasswordAndHash: func(userUse *mock_user.MockUserUsecase, pass string) {
				userUse.
					EXPECT().
					ComparePasswordAndHash(gomock.Any(), pass).
					Return(errors.Get(consts.CodeWrongPasswords))
			},
			mockBehaviourDeleteUserById: func(userUse *mock_user.MockUserUsecase, userId uint64) {},
			inUserId:                    1,
			inRequest: &Request{
				Password: "password",
			},
			expStatusCode: errors.Get(consts.CodeWrongPasswords).HttpCode,
			expRespBody: response.Response{
				Error: errors.Get(consts.CodeWrongPasswords),
			},
		},
		{
			name: "Error: CodeInternalError",
			mockBehaviourGetById: func(userUse *mock_user.MockUserUsecase, userId uint64) {
				userUse.
					EXPECT().
					GetById(userId).
					Return(&models.User{}, nil)
			},
			mockBehaviourComparePasswordAndHash: func(userUse *mock_user.MockUserUsecase, pass string) {
				userUse.
					EXPECT().
					ComparePasswordAndHash(gomock.Any(), pass).
					Return(nil)
			},
			mockBehaviourDeleteUserById: func(userUse *mock_user.MockUserUsecase, userId uint64) {
				userUse.
					EXPECT().
					DeleteUserById(gomock.Any()).
					Return(errors.Get(consts.CodeInternalError))
			},
			inUserId: 1,
			inRequest: &Request{
				Password: "password",
			},
			expStatusCode: errors.Get(consts.CodeInternalError).HttpCode,
			expRespBody: response.Response{
				Error: errors.Get(consts.CodeInternalError),
			},
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mx := mux.NewRouter()
			r := httptest.NewRequest("DELETE", "/api/v1/user/profile", converter.AnyBytesToString(testCase.inRequest))
			w := httptest.NewRecorder()
			ctx := r.Context()
			ctx = context.WithValue(ctx,
				contextHelper.UserId, testCase.inUserId,
			)
			userUse := mock_user.NewMockUserUsecase(ctrl)

			testCase.mockBehaviourGetById(userUse, testCase.inUserId)
			testCase.mockBehaviourComparePasswordAndHash(userUse, testCase.inRequest.Password)
			testCase.mockBehaviourDeleteUserById(userUse, testCase.inUserId)
			userHandler := NewUserHandler(userUse, nil)
			userHandler.Configure(mx, nil)

			userHandler.DeleteUserById()(w, r.WithContext(ctx))
			expResBody, err := converter.AnyToBytesBuffer(testCase.expRespBody)
			if err != nil {
				t.Error(err.Error())
				return
			}
			bytes := converter.ReadBytes(w.Body)

			assert.Equal(t, testCase.expStatusCode, w.Code)
			assert.JSONEq(t, expResBody.String(), string(bytes))
		})
	}
}
