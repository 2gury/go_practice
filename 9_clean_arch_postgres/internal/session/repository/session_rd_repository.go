package repository

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"go_practice/9_clean_arch_db/internal/models"
	"go_practice/9_clean_arch_db/internal/session"
)

type SessionRdRepository struct {
	rdConn redis.Conn
}

func NewSessionRdRepository(conn redis.Conn) session.SessionRepository {
	return &SessionRdRepository{
		rdConn: conn,
	}
}

func (r *SessionRdRepository) Create(session *models.Session) error {

	mkey := "sessions:" + session.Value
	sess, err := json.Marshal(session)
	if err != nil {
		return err
	}

	res, err := redis.String(r.rdConn.Do("SET", mkey, sess, "EX", session.GetTime()))
	if err != nil {
		return err
	}
	if res != "OK" {
		return fmt.Errorf("redis: not OK")
	}

	return nil
}

func (r *SessionRdRepository) Get(sessValue string) (*models.Session, error) {
	mkey := "sessions:" + sessValue
	sess := &models.Session{}

	bytes, err := redis.Bytes(r.rdConn.Do("GET", mkey))
	err = json.Unmarshal(bytes, sess)
	if err != nil {
		return nil, err
	}

	return sess, nil
}

func (r *SessionRdRepository) Delete(sessValue string) error {
	mkey := "sessions:" + sessValue

	_, err := redis.Int(r.rdConn.Do("DEL", mkey))
	if err != nil {
		return err
	}

	return nil
}