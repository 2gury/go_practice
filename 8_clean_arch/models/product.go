package models

type Product struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Price int    `json:"price"`
}

type ProductInput struct {
	Title string `json:"title"`
	Price int    `json:"price"`
}
