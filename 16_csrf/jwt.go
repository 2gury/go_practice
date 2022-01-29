package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JwtToken struct {
	Secret []byte
}

type JwtCsrfClaims struct {
	SessionID string `json:"sid"`
	UserID uint64 `json:"uid"`
	jwt.StandardClaims
}

func NewJwtString(secret string) (*JwtToken, error) {
	return &JwtToken{Secret: []byte(secret)}, nil
}

func (tk *JwtToken) Create(s *Session) (string, error) {
	data := JwtCsrfClaims{
		SessionID: s.ID,
		UserID: s.UserID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			IssuedAt: time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, data)
	return token.SignedString(tk.Secret)
}

func (tk *JwtToken) parseSecretGetter(token *jwt.Token) (interface{}, error) {
	method, ok := token.Method.(*jwt.SigningMethodHMAC)
	if !ok || method.Alg() != "HS256" {
		return nil, fmt.Errorf("error signing method")
	}
	return tk.Secret, nil
}

func (tk *JwtToken) Check(s *Session, inputToken string) error {
	payload := &JwtCsrfClaims{}
	_, err := jwt.ParseWithClaims(inputToken, payload, tk.parseSecretGetter)
	if err != nil {
		return err
	}
	if payload.Valid() != nil {
		return fmt.Errorf("not valid token")
	}
	ok := payload.SessionID == s.ID && payload.UserID == s.UserID
	if !ok {
		return fmt.Errorf("not valid token")
	}
	return nil
}
