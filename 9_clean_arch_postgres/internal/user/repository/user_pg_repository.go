package repository

import (
	"database/sql"
	"go_practice/9_clean_arch_db/internal/user"
)

type UserPgRepository struct {
	dbConn *sql.DB
}

func NewUserPgRepository(conn *sql.DB) user.UserRepository {
	return UserPgRepository{
		dbConn: conn,
	}
}
