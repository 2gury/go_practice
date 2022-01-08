package repository

import (
	"go_practice/8_clean_arch/models"
	"sync"
)

type InRuntimeRepository struct {
	Products []*models.Product
	mu *sync.Mutex
	nextId int
}

func NewInRuntimeRepository() *InRuntimeRepository {
	return &InRuntimeRepository{
		Products: []*models.Product{
			&models.Product{
				Id:     0,
				Title:   "Ипотека",
				Price: 1000,
			},
			&models.Product{
				Id:     1,
				Title:  "Кредит",
				Price: 500,
			},
			&models.Product{
				Id:     2,
				Title:   "Вклад",
				Price: 100,
			},
		},
		nextId: 3,
		mu: &sync.Mutex{},
	}
}
