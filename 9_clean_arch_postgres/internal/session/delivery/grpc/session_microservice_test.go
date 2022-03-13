package grpc

import (
	"context"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go_practice/9_clean_arch_db/internal/consts"
	"go_practice/9_clean_arch_db/internal/helpers/errors"
	mock_session "go_practice/9_clean_arch_db/internal/session/mocks"
	"google.golang.org/protobuf/types/known/durationpb"
	"testing"
)

func TestSessionService_Create(t *testing.T) {
	type mockBehaviour func(sessionRep *mock_session.MockSessionRepository)
	t.Parallel()

	testTable := []struct {
		name          string
		mockBehaviour mockBehaviour
		inSessionUserId  *SessionUserIdValue
		expError      error
	}{
		{
			name: "OK",
			mockBehaviour: func(sessionRep *mock_session.MockSessionRepository) {
				sessionRep.
					EXPECT().
					Create(gomock.Any()).
					Return(nil)
			},
			inSessionUserId: &SessionUserIdValue{
				Value: 1,
			},
			expError: nil,
		},
		{
			name: "Error: CodeInternalError",
			mockBehaviour: func(sessionRep *mock_session.MockSessionRepository) {
				sessionRep.
					EXPECT().
					Create(gomock.Any()).
					Return(fmt.Errorf("sql error"))
			},
			inSessionUserId: &SessionUserIdValue{
				Value: 1,
			},
			expError: errors.GetErrorFromGrpc(consts.CodeInternalError, fmt.Errorf("sql error")),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			sessionRep := mock_session.NewMockSessionRepository(ctrl)
			testCase.mockBehaviour(sessionRep)
			sessionSvc := NewSessionService(sessionRep)

			_, err := sessionSvc.Create(context.Background(), testCase.inSessionUserId)

			assert.Equal(t, err, testCase.expError)
		})
	}
}

func TestSessionService_Check(t *testing.T) {
	type mockBehaviour func(sessionRep *mock_session.MockSessionRepository, sessValue *SessionValue, sess *Session)
	t.Parallel()

	testTable := []struct {
		name          string
		mockBehaviour mockBehaviour
		inSessionValue *SessionValue
		outSession *Session
		expError      error
	}{
		{
			name: "OK",
			mockBehaviour: func(sessionRep *mock_session.MockSessionRepository, sessValue *SessionValue, sess *Session) {
				sessionRep.
					EXPECT().
					Get(sessValue.Value).
					Return(GrpcSessionToModel(sess), nil)
			},
			inSessionValue: &SessionValue{
				Value: "f823fg9gg3",
			},
			outSession: &Session{
				Value: "f823fg9gg3",
				UserId: 1,
				TimeDuration: &durationpb.Duration{
					Seconds: 1000,
				},
			},
			expError: nil,
		},
		{
			name: "Error: CodeInternalError",
			mockBehaviour: func(sessionRep *mock_session.MockSessionRepository, sessValue *SessionValue, sess *Session) {
				sessionRep.
					EXPECT().
					Get(sessValue.Value).
					Return(nil, fmt.Errorf("sql error"))
			},
			inSessionValue: &SessionValue{
				Value: "f823fg9gg3",
			},
			outSession: nil,
			expError: errors.GetErrorFromGrpc(consts.CodeInternalError, fmt.Errorf("sql error")),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			sessionRep := mock_session.NewMockSessionRepository(ctrl)
			testCase.mockBehaviour(sessionRep, testCase.inSessionValue, testCase.outSession)
			sessionSvc := NewSessionService(sessionRep)

			sess, err := sessionSvc.Check(context.Background(), testCase.inSessionValue)

			assert.Equal(t, sess, testCase.outSession)
			assert.Equal(t, err, testCase.expError)
		})
	}
}

func TestSessionService_Delete(t *testing.T) {
	type mockBehaviour func(sessionRep *mock_session.MockSessionRepository, sessValue *SessionValue)
	t.Parallel()

	testTable := []struct {
		name          string
		mockBehaviour mockBehaviour
		inSessionValue *SessionValue
		expError      error
	}{
		{
			name: "OK",
			mockBehaviour: func(sessionRep *mock_session.MockSessionRepository, sessValue *SessionValue) {
				sessionRep.
					EXPECT().
					Delete(sessValue.Value).
					Return( nil)
			},
			inSessionValue: &SessionValue{
				Value: "f823fg9gg3",
			},
			expError: nil,
		},
		{
			name: "Error: CodeInternalError",
			mockBehaviour: func(sessionRep *mock_session.MockSessionRepository, sessValue *SessionValue) {
				sessionRep.
					EXPECT().
					Delete(sessValue.Value).
					Return(fmt.Errorf("sql error"))
			},
			inSessionValue: &SessionValue{
				Value: "f823fg9gg3",
			},
			expError: errors.GetErrorFromGrpc(consts.CodeInternalError, fmt.Errorf("sql error")),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			sessionRep := mock_session.NewMockSessionRepository(ctrl)
			testCase.mockBehaviour(sessionRep, testCase.inSessionValue)
			sessionSvc := NewSessionService(sessionRep)

			_, err := sessionSvc.Delete(context.Background(), testCase.inSessionValue)

			assert.Equal(t, err, testCase.expError)
		})
	}
}
