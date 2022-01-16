package models

type Product struct {
	Id    uint64 `json:"id" valid:"int,optional"`
	Title string `json:"title" valid:",required"`
	Price int    `json:"price" valid:"int,required"`
}
