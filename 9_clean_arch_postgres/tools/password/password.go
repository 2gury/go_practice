package password

import (
	"crypto/md5"
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(inputPassword string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(inputPassword), 14)
	if err != nil {
		return "", err
	}

	return string(bytes), err
}

func VerifyPasswordAndHash(inputPassword string, hashFromDb string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashFromDb), []byte(inputPassword))
	if err != nil {
		return false
	}

	return true
}

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
