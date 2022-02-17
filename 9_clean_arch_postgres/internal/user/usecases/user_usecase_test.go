package usecases

import (
	"database/sql"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go_practice/9_clean_arch_db/internal/consts"
	"go_practice/9_clean_arch_db/internal/helpers/errors"
	"go_practice/9_clean_arch_db/internal/models"
	mock_user "go_practice/9_clean_arch_db/internal/user/mocks"
	"testing"
)

func TestUserUsecase_GetById(t *testing.T) {
	type mockBehaviour func(userRep *mock_user.MockUserRepository, usrId uint64, usr *models.User)
	t.Parallel()

	testTable := []struct {
		name          string
		mockBehaviour mockBehaviour
		inUserId uint64
		outUser *models.User
		expError *errors.Error
	}{
		{
			name: "OK",
			mockBehaviour: func(userRep *mock_user.MockUserRepository, usrId uint64, usr *models.User) {
				userRep.
					EXPECT().
					SelectById(usrId).
					Return(usr, nil)
			},
			inUserId: 1,
			outUser: &models.User{
				Id: 1,
				Email: "testmail@kek.ru",
				Password: "213f4g34gs3",
			},
			expError: nil,
		},
		{
			name: "Error: CodeUserDoesNotExist",
			mockBehaviour: func(userRep *mock_user.MockUserRepository, usrId uint64, usr *models.User) {
				userRep.
					EXPECT().
					SelectById(usrId).
					Return(nil, sql.ErrNoRows)
			},
			inUserId: 1,
			expError: errors.Get(consts.CodeUserDoesNotExist),
		},
		{
			name: "Error: CodeInternalError",
			mockBehaviour: func(userRep *mock_user.MockUserRepository, usrId uint64, usr *models.User) {
				userRep.
					EXPECT().
					SelectById(usrId).
					Return(nil, fmt.Errorf("sql error"))
			},
			inUserId: 1,
			expError: errors.Get(consts.CodeInternalError),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			userRep := mock_user.NewMockUserRepository(ctrl)
			testCase.mockBehaviour(userRep, testCase.inUserId, testCase.outUser)
			userUse := NewUserUsecase(userRep)

			usr, err := userUse.GetById(testCase.inUserId)

			assert.Equal(t, usr, testCase.outUser)
			assert.Equal(t, err, testCase.expError)
		})
	}
}

func TestUserUsecase_GetByEmail(t *testing.T) {
	type mockBehaviour func(userRep *mock_user.MockUserRepository, usrEmail string, usr *models.User)
	t.Parallel()

	testTable := []struct {
		name          string
		mockBehaviour mockBehaviour
		inUserEmail string
		outUser *models.User
		expError *errors.Error
	}{
		{
			name: "OK",
			mockBehaviour: func(userRep *mock_user.MockUserRepository, usrEmail string, usr *models.User) {
				userRep.
					EXPECT().
					SelectByEmail(usrEmail).
					Return(usr, nil)
			},
			inUserEmail: "testmail@kek.ru",
			outUser: &models.User{
				Id: 1,
				Email: "testmail@kek.ru",
				Password: "213f4g34gs3",
			},
			expError: nil,
		},
		{
			name: "Error: CodeUserDoesNotExist",
			mockBehaviour: func(userRep *mock_user.MockUserRepository, usrEmail string, usr *models.User) {
				userRep.
					EXPECT().
					SelectByEmail(usrEmail).
					Return(nil, sql.ErrNoRows)
			},
			inUserEmail: "testmail@kek.ru",
			expError: errors.Get(consts.CodeUserDoesNotExist),
		},
		{
			name: "Error: CodeInternalError",
			mockBehaviour: func(userRep *mock_user.MockUserRepository, usrEmail string, usr *models.User) {
				userRep.
					EXPECT().
					SelectByEmail(usrEmail).
					Return(nil, fmt.Errorf("sql error"))
			},
			inUserEmail: "testmail@kek.ru",
			expError: errors.Get(consts.CodeInternalError),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			userRep := mock_user.NewMockUserRepository(ctrl)
			testCase.mockBehaviour(userRep, testCase.inUserEmail, testCase.outUser)
			userUse := NewUserUsecase(userRep)

			usr, err := userUse.GetByEmail(testCase.inUserEmail)

			assert.Equal(t, usr, testCase.outUser)
			assert.Equal(t, err, testCase.expError)
		})
	}
}

func TestUserUsecase_Create(t *testing.T) {
	type mockBehaviour func(userRep *mock_user.MockUserRepository, usr *models.User, lastId uint64)
	t.Parallel()

	testTable := []struct {
		name          string
		mockBehaviour mockBehaviour
		inUser *models.User
		outLastId uint64
		expError *errors.Error
	}{
		{
			name: "OK",
			mockBehaviour: func(userRep *mock_user.MockUserRepository, usr *models.User, lastId uint64) {
				userRep.
					EXPECT().
					Insert(usr).
					Return(lastId, nil)
			},
			inUser: &models.User{
				Email: "testmail@kek.ru",
				Password: "213f4g34gs3",
			},
			outLastId: 1,
			expError: nil,
		},
		{
			name: "Error: CodeInternalError",
			mockBehaviour: func(userRep *mock_user.MockUserRepository, usr *models.User, lastId uint64) {
				userRep.
					EXPECT().
					Insert(usr).
					Return(lastId, fmt.Errorf("sql error"))
			},
			inUser: &models.User{
				Email: "testmail@kek.ru",
				Password: "213f4g34gs3",
			},
			outLastId: 0,
			expError: errors.Get(consts.CodeInternalError),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			userRep := mock_user.NewMockUserRepository(ctrl)
			testCase.mockBehaviour(userRep, testCase.inUser, testCase.outLastId)
			userUse := NewUserUsecase(userRep)

			lastId, err := userUse.Create(testCase.inUser)

			assert.Equal(t, lastId, testCase.outLastId)
			assert.Equal(t, err, testCase.expError)
		})
	}
}

func TestUserUsecase_UpdateUserPassword(t *testing.T) {
	type mockBehaviour func(userRep *mock_user.MockUserRepository, usr *models.User)
	t.Parallel()

	testTable := []struct {
		name          string
		mockBehaviour mockBehaviour
		inUser *models.User
		expError *errors.Error
	}{
		{
			name: "OK",
			mockBehaviour: func(userRep *mock_user.MockUserRepository, usr *models.User) {
				userRep.
					EXPECT().
					UpdatePassword(usr).
					Return(nil)
			},
			inUser: &models.User{
				Email: "testmail@kek.ru",
				Password: "213f4g34gs3",
			},
			expError: nil,
		},
		{
			name: "Error: CodeInternalError",
			mockBehaviour: func(userRep *mock_user.MockUserRepository, usr *models.User) {
				userRep.
					EXPECT().
					UpdatePassword(usr).
					Return(fmt.Errorf("sql error"))
			},
			inUser: &models.User{
				Email: "testmail@kek.ru",
				Password: "213f4g34gs3",
			},
			expError: errors.Get(consts.CodeInternalError),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			userRep := mock_user.NewMockUserRepository(ctrl)
			testCase.mockBehaviour(userRep, testCase.inUser)
			userUse := NewUserUsecase(userRep)

			err := userUse.UpdateUserPassword(testCase.inUser)

			assert.Equal(t, err, testCase.expError)
		})
	}
}

func TestUserUsecase_DeleteUserById(t *testing.T) {
	type mockBehaviour func(userRep *mock_user.MockUserRepository, usrId uint64)
	t.Parallel()

	testTable := []struct {
		name          string
		mockBehaviour mockBehaviour
		inUserId uint64
		expError *errors.Error
	}{
		{
			name: "OK",
			mockBehaviour: func(userRep *mock_user.MockUserRepository, usrId uint64) {
				userRep.
					EXPECT().
					DeleteById(usrId).
					Return(nil)
			},
			inUserId: 1,
			expError: nil,
		},
		{
			name: "Error: CodeInternalError",
			mockBehaviour: func(userRep *mock_user.MockUserRepository, usrId uint64) {
				userRep.
					EXPECT().
					DeleteById(usrId).
					Return(fmt.Errorf("sql error"))
			},
			inUserId: 1,
			expError: errors.Get(consts.CodeInternalError),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			userRep := mock_user.NewMockUserRepository(ctrl)
			testCase.mockBehaviour(userRep, testCase.inUserId)
			userUse := NewUserUsecase(userRep)

			err := userUse.DeleteUserById(testCase.inUserId)

			assert.Equal(t, err, testCase.expError)
		})
	}
}

func TestUserUsecase_ComparePasswords(t *testing.T) {
	t.Parallel()

	testTable := []struct {
		name          string
		inUser *models.User
		inPassword string
		expError *errors.Error
	}{
		{
			name: "OK",
			inUser: &models.User{
				Password: "$2a$14$oXhLcYUgwrsOSiReOtDt4u.TSvY1kM5U4K4.hDp5rBpohbShk.gny",
			},
			inPassword: "password",
			expError: nil,
		},
		{
			name: "Error: CodeWrongPasswords",
			inUser: &models.User{
				Password: "kek",
			},
			inPassword: "password",
			expError: errors.Get(consts.CodeWrongPasswords),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			userRep := mock_user.NewMockUserRepository(ctrl)
			userUse := NewUserUsecase(userRep)

			err := userUse.ComparePasswordAndHash(testCase.inUser, testCase.inPassword)

			assert.Equal(t, err, testCase.expError)
		})
	}
}