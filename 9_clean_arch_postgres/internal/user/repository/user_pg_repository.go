package repository

import (
	"context"
	"database/sql"
	"go_practice/9_clean_arch_db/internal/models"
	"go_practice/9_clean_arch_db/internal/user"
	"log"
)

type UserPgRepository struct {
	dbConn *sql.DB
}

func NewUserPgRepository(conn *sql.DB) user.UserRepository {
	return &UserPgRepository{
		dbConn: conn,
	}
}

func (r *UserPgRepository) SelectById(id uint64) (*models.User, error) {
	usr := &models.User{}

	err := r.dbConn.QueryRow(
		`SELECT id, email, password, role FROM users
                WHERE id = $1`, id).
		Scan(&usr.Id, &usr.Email, &usr.Password, &usr.Role)
	if err != nil {
		return nil, err
	}

	return usr, nil
}

func (r *UserPgRepository) Insert(usr *models.User) (uint64, error) {
	tx, err := r.dbConn.BeginTx(context.Background(), &sql.TxOptions{})
	if err != nil {
		return 0, err
	}
	var lastId int64
	err = tx.QueryRow(
		`INSERT INTO users(email, password)
			    VALUES ($1, $2) RETURNING id`,
		usr.Email, usr.Password).Scan(&lastId)
	if err != nil {
		if rollBackError := tx.Rollback(); rollBackError != nil {
			log.Fatal(rollBackError.Error())
		}
		return 0, err
	}
	if err := tx.Commit(); err != nil {
		return 0, err
	}

	return uint64(lastId), nil
}

func (r *UserPgRepository) SelectByEmail(email string) (*models.User, error) {
	usr := &models.User{}

	err := r.dbConn.QueryRow(
		`SELECT id, email, password, role FROM users
                WHERE email = $1`, email).
		Scan(&usr.Id, &usr.Email, &usr.Password, &usr.Role)
	if err != nil {
		return nil, err
	}

	return usr, nil
}

func (r *UserPgRepository) UpdatePassword(usr *models.User) error {
	tx, err := r.dbConn.BeginTx(context.Background(), &sql.TxOptions{})
	if err != nil {
		return err
	}
	_, err = tx.Exec(
		`UPDATE users
			    SET password=$1
			    WHERE id=$2`,
		usr.Password, usr.Id)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			log.Fatal(rollbackErr)
		}
		return nil
	}
	if err := tx.Commit(); err != nil {
		return nil
	}

	return nil
}

func (r *UserPgRepository) DeleteById(id uint64) error {
	tx, err := r.dbConn.BeginTx(context.Background(), &sql.TxOptions{})
	if err != nil {
		return err
	}
	_, err = tx.Exec(
		`DELETE FROM users
			    WHERE id=$1 `, id)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			log.Fatal(rollbackErr)
		}
		return nil
	}
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
