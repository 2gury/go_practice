package main

import "sync"

type BookInput struct {
	Title string `json:"title"`
	Price int    `json:"price"`
}

type Book struct {
	Id    int
	Title string
	Price int
}

type BookStore struct {
	books  []*Book
	mu     *sync.Mutex
	nextId int
}

func NewBookStore() *BookStore {
	return &BookStore{
		books:  []*Book{},
		mu:     &sync.Mutex{},
		nextId: 0,
	}
}

func (bs *BookStore) AddBook(book BookInput) (int, error) {
	bs.mu.Lock()
	defer bs.mu.Unlock()
	bs.books = append(bs.books, &Book{
		Id:    bs.nextId,
		Title: book.Title,
		Price: book.Price,
	})
	bs.nextId++
	return bs.nextId - 1, nil
}

func (bs *BookStore) GetBooks() (*[]*Book, error) {
	bs.mu.Lock()
	defer bs.mu.Unlock()
	return &bs.books, nil
}

func (bs *BookStore) GetBookById(id int) (*Book, error) {
	bs.mu.Lock()
	defer bs.mu.Unlock()
	for i := 0; i < len(bs.books); i++ {
		if bs.books[i].Id == id {
			return bs.books[i], nil
		}
	}
	return nil, nil
}
