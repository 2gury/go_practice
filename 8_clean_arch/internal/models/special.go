package models

type Special struct {
	Id        uint64 `json:"id"`
	ProductId uint64 `json:"product_id"`
	Slogan    string `json:"slogan"`
	Discount  int    `json:"discount"`
}
