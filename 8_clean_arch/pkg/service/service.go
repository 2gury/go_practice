package service

import (
	"go_practice/8_clean_arch/models"
	"go_practice/8_clean_arch/pkg/repository"
)

type Service struct {
	BankService
}

type BankService interface {
	GetBanks() ([]*models.Bank, error)
	GetBankById(id int) (*models.Bank, error)
	AddBank(bank models.BankInput) (int, error)
}

func NewService(rep *repository.Repository) *Service {
	return &Service{
		BankService: NewBankSvc(rep),
	}
}