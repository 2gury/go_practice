package models

type User struct {
	Id       uint64 `json:"id"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Role     string `json:"-"`
}
