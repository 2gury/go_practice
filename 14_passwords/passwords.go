package main

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/argon2"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/pbkdf2"
	"golang.org/x/crypto/scrypt"
	"log"
	"math/rand"
	"net/http"
)

func Root(w http.ResponseWriter, r *http.Request) {
	inputPassword := []byte(r.FormValue("pass"))
	if len(inputPassword) == 0 {
		w.Write([]byte("Incorrect input params: [ pass]"))
		return
	}
	var salt []byte
	salt = append(salt, RandBytes(10)[:]...)
	w.Write([]byte(fmt.Sprintf("Salt: %x\n", salt)))
	w.Write([]byte(fmt.Sprintf("Input password: %x\n", inputPassword)))
	passwordWithSalt := append(salt[:], inputPassword[:]...)
	w.Write([]byte(fmt.Sprintf("Password with salt: %x\n\n", passwordWithSalt)))

	md5Hash := md5.New().Sum(passwordWithSalt)
	var resMD5Hash []byte
	resMD5Hash = append(salt, md5Hash[:]...)
	w.Write([]byte(fmt.Sprintf("MD5: %x\n", resMD5Hash)))

	bcryptHash, err := bcrypt.GenerateFromPassword(passwordWithSalt, 15)
	var resBcryptHash []byte
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Bcrypt: %s\n", err.Error())))
	} else {
		resBcryptHash = append(salt, bcryptHash[:]...)
		w.Write([]byte(fmt.Sprintf("Bcrypt: %x\n", resBcryptHash)))
	}

	PBKDF2Hash := pbkdf2.Key(inputPassword, salt, 4096, 32, sha1.New)
	w.Write([]byte(fmt.Sprintf("PBKDF2: %x\n", PBKDF2Hash)))

	scryptHash, err := scrypt.Key(inputPassword, salt, 4096, 8, 1, 32)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Scrypt: %s\n", err.Error())))
	} else {
		w.Write([]byte(fmt.Sprintf("Scrypt: %x\n", scryptHash)))
	}

	argon2Hash := argon2.IDKey(inputPassword, salt, 1, 64 * 1024, 4, 32)
	w.Write([]byte(fmt.Sprintf("Argon2: %x\n\n", argon2Hash)))

	verifySaltMD5 := resMD5Hash[:10]
	verifyPasswordWithSaltMD5 := append(verifySaltMD5, inputPassword[:]...)

	verifyMD5Hash := md5.New().Sum(verifyPasswordWithSaltMD5)
	verifyResMD5Hash := append(verifySaltMD5, verifyMD5Hash[:]...)
	w.Write([]byte(fmt.Sprintf("Compare MD5 hash: %d\n", bytes.Compare(verifyResMD5Hash, resMD5Hash[:]))))

	verifySaltBcrypt := resBcryptHash[:10]
	verifyPasswordWithSaltBcrypt := append(verifySaltBcrypt, inputPassword[:]...)
	verifyBcryptHash, err := bcrypt.GenerateFromPassword(verifyPasswordWithSaltBcrypt, 15)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Compare Bcrypt hash: %s\n", err.Error())))
	} else {
		resBcryptHash = append(salt, verifyBcryptHash[:]...)
		w.Write([]byte(fmt.Sprintf("Compare MD5 hash: %d\n", bytes.Compare(verifyBcryptHash, resBcryptHash[:]))))
	}
}

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/", Root)
	log.Println("Launch server att :8080 port")
	http.ListenAndServe(":8080", mux)
}

var letterRunes = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandBytes(n int) []byte {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return b
}