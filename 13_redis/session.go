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
	id := &SessionId{RandStringRunes(sessKeyLen)}
	mkey := "sessions:" + id.Id
	sess, err := json.Marshal(in)
	if err != nil {
		return nil, err
	}
	res, err := redis.String(sm.redisConn.Do("SET", mkey, sess, "EX", 86400))
	if err != nil {
		return nil, err
	}
	if res != "OK" {
		return nil, fmt.Errorf("redis: not OK")
	}
	return id, nil
}

func (sm *SessionManager) Check(in *SessionId) (*Session, error) {
	mkey := "sessions:" + in.Id
	bytes, err := redis.Bytes(sm.redisConn.Do("GET", mkey))
	session := &Session{}
	err = json.Unmarshal(bytes, session)
	if err != nil {
		return nil, err
	}
	return session, nil
}

func (sm *SessionManager) Delete(in *SessionId) (error) {
	mkey := "sessions:" + in.Id
	res, err := redis.String(sm.redisConn.Do("DEL", mkey))
	if err != nil {
		return err
	}
	if res != "OK" {
		return fmt.Errorf("redis: not OK")
	}
	return nil
}