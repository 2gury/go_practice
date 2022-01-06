package models

type Bank struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

type BankInput struct {
	Name string `json:"name"`
}
