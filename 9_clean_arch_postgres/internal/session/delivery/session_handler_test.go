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
	mock_session "go_practice/9_clean_arch_db/internal/session/mocks"
	mock_user "go_practice/9_clean_arch_db/internal/user/mocks"
	"go_practice/9_clean_arch_db/tools/converter"
	"go_practice/9_clean_arch_db/tools/response"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSessionHandler_Login(t *testing.T) {
	type Request struct {
		Email            string `json:"email" valid:"email,required"`
		Password         string `json:"password" valid:"stringlength(6|32),required"`
	}

	type mockBehaviourUserUseGetByEmail func(userUse *mock_user.MockUserUsecase, email string, usr *models.User)
	type mockBehaviourUserUseComparePasswordAndHash func(userUse *mock_user.MockUserUsecase, pass string, usr *models.User)
	type mockBehaviourSessionUseCreate func(sessUse *mock_session.MockSessionUsecase, usrId uint64, sess *models.Session)
	t.Parallel()

	testTable := []struct {
		name string
		userUseGetByEmail mockBehaviourUserUseGetByEmail
		userUseComparePasswordAndHash mockBehaviourUserUseComparePasswordAndHash
		sessUseCreate mockBehaviourSessionUseCreate
		inRequest *Request
		outUser *models.User
		outSession *models.Session
		expStatusCode int
		expRespBody response.Response
	}{
		{
			name: "OK",
			userUseGetByEmail: func(userUse *mock_user.MockUserUsecase, email string, usr *models.User) {
				userUse.
					EXPECT().
					GetByEmail(email).
					Return(usr, nil)
			},
			userUseComparePasswordAndHash: func(userUse *mock_user.MockUserUsecase, pass string, usr *models.User) {
				userUse.
					EXPECT().
					ComparePasswordAndHash(gomock.Any(), pass).
					Return(nil)
			},
			sessUseCreate: func(sessUse *mock_session.MockSessionUsecase, usrId uint64, sess *models.Session) {
				sessUse.
					EXPECT().
					Create(usrId).
					Return(sess, nil)
			},
			inRequest: &Request{
				"testmail@kek.ru",
				"test_password",
			},
			outUser: &models.User{
				Id: 1,
				Email: "testmail@kek.ru",
				Password: "fsd8fds8sdfd9",
				Role: "user",
			},
			outSession: models.NewSession(1),
			expStatusCode: http.StatusOK,
			expRespBody: response.Response{Body: &response.Body{
					"status": "OK",
				},
			},
		},
		{
			name: "Error: CodeUserDoesNotExist",
			userUseGetByEmail: func(userUse *mock_user.MockUserUsecase, email string, usr *models.User) {
				userUse.
					EXPECT().
					GetByEmail(email).
					Return(usr, nil)
			},
			userUseComparePasswordAndHash: func(userUse *mock_user.MockUserUsecase, pass string, usr *models.User) {
				userUse.
					EXPECT().
					ComparePasswordAndHash(gomock.Any(), pass).
					Return(errors.Get(consts.CodeWrongPasswords))
			},
			sessUseCreate: func(sessUse *mock_session.MockSessionUsecase, usrId uint64, sess *models.Session) {},
			inRequest: &Request{
				"testmail@kek.ru",
				"test_password",
			},
			outUser: &models.User{
				Id: 1,
				Email: "testmail@kek.ru",
				Password: "fsd8fds8sdfd9",
				Role: "user",
			},
			outSession: models.NewSession(1),
			expStatusCode: errors.Get(consts.CodeWrongPasswords).HttpCode,
			expRespBody: response.Response{
				Error: errors.Get(consts.CodeWrongPasswords),
			},
		},
		{
			name: "Error: CodeUserDoesNotExist",
			userUseGetByEmail: func(userUse *mock_user.MockUserUsecase, email string, usr *models.User) {
				userUse.
					EXPECT().
					GetByEmail(email).
					Return(nil, errors.Get(consts.CodeUserDoesNotExist))
			},
			userUseComparePasswordAndHash: func(userUse *mock_user.MockUserUsecase, pass string, usr *models.User) {},
			sessUseCreate: func(sessUse *mock_session.MockSessionUsecase, usrId uint64, sess *models.Session) {},
			inRequest: &Request{
				"testmail@kek.ru",
				"test_password",
			},
			outUser: &models.User{
				Id: 1,
				Email: "testmail@kek.ru",
				Password: "fsd8fds8sdfd9",
				Role: "user",
			},
			outSession: models.NewSession(1),
			expStatusCode: errors.Get(consts.CodeUserDoesNotExist).HttpCode,
			expRespBody: response.Response{
				Error: errors.Get(consts.CodeUserDoesNotExist),
			},
		},
		{
			name: "Error: CodeInternalError",
			userUseGetByEmail: func(userUse *mock_user.MockUserUsecase, email string, usr *models.User) {
				userUse.
					EXPECT().
					GetByEmail(email).
					Return(usr, nil)
			},
			userUseComparePasswordAndHash: func(userUse *mock_user.MockUserUsecase, pass string, usr *models.User) {
				userUse.
					EXPECT().
					ComparePasswordAndHash(gomock.Any(), pass).
					Return(nil)
			},
			sessUseCreate: func(sessUse *mock_session.MockSessionUsecase, usrId uint64, sess *models.Session) {
				sessUse.
					EXPECT().
					Create(usrId).
					Return(nil, errors.Get(consts.CodeInternalError))
			},
			inRequest: &Request{
				"testmail@kek.ru",
				"test_password",
			},
			outUser: &models.User{
				Id: 1,
				Email: "testmail@kek.ru",
				Password: "fsd8fds8sdfd9",
				Role: "user",
			},
			outSession: models.NewSession(1),
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
			r := httptest.NewRequest("PUT", "/api/v1/session", converter.AnyBytesToString(testCase.inRequest))
			w := httptest.NewRecorder()
			userUse := mock_user.NewMockUserUsecase(ctrl)
			sessUse := mock_session.NewMockSessionUsecase(ctrl)

			testCase.userUseGetByEmail(userUse, testCase.inRequest.Email, testCase.outUser)
			testCase.userUseComparePasswordAndHash(userUse, testCase.inRequest.Password, testCase.outUser)
			testCase.sessUseCreate(sessUse, testCase.outUser.Id ,testCase.outSession)

			sessHandler := NewSessionHandler(sessUse, userUse)
			sessHandler.Configure(mx, nil)

			sessHandler.Login()(w, r)
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

func TestSessionHandler_Logout(t *testing.T) {
	type mockBehaviour func(sessUse *mock_session.MockSessionUsecase, sessValue string)
	t.Parallel()

	testTable := []struct {
		name string
		mockBehaviour mockBehaviour
		inSession *models.Session
		expStatusCode int
		expRespBody response.Response
	}{
		{
			name: "OK",
			mockBehaviour: func(sessUse *mock_session.MockSessionUsecase, sessValue string) {
				sessUse.
					EXPECT().
					Delete(sessValue).
					Return(nil)
			},
			inSession: models.NewSession(1),
			expStatusCode: http.StatusOK,
			expRespBody: response.Response{Body: &response.Body{
					"status": "OK",
				},
			},
		},
		{
			name: "Error: CodeInternalError",
			mockBehaviour: func(sessUse *mock_session.MockSessionUsecase, sessValue string) {
				sessUse.
					EXPECT().
					Delete(sessValue).
					Return(errors.Get(consts.CodeInternalError))
			},
			inSession: models.NewSession(1),
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
			r := httptest.NewRequest("PUT", "/api/v1/session", nil)
			w := httptest.NewRecorder()
			ctx := r.Context()
			ctx = context.WithValue(ctx,
				contextHelper.SessionValue, testCase.inSession.Value,
			)

			userUse := mock_user.NewMockUserUsecase(ctrl)
			sessUse := mock_session.NewMockSessionUsecase(ctrl)

			testCase.mockBehaviour(sessUse, testCase.inSession.Value)

			sessHandler := NewSessionHandler(sessUse, userUse)
			sessHandler.Configure(mx, nil)

			sessHandler.Logout()(w, r.WithContext(ctx))
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
