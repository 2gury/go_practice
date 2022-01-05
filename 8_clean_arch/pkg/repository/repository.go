package repository

import "go_practice/8_clean_arch/models"

type Repository struct {
	BankRepository
	CityRepository
}

type BankRepository interface {
	GetBanks() ([]*models.Bank, error)
}

type CityRepository interface {

}

func NewRepository(rep *MapRepository) *Repository {
	return &Repository{
		BankRepository: NewBankRep(rep),
	}
}
