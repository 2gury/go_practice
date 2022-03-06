package grpc

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go_practice/9_clean_arch_db/internal/consts"
	"go_practice/9_clean_arch_db/internal/helpers/errors"
	mock_user "go_practice/9_clean_arch_db/internal/user/mocks"
	"testing"
)

func TestUserService_GetById(t *testing.T) {
	type mockBehaviour func(userRep *mock_user.MockUserRepository, userId *UserIdValue, user *User)
	t.Parallel()

	testTable := []struct {
		name          string
		mockBehaviour mockBehaviour
		inUserId      *UserIdValue
		outUser       *User
		expError      error
	}{
		{
			name: "OK",
			mockBehaviour: func(userRep *mock_user.MockUserRepository, userId *UserIdValue, user *User) {
				userRep.
					EXPECT().
					SelectById(userId.Value).
					Return(GrpcUserToModel(user), nil)
			},
			inUserId: &UserIdValue{
				Value: 1,
			},
			outUser: &User{
				Id:    1,
				Email: "testmail@ya.ru",
			},
			expError: nil,
		},
		{
			name: "Error: CodeInternalError",
			mockBehaviour: func(userRep *mock_user.MockUserRepository, userId *UserIdValue, user *User) {
				userRep.
					EXPECT().
					SelectById(userId.Value).
					Return(nil, fmt.Errorf("sql error"))
			},
			inUserId: &UserIdValue{
				Value: 1,
			},
			outUser:  nil,
			expError: errors.GetErrorFromGrpc(consts.CodeInternalError, fmt.Errorf("sql error")),
		},
		{
			name: "Error: CodeUserDoesNotExist",
			mockBehaviour: func(userRep *mock_user.MockUserRepository, userId *UserIdValue, user *User) {
				userRep.
					EXPECT().
					SelectById(userId.Value).
					Return(nil, sql.ErrNoRows)
			},
			inUserId: &UserIdValue{
				Value: 1,
			},
			outUser:  nil,
			expError: errors.GetErrorFromGrpc(consts.CodeUserDoesNotExist, fmt.Errorf("sql: no rows in result set")),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			userRep := mock_user.NewMockUserRepository(ctrl)
			testCase.mockBehaviour(userRep, testCase.inUserId, testCase.outUser)
			userSvc := NewUserService(userRep)

			user, err := userSvc.GetById(context.Background(), testCase.inUserId)

			assert.Equal(t, user, testCase.outUser)
			assert.Equal(t, err, testCase.expError)
		})
	}
}

func TestUserService_Create(t *testing.T) {
	type mockBehaviour func(userRep *mock_user.MockUserRepository, outUserId *UserIdValue)
	t.Parallel()

	testTable := []struct {
		name          string
		mockBehaviour mockBehaviour
		inUser        *User
		outUserId     *UserIdValue
		expError      error
	}{
		{
			name: "OK",
			mockBehaviour: func(userRep *mock_user.MockUserRepository, outUserId *UserIdValue) {
				userRep.
					EXPECT().
					Insert(gomock.Any()).
					Return(outUserId.Value, nil)
			},
			inUser: &User{
				Id:    1,
				Email: "testmail@ya.ru",
			},
			outUserId: &UserIdValue{
				Value: 1,
			},
			expError: nil,
		},
		{
			name: "Error: CodeInternalError",
			mockBehaviour: func(userRep *mock_user.MockUserRepository, outUserId *UserIdValue) {
				userRep.
					EXPECT().
					Insert(gomock.Any()).
					Return(uint64(0), fmt.Errorf("sql error"))
			},
			inUser: &User{
				Id:    1,
				Email: "testmail@ya.ru",
			},
			outUserId: nil,
			expError:  errors.GetErrorFromGrpc(consts.CodeInternalError, fmt.Errorf("sql error")),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			userRep := mock_user.NewMockUserRepository(ctrl)
			testCase.mockBehaviour(userRep, testCase.outUserId)
			userSvc := NewUserService(userRep)

			lastId, err := userSvc.Create(context.Background(), testCase.inUser)

			assert.Equal(t, lastId, testCase.outUserId)
			assert.Equal(t, err, testCase.expError)
		})
	}
}

func TestUserService_GetByEmail(t *testing.T) {
	type mockBehaviour func(userRep *mock_user.MockUserRepository, email *EmailValue, user *User)
	t.Parallel()

	testTable := []struct {
		name          string
		mockBehaviour mockBehaviour
		inUserEmail  *EmailValue
		outUser       *User
		expError      error
	}{
		{
			name: "OK",
			mockBehaviour: func(userRep *mock_user.MockUserRepository, email *EmailValue, user *User) {
				userRep.
					EXPECT().
					SelectByEmail(email.Value).
					Return(GrpcUserToModel(user), nil)
			},
			inUserEmail: &EmailValue{
				Value: "testmail@ya.ru",
			},
			outUser: &User{
				Id:    1,
				Email: "testmail@ya.ru",
			},
			expError: nil,
		},
		{
			name: "Error: CodeInternalError",
			mockBehaviour: func(userRep *mock_user.MockUserRepository, email *EmailValue, user *User)  {
				userRep.
					EXPECT().
					SelectByEmail(email.Value).
					Return(nil, fmt.Errorf("sql error"))
			},
			inUserEmail: &EmailValue{
				Value: "testmail@ya.ru",
			},
			outUser:  nil,
			expError: errors.GetErrorFromGrpc(consts.CodeInternalError, fmt.Errorf("sql error")),
		},
		{
			name: "Error: CodeUserDoesNotExist",
			mockBehaviour: func(userRep *mock_user.MockUserRepository, email *EmailValue, user *User)  {
				userRep.
					EXPECT().
					SelectByEmail(email.Value).
					Return(nil, sql.ErrNoRows)
			},
			inUserEmail: &EmailValue{
				Value: "testmail@ya.ru",
			},
			outUser:  nil,
			expError: errors.GetErrorFromGrpc(consts.CodeUserDoesNotExist, fmt.Errorf("sql: no rows in result set")),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			userRep := mock_user.NewMockUserRepository(ctrl)
			testCase.mockBehaviour(userRep, testCase.inUserEmail, testCase.outUser)
			userSvc := NewUserService(userRep)

			user, err := userSvc.GetByEmail(context.Background(), testCase.inUserEmail)

			assert.Equal(t, user, testCase.outUser)
			assert.Equal(t, err, testCase.expError)
		})
	}
}

func TestUserService_UpdateUserPassword(t *testing.T) {
	type mockBehaviour func(userRep *mock_user.MockUserRepository)
	t.Parallel()

	testTable := []struct {
		name          string
		mockBehaviour mockBehaviour
		inUser        *User
		expError      error
	}{
		{
			name: "OK",
			mockBehaviour: func(userRep *mock_user.MockUserRepository) {
				userRep.
					EXPECT().
					UpdatePassword(gomock.Any()).
					Return(nil)
			},
			inUser: &User{
				Id: 1,
				Email: "testmail@ya.ru",
				Password: "pass",
			},
			expError: nil,
		},
		{
			name: "Error: CodeInternalError",
			mockBehaviour: func(userRep *mock_user.MockUserRepository) {
				userRep.
					EXPECT().
					UpdatePassword(gomock.Any()).
					Return(fmt.Errorf("sql error"))
			},
			inUser: &User{
				Id:    1,
				Email: "testmail@ya.ru",
			},
			expError:  errors.GetErrorFromGrpc(consts.CodeInternalError, fmt.Errorf("sql error")),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			userRep := mock_user.NewMockUserRepository(ctrl)
			testCase.mockBehaviour(userRep)
			userSvc := NewUserService(userRep)

			_, err := userSvc.UpdateUserPassword(context.Background(), testCase.inUser)

			assert.Equal(t, err, testCase.expError)
		})
	}
}

func TestUserService_DeleteUserById(t *testing.T) {
	type mockBehaviour func(userRep *mock_user.MockUserRepository, userId *UserIdValue)
	t.Parallel()

	testTable := []struct {
		name          string
		mockBehaviour mockBehaviour
		inUserId      *UserIdValue
		expError      error
	}{
		{
			name: "OK",
			mockBehaviour: func(userRep *mock_user.MockUserRepository, userId *UserIdValue) {
				userRep.
					EXPECT().
					DeleteById(userId.Value).
					Return(nil)
			},
			inUserId: &UserIdValue{
				Value: 1,
			},
			expError: nil,
		},
		{
			name: "Error: CodeInternalError",
			mockBehaviour: func(userRep *mock_user.MockUserRepository, userId *UserIdValue) {
				userRep.
					EXPECT().
					DeleteById(userId.Value).
					Return(fmt.Errorf("sql error"))
			},
			inUserId: &UserIdValue{
				Value: 1,
			},
			expError: errors.GetErrorFromGrpc(consts.CodeInternalError, fmt.Errorf("sql error")),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			userRep := mock_user.NewMockUserRepository(ctrl)
			testCase.mockBehaviour(userRep, testCase.inUserId)
			userSvc := NewUserService(userRep)

			_, err := userSvc.DeleteUserById(context.Background(), testCase.inUserId)

			assert.Equal(t, err, testCase.expError)
		})
	}
}
