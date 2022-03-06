package usecases

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go_practice/9_clean_arch_db/internal/consts"
	"go_practice/9_clean_arch_db/internal/helpers/errors"
	"go_practice/9_clean_arch_db/internal/models"
	"go_practice/9_clean_arch_db/internal/user/delivery/grpc"
	mock_grpc "go_practice/9_clean_arch_db/internal/user/delivery/grpc/mocks"
	"google.golang.org/protobuf/types/known/emptypb"
	"testing"
)

func TestUserUsecase_GetById(t *testing.T) {
	type mockBehaviour func(userSvc *mock_grpc.MockUserServiceClient, usrId uint64, usr *models.User)
	t.Parallel()

	testTable := []struct {
		name          string
		mockBehaviour mockBehaviour
		inUserId      uint64
		outUser       *models.User
		expError      *errors.Error
	}{
		{
			name: "OK",
			mockBehaviour: func(userSvc *mock_grpc.MockUserServiceClient, usrId uint64, usr *models.User) {
				userSvc.
					EXPECT().
					GetById(context.Background(), &grpc.UserIdValue{Value: usrId}).
					Return(grpc.ModelUserToGrpc(usr), nil)
			},
			inUserId: 1,
			outUser: &models.User{
				Id:       1,
				Email:    "testmail@kek.ru",
				Password: "213f4g34gs3",
			},
			expError: nil,
		},
		{
			name: "Error: CodeInternalError",
			mockBehaviour: func(userSvc *mock_grpc.MockUserServiceClient, usrId uint64, usr *models.User) {
				userSvc.
					EXPECT().
					GetById(context.Background(), &grpc.UserIdValue{Value: usrId}).
					Return(nil, errors.GetErrorFromGrpc(consts.CodeInternalError, errors.NilErrror))
			},
			inUserId: 1,
			expError: errors.Get(consts.CodeInternalError),
		},
		{
			name: "Error: CodeUserDoesNotExist",
			mockBehaviour: func(userSvc *mock_grpc.MockUserServiceClient, usrId uint64, usr *models.User) {
				userSvc.
					EXPECT().
					GetById(context.Background(), &grpc.UserIdValue{Value: usrId}).
					Return(nil, errors.GetErrorFromGrpc(consts.CodeUserDoesNotExist, errors.NilErrror))
			},
			inUserId: 1,
			expError: errors.Get(consts.CodeUserDoesNotExist),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			userSvc := mock_grpc.NewMockUserServiceClient(ctrl)
			testCase.mockBehaviour(userSvc, testCase.inUserId, testCase.outUser)
			userUse := NewUserUsecase(userSvc)

			usr, err := userUse.GetById(testCase.inUserId)

			assert.Equal(t, usr, testCase.outUser)
			assert.Equal(t, err, testCase.expError)
		})
	}
}

func TestUserUsecase_GetByEmail(t *testing.T) {
	type mockBehaviour func(userSvc *mock_grpc.MockUserServiceClient, usrEmail string, usr *models.User)
	t.Parallel()

	testTable := []struct {
		name          string
		mockBehaviour mockBehaviour
		inUserEmail   string
		outUser       *models.User
		expError      *errors.Error
	}{
		{
			name: "OK",
			mockBehaviour: func(userSvc *mock_grpc.MockUserServiceClient, usrEmail string, usr *models.User) {
				userSvc.
					EXPECT().
					GetByEmail(context.Background(), &grpc.EmailValue{Value: usrEmail}).
					Return(grpc.ModelUserToGrpc(usr), nil)
			},
			inUserEmail: "testmail@kek.ru",
			outUser: &models.User{
				Id:       1,
				Email:    "testmail@kek.ru",
				Password: "213f4g34gs3",
			},
			expError: nil,
		},
		{
			name: "Error: CodeInternalError",
			mockBehaviour: func(userSvc *mock_grpc.MockUserServiceClient, usrEmail string, usr *models.User) {
				userSvc.
					EXPECT().
					GetByEmail(context.Background(), &grpc.EmailValue{Value: usrEmail}).
					Return(nil, errors.GetErrorFromGrpc(consts.CodeInternalError, errors.NilErrror))
			},
			inUserEmail: "testmail@kek.ru",
			expError:    errors.Get(consts.CodeInternalError),
		},
		{
			name: "Error: CodeUserDoesNotExist",
			mockBehaviour: func(userSvc *mock_grpc.MockUserServiceClient, usrEmail string, usr *models.User) {
				userSvc.
					EXPECT().
					GetByEmail(context.Background(), &grpc.EmailValue{Value: usrEmail}).
					Return(nil, errors.GetErrorFromGrpc(consts.CodeUserDoesNotExist, errors.NilErrror))
			},
			inUserEmail: "testmail@kek.ru",
			expError:    errors.Get(consts.CodeUserDoesNotExist),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			userSvc := mock_grpc.NewMockUserServiceClient(ctrl)
			testCase.mockBehaviour(userSvc, testCase.inUserEmail, testCase.outUser)
			userUse := NewUserUsecase(userSvc)

			usr, err := userUse.GetByEmail(testCase.inUserEmail)

			assert.Equal(t, usr, testCase.outUser)
			assert.Equal(t, err, testCase.expError)
		})
	}
}

func TestUserUsecase_Create(t *testing.T) {
	type mockBehaviour func(userSvc *mock_grpc.MockUserServiceClient, usr *models.User, lastId uint64)
	t.Parallel()

	testTable := []struct {
		name          string
		mockBehaviour mockBehaviour
		inUser        *models.User
		outLastId     uint64
		expError      *errors.Error
	}{
		{
			name: "OK",
			mockBehaviour: func(userSvc *mock_grpc.MockUserServiceClient, usr *models.User, lastId uint64) {
				userSvc.
					EXPECT().
					Create(context.Background(), grpc.ModelUserToGrpc(usr)).
					Return(&grpc.UserIdValue{Value: lastId}, nil)
			},
			inUser: &models.User{
				Email:    "testmail@kek.ru",
				Password: "213f4g34gs3",
			},
			outLastId: 1,
			expError:  nil,
		},
		{
			name: "Error: CodeInternalError",
			mockBehaviour: func(userSvc *mock_grpc.MockUserServiceClient, usr *models.User, lastId uint64) {
				userSvc.
					EXPECT().
					Create(context.Background(), grpc.ModelUserToGrpc(usr)).
					Return(nil, errors.GetErrorFromGrpc(consts.CodeInternalError, errors.NilErrror))
			},
			inUser: &models.User{
				Email:    "testmail@kek.ru",
				Password: "213f4g34gs3",
			},
			outLastId: 0,
			expError:  errors.Get(consts.CodeInternalError),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			userSvc := mock_grpc.NewMockUserServiceClient(ctrl)
			testCase.mockBehaviour(userSvc, testCase.inUser, testCase.outLastId)
			userUse := NewUserUsecase(userSvc)

			lastId, err := userUse.Create(testCase.inUser)

			assert.Equal(t, lastId, testCase.outLastId)
			assert.Equal(t, err, testCase.expError)
		})
	}
}

func TestUserUsecase_UpdateUserPassword(t *testing.T) {
	type mockBehaviour func(userSvc *mock_grpc.MockUserServiceClient, usr *models.User)
	t.Parallel()

	testTable := []struct {
		name          string
		mockBehaviour mockBehaviour
		inUser        *models.User
		expError      *errors.Error
	}{
		{
			name: "OK",
			mockBehaviour: func(userSvc *mock_grpc.MockUserServiceClient, usr *models.User) {
				userSvc.
					EXPECT().
					UpdateUserPassword(context.Background(), grpc.ModelUserToGrpc(usr)).
					Return(&emptypb.Empty{}, nil)
			},
			inUser: &models.User{
				Email:    "testmail@kek.ru",
				Password: "213f4g34gs3",
			},
			expError: nil,
		},
		{
			name: "Error: CodeInternalError",
			mockBehaviour: func(userSvc *mock_grpc.MockUserServiceClient, usr *models.User) {
				userSvc.
					EXPECT().
					UpdateUserPassword(context.Background(), grpc.ModelUserToGrpc(usr)).
					Return(&emptypb.Empty{}, errors.GetErrorFromGrpc(consts.CodeInternalError, errors.NilErrror))
			},
			inUser: &models.User{
				Email:    "testmail@kek.ru",
				Password: "213f4g34gs3",
			},
			expError: errors.Get(consts.CodeInternalError),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			userSvc := mock_grpc.NewMockUserServiceClient(ctrl)
			testCase.mockBehaviour(userSvc, testCase.inUser)
			userUse := NewUserUsecase(userSvc)

			err := userUse.UpdateUserPassword(testCase.inUser)

			assert.Equal(t, err, testCase.expError)
		})
	}
}

func TestUserUsecase_DeleteUserById(t *testing.T) {
	type mockBehaviour func(userSvc *mock_grpc.MockUserServiceClient, usrId uint64)
	t.Parallel()

	testTable := []struct {
		name          string
		mockBehaviour mockBehaviour
		inUserId      uint64
		expError      *errors.Error
	}{
		{
			name: "OK",
			mockBehaviour: func(userSvc *mock_grpc.MockUserServiceClient, usrId uint64) {
				userSvc.
					EXPECT().
					DeleteUserById(context.Background(), &grpc.UserIdValue{Value: usrId}).
					Return(&emptypb.Empty{}, nil)
			},
			inUserId: 1,
			expError: nil,
		},
		{
			name: "Error: CodeInternalError",
			mockBehaviour: func(userSvc *mock_grpc.MockUserServiceClient, usrId uint64) {
				userSvc.
					EXPECT().
					DeleteUserById(context.Background(), &grpc.UserIdValue{Value: usrId}).
					Return(&emptypb.Empty{}, errors.GetErrorFromGrpc(consts.CodeInternalError, errors.NilErrror))
			},
			inUserId: 1,
			expError: errors.Get(consts.CodeInternalError),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			userSvc := mock_grpc.NewMockUserServiceClient(ctrl)
			testCase.mockBehaviour(userSvc, testCase.inUserId)
			userUse := NewUserUsecase(userSvc)

			err := userUse.DeleteUserById(testCase.inUserId)

			assert.Equal(t, err, testCase.expError)
		})
	}
}

func TestUserUsecase_ComparePasswords(t *testing.T) {
	t.Parallel()

	testTable := []struct {
		name       string
		inUser     *models.User
		inPassword string
		expError   *errors.Error
	}{
		{
			name: "OK",
			inUser: &models.User{
				Password: "$2a$14$oXhLcYUgwrsOSiReOtDt4u.TSvY1kM5U4K4.hDp5rBpohbShk.gny",
			},
			inPassword: "password",
			expError:   nil,
		},
		{
			name: "Error: CodeWrongPasswords",
			inUser: &models.User{
				Password: "kek",
			},
			inPassword: "password",
			expError:   errors.Get(consts.CodeWrongPasswords),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			userSvc := mock_grpc.NewMockUserServiceClient(ctrl)
			userUse := NewUserUsecase(userSvc)

			err := userUse.ComparePasswordAndHash(testCase.inUser, testCase.inPassword)

			assert.Equal(t, err, testCase.expError)
		})
	}
}
