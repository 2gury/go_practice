package usecases

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go_practice/9_clean_arch_db/internal/consts"
	"go_practice/9_clean_arch_db/internal/helpers/errors"
	"go_practice/9_clean_arch_db/internal/models"
	mock_session "go_practice/9_clean_arch_db/internal/session/mocks"
	"testing"
)

func TestSessionUsecase_Create(t *testing.T) {
	type mockBehaviour func(sessionRep *mock_session.MockSessionRepository)
	t.Parallel()

	testTable := []struct {
		name string
		mockBehaviour mockBehaviour
		inUserId uint64
		expError *errors.Error
	}{
		{
			name: "OK",
			mockBehaviour: func(sessionRep *mock_session.MockSessionRepository) {
				sessionRep.
					EXPECT().
					Create(gomock.Any()).
					Return(nil)
			},
			inUserId: 1,
			expError: nil,
		},
		{
			name: "Error: CodeInternalError",
			mockBehaviour: func(sessionRep *mock_session.MockSessionRepository) {
				sessionRep.
					EXPECT().
					Create(gomock.Any()).
					Return(fmt.Errorf("redis error"))
			},
			inUserId: 1,
			expError: errors.Get(consts.CodeInternalError),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			sessionRep := mock_session.NewMockSessionRepository(ctrl)
			testCase.mockBehaviour(sessionRep)
			sessionUse := NewSessionUsecase(sessionRep)

			_, err := sessionUse.Create(testCase.inUserId)

			assert.Equal(t, err, testCase.expError)
		})
	}
}

func TestSessionUsecase_Check(t *testing.T) {
	type mockBehaviour func(sessionRep *mock_session.MockSessionRepository, sessValue string, sess *models.Session)
	t.Parallel()

	testTable := []struct {
		name string
		mockBehaviour mockBehaviour
		inSessValue string
		outSess *models.Session
		expError *errors.Error
	}{
		{
			name: "OK",
			mockBehaviour: func(sessionRep *mock_session.MockSessionRepository, sessValue string, sess *models.Session) {
				sessionRep.
					EXPECT().
					Get(sessValue).
					Return(sess, nil)
			},
			inSessValue: "hyufsd9sdf9sfsn",
			outSess: &models.Session{
				Value: "hyufsd9sdf9sfsn",
				UserId: 1,
			},
			expError: nil,
		},
		{
			name: "Error: CodeInternalError",
			mockBehaviour: func(sessionRep *mock_session.MockSessionRepository, sessValue string, sess *models.Session) {
				sessionRep.
					EXPECT().
					Get(sessValue).
					Return(nil, fmt.Errorf("redis error"))
			},
			inSessValue: "hyufsd9sdf9sfsn",
			outSess: nil,
			expError: errors.Get(consts.CodeInternalError),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			sessionRep := mock_session.NewMockSessionRepository(ctrl)
			testCase.mockBehaviour(sessionRep, testCase.inSessValue, testCase.outSess)
			sessionUse := NewSessionUsecase(sessionRep)

			sess, err := sessionUse.Check(testCase.inSessValue)

			assert.Equal(t, sess, testCase.outSess)
			assert.Equal(t, err, testCase.expError)
		})
	}
}

func TestSessionUsecase_Delete(t *testing.T) {
	type mockBehaviour func(sessionRep *mock_session.MockSessionRepository, sessValue string)
	t.Parallel()

	testTable := []struct {
		name          string
		mockBehaviour mockBehaviour
		inSessValue   string
		expError      *errors.Error
	}{
		{
			name: "OK",
			mockBehaviour: func(sessionRep *mock_session.MockSessionRepository, sessValue string) {
				sessionRep.
					EXPECT().
					Delete(sessValue).
					Return(nil)
			},
			inSessValue: "hyufsd9sdf9sfsn",
			expError: nil,
		},
		{
			name: "Error: CodeInternalError",
			mockBehaviour: func(sessionRep *mock_session.MockSessionRepository, sessValue string) {
				sessionRep.
					EXPECT().
					Delete(sessValue).
					Return(fmt.Errorf("redis error"))
			},
			inSessValue: "hyufsd9sdf9sfsn",
			expError: errors.Get(consts.CodeInternalError),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			sessionRep := mock_session.NewMockSessionRepository(ctrl)
			testCase.mockBehaviour(sessionRep, testCase.inSessValue)
			sessionUse := NewSessionUsecase(sessionRep)

			err := sessionUse.Delete(testCase.inSessValue)

			assert.Equal(t, err, testCase.expError)
		})
	}
}
