package usecases

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go_practice/9_clean_arch_db/internal/consts"
	"go_practice/9_clean_arch_db/internal/helpers/errors"
	"go_practice/9_clean_arch_db/internal/models"
	"go_practice/9_clean_arch_db/internal/session/delivery/grpc"
	mock_grpc "go_practice/9_clean_arch_db/internal/session/delivery/grpc/mocks"
	"google.golang.org/protobuf/types/known/emptypb"
	"testing"
	"time"
)

func TestSessionUsecase_Create(t *testing.T) {
	type mockBehaviour func(sessionSvc *mock_grpc.MockSessionServiceClient, sess *models.Session)
	t.Parallel()

	testTable := []struct {
		name          string
		mockBehaviour mockBehaviour
		inUserId uint64
		outSession      *models.Session
		expError      *errors.Error
	}{
		{
			name: "OK",
			mockBehaviour: func(sessionSvc *mock_grpc.MockSessionServiceClient, sess *models.Session) {
				sessionSvc.
					EXPECT().
					Create(context.Background(), gomock.Any()).
					Return(grpc.ModelSessionToGrpc(sess), nil)
			},
			inUserId: 1,
			outSession: &models.Session{
				Value: "fsd7gs9segs",
				UserId: 1,
				TimeDuration: time.Hour,
			},
			expError: nil,
		},
		{
			name: "Error: CodeInternalError",
			mockBehaviour: func(sessionSvc *mock_grpc.MockSessionServiceClient, sess *models.Session) {
				sessionSvc.
					EXPECT().
					Create(context.Background(), gomock.Any()).
					Return(nil, errors.GetErrorFromGrpc(consts.CodeInternalError, errors.NilErrror))
			},
			inUserId: 1,
			outSession: nil,
			expError: errors.Get(consts.CodeInternalError),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			sessionSvc := mock_grpc.NewMockSessionServiceClient(ctrl)
			testCase.mockBehaviour(sessionSvc, testCase.outSession)
			sessionUse := NewSessionUsecase(sessionSvc)

			sess, err := sessionUse.Create(testCase.inUserId)

			assert.Equal(t, sess, testCase.outSession)
			assert.Equal(t, err, testCase.expError)
		})
	}
}

func TestSessionUsecase_Check(t *testing.T) {
	type mockBehaviour func(sessionSvc *mock_grpc.MockSessionServiceClient, sessValue string, sess *models.Session)
	t.Parallel()

	testTable := []struct {
		name          string
		mockBehaviour mockBehaviour
		inSessValue   string
		outSess       *models.Session
		expError      *errors.Error
	}{
		{
			name: "OK",
			mockBehaviour: func(sessionSvc *mock_grpc.MockSessionServiceClient, sessValue string, sess *models.Session) {
				sessionSvc.
					EXPECT().
					Check(context.Background(), gomock.Any()).
					Return(grpc.ModelSessionToGrpc(sess), nil)
			},
			inSessValue: "hyufsd9sdf9sfsn",
			outSess: &models.Session{
				Value:  "hyufsd9sdf9sfsn",
				UserId: 1,
			},
			expError: nil,
		},
		{
			name: "Error: CodeInternalError",
			mockBehaviour: func(sessionSvc *mock_grpc.MockSessionServiceClient, sessValue string, sess *models.Session) {
				sessionSvc.
					EXPECT().
					Check(context.Background(), gomock.Any()).
					Return(nil, errors.GetErrorFromGrpc(consts.CodeInternalError, errors.NilErrror))
			},
			inSessValue: "hyufsd9sdf9sfsn",
			outSess:     nil,
			expError:    errors.Get(consts.CodeInternalError),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			sessionSvc := mock_grpc.NewMockSessionServiceClient(ctrl)
			testCase.mockBehaviour(sessionSvc, testCase.inSessValue, testCase.outSess)
			sessionUse := NewSessionUsecase(sessionSvc)

			sess, err := sessionUse.Check(testCase.inSessValue)

			assert.Equal(t, sess, testCase.outSess)
			assert.Equal(t, err, testCase.expError)
		})
	}
}

func TestSessionUsecase_Delete(t *testing.T) {
	type mockBehaviour func(sessionSvc *mock_grpc.MockSessionServiceClient, sessValue string)
	t.Parallel()

	testTable := []struct {
		name          string
		mockBehaviour mockBehaviour
		inSessValue   string
		expError      *errors.Error
	}{
		{
			name: "OK",
			mockBehaviour: func(sessionSvc *mock_grpc.MockSessionServiceClient, sessValue string) {
				sessionSvc.
					EXPECT().
					Delete(context.Background(), gomock.Any()).
					Return(&emptypb.Empty{}, nil)
			},
			inSessValue: "hyufsd9sdf9sfsn",
			expError:    nil,
		},
		{
			name: "Error: CodeInternalError",
			mockBehaviour: func(sessionSvc *mock_grpc.MockSessionServiceClient, sessValue string) {
				sessionSvc.
					EXPECT().
					Delete(context.Background(), gomock.Any()).
					Return(&emptypb.Empty{}, errors.GetErrorFromGrpc(consts.CodeInternalError, errors.NilErrror))
			},
			inSessValue: "hyufsd9sdf9sfsn",
			expError:    errors.Get(consts.CodeInternalError),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			sessionSvc := mock_grpc.NewMockSessionServiceClient(ctrl)
			testCase.mockBehaviour(sessionSvc, testCase.inSessValue)
			sessionUse := NewSessionUsecase(sessionSvc)

			err := sessionUse.Delete(testCase.inSessValue)

			assert.Equal(t, err, testCase.expError)
		})
	}
}
