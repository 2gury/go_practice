package main

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

type Session struct {
	Login string
	Useragent string
}

type SessionId struct {
	Id string
}

type SessionManager struct {
	redisConn redis.Conn
}

func NewSessionManager(conn redis.Conn)  *SessionManager {
	return &SessionManager{
		redisConn: conn,
	}
}

func (sm *SessionManager) Create(in *Session) (*SessionId, error) {
	id := SessionId{RandStringRunes(sessKeyLen)}
	dataSerialized, _ := json.Marshal(in)
	mkey := "sessions:" + id.Id
	res, err := redis.String(sm.redisConn.Do("SET", mkey, dataSerialized, "EX", 86400))
	if err != nil {
		return nil, err
	}
	if res != "OK" {
		return nil, fmt.Errorf("Redis: result not OK")
	}
	return &id, nil
}

func (sm *SessionManager) Check(in *SessionId) (*Session, error) {
	mkey := "sessions:" + in.Id
	data, err := redis.Bytes(sm.redisConn.Do("GET", mkey))
	if err != nil {
		return nil, err
	}
	sess := &Session{}
	err = json.Unmarshal(data, sess)
	if err != nil {
		return nil, err
	}
	return sess, nil
}

func (sm *SessionManager) Delete(in *SessionId) (error) {
	mkey := "sessions:" + in.Id
	_, err := redis.Int(sm.redisConn.Do("DEL", mkey))
	return err
}