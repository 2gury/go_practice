package repository

import "go_practice/8_clean_arch/models"

type MapRepository struct {
	Banks []*models.Bank
	Cities []*models.City
}

func NewMapRepository() *MapRepository {
	return &MapRepository{
		Banks: []*models.Bank{
			&models.Bank{
				Id:     0,
				Name:   "Sberbank",
				Cities: []int{0, 1, 2},
			},
			&models.Bank{
				Id:     1,
				Name:   "Orenburg",
				Cities: []int{2},
			},
			&models.Bank{
				Id:     2,
				Name:   "VTB",
				Cities: []int{0, 1},
			},
		},
		Cities: []*models.City{
			&models.City{
				Id:     0,
				Name:   "Moskva",
			},
			&models.City{
				Id:     1,
				Name:   "Sankt-Peterburg",
			},
			&models.City{
				Id:     2,
				Name:   "Orenburg",
			},
		},
	}
}
