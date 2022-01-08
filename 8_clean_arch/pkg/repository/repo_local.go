package repository

import (
	"go_practice/8_clean_arch/models"
	"sync"
)

type LocalRepository struct {
	Products []*models.Product
	Specials []*models.Special
	mu *sync.Mutex
}

func NewLocalRepository() *LocalRepository {
	return &LocalRepository{
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
		Specials: []*models.Special{
			&models.Special{
				Id:     0,
				ProductId: 2,
				Slogan: "Дешевле всего",
				Discount: 15,
			},
			&models.Special{
				Id:     1,
				ProductId: 1,
				Slogan: "Очень дорого",
				Discount: 5,
			},
		},
		mu: &sync.Mutex{},
	}
}
