package repository

import "go_practice/8_clean_arch/models"

type Repository struct {
	BankRepository
}

type BankRepository interface {
	GetBanks() ([]*models.Bank, error)
	GetBankById(id int) (*models.Bank, error)
	AddBank(bank models.BankInput) (int, error)
}

func NewRepository(rep *InRuntimeRepository) *Repository {
	return &Repository{
		BankRepository: NewBankRep(rep),
	}
}
