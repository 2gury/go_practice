package main

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/argon2"
	"golang.org/x/crypto/pbkdf2"
	"golang.org/x/crypto/scrypt"
	"log"
	"math/rand"
	"net/http"
)

func Root(w http.ResponseWriter, r *http.Request) {
	inputPassword := []byte(r.FormValue("pass"))
	if len(inputPassword) == 0 {
		w.Write([]byte("Incorrect input params: [pass]"))
		return
	}
	var salt []byte
	salt = append(salt, RandBytes(10)[:]...)
	w.Write([]byte(fmt.Sprintf("Salt: %x\n", salt)))
	w.Write([]byte(fmt.Sprintf("Input password: %x\n", inputPassword)))
	passwordWithSalt := append(salt[:], inputPassword[:]...)
	w.Write([]byte(fmt.Sprintf("Password with salt: %x\n\n", passwordWithSalt)))
	{
		//Generate MD5 Hash
		md5Hash := md5.New().Sum(passwordWithSalt)
		resMD5Hash := append(salt, md5Hash[:]...)
		w.Write([]byte(fmt.Sprintf("MD5: %x\n", resMD5Hash)))

		//Verify MD5 Hash
		verifySaltMD5 := resMD5Hash[:10]
		verifyPasswordWithSaltMD5 := append(verifySaltMD5, inputPassword[:]...)
		verifyMD5Hash := md5.New().Sum(verifyPasswordWithSaltMD5)
		verifyResMD5Hash := append(verifySaltMD5, verifyMD5Hash[:]...)
		w.Write([]byte(fmt.Sprintf("Compare MD5 hash: %d\n\n", bytes.Compare(verifyResMD5Hash, resMD5Hash[:]))))
	}
	{
		//Generate PBKDF2 Hash
		PBKDF2Hash := pbkdf2.Key(inputPassword, salt, 4096, 32, sha1.New)
		resPBKDF2Hash := append(salt, PBKDF2Hash[:]...)
		w.Write([]byte(fmt.Sprintf("PBKDF2: %x\n", resPBKDF2Hash)))

		//Verify PBKDF2 Hash
		verifyPBKDF2Salt := resPBKDF2Hash[:10]
		verifyPBKDF2Hash := pbkdf2.Key(inputPassword, verifyPBKDF2Salt, 4096, 32, sha1.New)
		verifyResPBKDF2Hash := append(verifyPBKDF2Salt, verifyPBKDF2Hash[:]...)
		w.Write([]byte(fmt.Sprintf("Compare PBKDF2 hash: %d\n\n", bytes.Compare(verifyResPBKDF2Hash, resPBKDF2Hash[:]))))
	}
	{
		//Generate Scrypt Hash
		scryptHash, _ := scrypt.Key(inputPassword, salt, 4096, 8, 1, 32)
		resScryptHash := append(salt, scryptHash[:]...)
		w.Write([]byte(fmt.Sprintf("Scrypt: %x\n", resScryptHash)))

		//Verify Scrypt Hash
		verifyScryptSalt := resScryptHash[:10]
		verifyScryptHash, _ := scrypt.Key(inputPassword, verifyScryptSalt, 4096, 8, 1, 32)
		verifyResScryptHash := append(salt, verifyScryptHash[:]...)
		w.Write([]byte(fmt.Sprintf("Compare Scrypt hash: %d\n\n", bytes.Compare(verifyResScryptHash, resScryptHash[:]))))
	}
	{
		//Generate Argon2 Hash
		argon2Hash := argon2.IDKey(inputPassword, salt, 1, 64 * 1024, 4, 32)
		resArgon2Hash := append(salt, argon2Hash[:]...)
		w.Write([]byte(fmt.Sprintf("Argon2: %x\n", resArgon2Hash)))

		//Verify Argon2 Hash
		verifyArgon2Salt := resArgon2Hash[:10]
		verifyArgon2Hash := argon2.IDKey(inputPassword, verifyArgon2Salt, 1, 64 * 1024, 4, 32)
		verifyResArgon2Hash := append(salt, verifyArgon2Hash[:]...)
		w.Write([]byte(fmt.Sprintf("Compare Argon2 hash: %d\n\n", bytes.Compare(verifyResArgon2Hash, resArgon2Hash[:]))))
	}
}

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/", Root)
	log.Println("Launch server at :8080 port")
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
