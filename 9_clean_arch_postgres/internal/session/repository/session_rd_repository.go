package repository

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"go_practice/9_clean_arch_db/internal/models"
	"go_practice/9_clean_arch_db/internal/session"
	"go_practice/9_clean_arch_db/tools/password"
)

type SessionRdRepository struct {
	rdConn redis.Conn
}

func NewSessionRdRepository(conn redis.Conn) session.SessionRepository {
	return &SessionRdRepository{
		rdConn: conn,
	}
}

func (r *SessionRdRepository) Create(usrId uint64) (*models.Session, error) {
	inputSess := models.NewSession(usrId)
	inputSess.Value = password.GetMD5Hash(inputSess.Value)
	mkey := "sessions:" + inputSess.Value
	sess, err := json.Marshal(inputSess)
	if err != nil {
		return nil, err
	}
	res, err := redis.String(r.rdConn.Do("SET", mkey, sess, "EX", inputSess.GetTime()))
	if err != nil {
		return nil, err
	}
	if res != "OK" {
		return nil, fmt.Errorf("redis: not OK")
	}
	return inputSess, nil
}

func (r *SessionRdRepository) Check(sessValue string) (*models.Session, error) {
	mkey := "sessions:" + sessValue
	bytes, err := redis.Bytes(r.rdConn.Do("GET", mkey))
	sess := &models.Session{}
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