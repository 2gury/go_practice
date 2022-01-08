package models

type Special struct {
	Id int `json:"id"`
	ProductId int `json:"product_id"`
	Slogan string `json:"slogan"`
	Discount string `json:"discount"`
}
