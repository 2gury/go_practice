package repository

import (
	"go_practice/8_clean_arch/models"
	"sync"
)

type InRuntimeRepository struct {
	Banks []*models.Bank
	mu *sync.Mutex
	nextId int
}

func NewInRuntimeRepository() *InRuntimeRepository {
	return &InRuntimeRepository{
		Banks: []*models.Bank{
			&models.Bank{
				Id:     0,
				Name:   "Sberbank",
			},
			&models.Bank{
				Id:     1,
				Name:   "Orenburg",
			},
			&models.Bank{
				Id:     2,
				Name:   "VTB",
			},
		},
		nextId: 3,
		mu: &sync.Mutex{},
	}
}
