package repository

import (
	"database/sql/driver"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"go_practice/9_clean_arch_db/internal/models"
	"testing"
)

func TestUserPgRepository_SelectById(t *testing.T) {
	type mockBehaviour func(mock sqlmock.Sqlmock, usrId uint64, usr *models.User)

	t.Parallel()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	testTable := []struct {
		name string
		mockBehaviour mockBehaviour
		inUserId uint64
		outUser *models.User
		expError error
	}{
		{
			name: "OK",
			mockBehaviour: func(mock sqlmock.Sqlmock, usrId uint64, usr *models.User) {
				rows := sqlmock.NewRows([]string{"id", "email", "password", "role"})
				rows.AddRow(usr.Id, usr.Email, usr.Password, usr.Role)
				mock.ExpectQuery(`SELECT`).WithArgs(usrId).WillReturnRows(rows)
			},
			inUserId: 1,
			outUser: &models.User{
				Id:    1,
				Email: "testmail@kek.ru",
				Password: "dsf32g2434g",
				Role: "user",
			},
			expError: nil,
		},
		{
			name: "Error: sql error",
			mockBehaviour: func(mock sqlmock.Sqlmock, usrId uint64, usr *models.User) {
				mock.ExpectQuery(`SELECT`).WithArgs(usrId).WillReturnError(fmt.Errorf("sql error"))
			},
			inUserId: 1,
			outUser: nil,
			expError: fmt.Errorf("sql error"),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			userRep := NewUserPgRepository(db)
			testCase.mockBehaviour(mock, testCase.inUserId, testCase.outUser)
			userFromDb, err := userRep.SelectById(testCase.inUserId)

			assert.Equal(t, userFromDb, testCase.outUser)
			assert.Equal(t, err, testCase.expError)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("were met expectation: %s", err)
			}
		})
	}
}

func TestUserPgRepository_Insert(t *testing.T) {
	type mockBehaviour func(mock sqlmock.Sqlmock, usr *models.User, usrId uint64)

	t.Parallel()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	testTable := []struct {
		name string
		mockBehaviour mockBehaviour
		inUser *models.User
		outUserId uint64
		expError error
	}{
		{
			name: "OK",
			mockBehaviour: func(mock sqlmock.Sqlmock, usr *models.User, usrId uint64) {
				mock.ExpectBegin()
				insertedUsrId := sqlmock.NewRows([]string{"id"}).AddRow(usrId)
				mock.ExpectQuery(`INSERT INTO users`).
					WithArgs(usr.Email, usr.Password).
					WillReturnRows(insertedUsrId)
				mock.ExpectCommit()
			},
			inUser: &models.User{
				Email: "testmail@kek.ru",
				Password: "dsf32g2434g",
				Role: "user",
			},
			outUserId: 1,
			expError: nil,
		},
		{
			name: "Error: sql  error",
			mockBehaviour: func(mock sqlmock.Sqlmock, usr *models.User, usrId uint64) {
				mock.ExpectBegin()
				mock.ExpectQuery(`INSERT INTO users`).
					WithArgs(usr.Email, usr.Password).
					WillReturnError(fmt.Errorf("sql error"))
				mock.ExpectRollback()
			},
			inUser: &models.User{
				Email: "testmail@kek.ru",
				Password: "dsf32g2434g",
				Role: "user",
			},
			outUserId: 0,
			expError: fmt.Errorf("sql error"),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			userRep := NewUserPgRepository(db)
			testCase.mockBehaviour(mock, testCase.inUser, testCase.outUserId)
			lastId, err := userRep.Insert(testCase.inUser)

			assert.Equal(t, lastId, testCase.outUserId)
			assert.Equal(t, err, testCase.expError)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("were met expectation: %s", err)
			}
		})
	}
}

func TestUserPgRepository_SelectByEmail(t *testing.T) {
	type mockBehaviour func(mock sqlmock.Sqlmock, usrEmail string, usr *models.User)

	t.Parallel()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	testTable := []struct {
		name string
		mockBehaviour mockBehaviour
		inUserEmail string
		outUser *models.User
		expError error
	}{
		{
			name: "OK",
			mockBehaviour: func(mock sqlmock.Sqlmock, usrEmail string, usr *models.User) {
				rows := sqlmock.NewRows([]string{"id", "email", "password", "role"})
				rows.AddRow(usr.Id, usr.Email, usr.Password, usr.Role)
				mock.ExpectQuery(`SELECT`).WithArgs(usrEmail).WillReturnRows(rows)
			},
			inUserEmail: "testmail@kek.ru",
			outUser: &models.User{
				Id:    1,
				Email: "testmail@kek.ru",
				Password: "dsf32g2434g",
				Role: "user",
			},
			expError: nil,
		},
		{
			name: "Error: sql error",
			mockBehaviour: func(mock sqlmock.Sqlmock, usrEmail string, usr *models.User) {
				mock.ExpectQuery(`SELECT`).WithArgs(usrEmail).WillReturnError(fmt.Errorf("sql error"))
			},
			inUserEmail: "testmail@kek.ru",
			outUser: nil,
			expError: fmt.Errorf("sql error"),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			userRep := NewUserPgRepository(db)
			testCase.mockBehaviour(mock, testCase.inUserEmail, testCase.outUser)
			userFromDb, err := userRep.SelectByEmail(testCase.inUserEmail)

			assert.Equal(t, userFromDb, testCase.outUser)
			assert.Equal(t, err, testCase.expError)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("were met expectation: %s", err)
			}
		})
	}
}

func TestUserPgRepository_UpdatePassword(t *testing.T) {
	type mockBehaviour func(mock sqlmock.Sqlmock, usrPass string, usrId uint64)

	t.Parallel()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	testTable := []struct {
		name string
		mockBehaviour mockBehaviour
		inUser *models.User
		outUser *models.User
		expError error
	}{
		{
			name: "OK",
			mockBehaviour: func(mock sqlmock.Sqlmock, usrPass string, usrId uint64) {
				mock.ExpectBegin()
				mock.ExpectExec(`UPDATE users`).
					WithArgs(usrPass, usrId).
					WillReturnResult(driver.ResultNoRows)
				mock.ExpectCommit()
			},
			inUser: &models.User{
				Id: 1,
				Email: "testmail@kek.ru",
				Password: "8dfs8seg9dfg",
			},
			expError: nil,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			userRep := NewUserPgRepository(db)
			testCase.mockBehaviour(mock, testCase.inUser.Password, testCase.inUser.Id)
			err := userRep.UpdatePassword(testCase.inUser)

			assert.Equal(t, err, testCase.expError)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("were met expectation: %s", err)
			}
		})
	}
}

func TestUserPgRepository_DeleteById(t *testing.T) {
	type mockBehaviour func(mock sqlmock.Sqlmock, usrId uint64)

	t.Parallel()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	testTable := []struct {
		name string
		mockBehaviour mockBehaviour
		inUserId uint64
		expError error
	}{
		{
			name: "OK",
			mockBehaviour: func(mock sqlmock.Sqlmock, usrId uint64) {
				mock.ExpectBegin()
				mock.ExpectExec(`DELETE FROM users`).
					WithArgs(usrId).
					WillReturnResult(driver.ResultNoRows)
				mock.ExpectCommit()
			},
			inUserId: 1,
			expError: nil,
		},
		{
			name: "OK",
			mockBehaviour: func(mock sqlmock.Sqlmock, usrId uint64) {
				mock.ExpectBegin()
				mock.ExpectExec(`DELETE FROM users`).
					WithArgs(usrId).
					WillReturnResult(driver.ResultNoRows)
				mock.ExpectCommit()
			},
			inUserId: 1,
			expError: nil,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			userRep := NewUserPgRepository(db)
			testCase.mockBehaviour(mock, testCase.inUserId)
			err := userRep.DeleteById(testCase.inUserId)

			assert.Equal(t, err, testCase.expError)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("were met expectation: %s", err)
			}
		})
	}
}