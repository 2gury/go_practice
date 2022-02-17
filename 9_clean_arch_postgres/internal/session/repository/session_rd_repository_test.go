package repository

import (
	"encoding/json"
	"fmt"
	"github.com/rafaeljusto/redigomock"
	"github.com/stretchr/testify/assert"
	"go_practice/9_clean_arch_db/internal/models"
	"testing"
)

func TestSessionRdRepository_Get(t *testing.T) {
	type mockBehaviour func(conn *redigomock.Conn, sessionValue string, session []byte)
	t.Parallel()

	rdConn := redigomock.NewConn()
	defer rdConn.Close()

	testTable := []struct {
		name          string
		mockBehaviour mockBehaviour
		inSession     *models.Session
		expError      error
	}{
		{
			name: "OK",
			mockBehaviour: func(conn *redigomock.Conn, sessionValue string, session []byte) {
				conn.Command("GET", sessionValue).Expect(session)
			},
			inSession: models.NewSession(1),
			expError:  nil,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			bytesSess, _ := json.Marshal(testCase.inSession)
			testCase.mockBehaviour(rdConn, testCase.inSession.Value, bytesSess)

			sessionRep := NewSessionRdRepository(rdConn)
			sessionFromRd, err := sessionRep.Get(testCase.inSession.Value)

			assert.Equal(t, sessionFromRd, testCase.inSession)
			assert.Equal(t, err, testCase.expError)
		})
	}
}

func TestSessionRdRepository_Create(t *testing.T) {
	type mockBehaviour func(conn *redigomock.Conn, sessionValue string, session []byte, timeExpire int)
	t.Parallel()

	rdConn := redigomock.NewConn()
	defer rdConn.Close()

	testTable := []struct {
		name          string
		mockBehaviour mockBehaviour
		inSession     *models.Session
		expError      error
	}{
		{
			name: "OK",
			mockBehaviour: func(conn *redigomock.Conn, sessionValue string, session []byte, timeExpire int) {
				conn.Command("SET", sessionValue, session, "EX", timeExpire).Expect("OK")
			},
			inSession: models.NewSession(1),
			expError:  nil,
		},
		{
			name: "Error: redis not OK",
			mockBehaviour: func(conn *redigomock.Conn, sessionValue string, session []byte, timeExpire int) {
				conn.Command("SET", sessionValue, session, "EX", timeExpire).Expect("KEK")
			},
			inSession: models.NewSession(1),
			expError:  fmt.Errorf("redis: not OK"),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			bytesSess, _ := json.Marshal(testCase.inSession)
			testCase.mockBehaviour(rdConn, testCase.inSession.Value, bytesSess, testCase.inSession.GetTime())

			sessionRep := NewSessionRdRepository(rdConn)
			err := sessionRep.Create(testCase.inSession)

			assert.Equal(t, err, testCase.expError)
		})
	}
}

func TestSessionRdRepository_Delete(t *testing.T) {
	type mockBehaviour func(conn *redigomock.Conn, sessionValue string)
	t.Parallel()

	rdConn := redigomock.NewConn()
	defer rdConn.Close()

	testTable := []struct {
		name           string
		mockBehaviour  mockBehaviour
		inSessionValue string
		expError       error
	}{
		{
			name: "OK",
			mockBehaviour: func(conn *redigomock.Conn, sessionValue string) {
				conn.Command("DEL", sessionValue).Expect([]byte("1"))
			},
			inSessionValue: "session:dsf8gkmw34mkdrg9",
			expError:       nil,
		},
		{
			name: "OK",
			mockBehaviour: func(conn *redigomock.Conn, sessionValue string) {
				conn.Command("DEL", sessionValue).ExpectError(fmt.Errorf("redis: not OK"))
			},
			inSessionValue: "session:dsf8gkmw34mkdrg9",
			expError:       fmt.Errorf("redis: not OK"),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehaviour(rdConn, testCase.inSessionValue)

			sessionRep := NewSessionRdRepository(rdConn)
			err := sessionRep.Delete(testCase.inSessionValue)

			assert.Equal(t, err, testCase.expError)
		})
	}
}
