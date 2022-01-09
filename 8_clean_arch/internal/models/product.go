package models

type Product struct {
	Id    uint64 `json:"id"`
	Title string `json:"title"`
	Price int    `json:"price"`
}

type ProductInput struct {
	Title string `json:"title"`
	Price int    `json:"price"`
}
