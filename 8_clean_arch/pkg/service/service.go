package service

import (
	"go_practice/8_clean_arch/models"
	"go_practice/8_clean_arch/pkg/repository"
)

type Service struct {
	BankService
	CityService
}

type BankService interface {
	GetBanks() ([]*models.Bank, error)
}

type CityService interface {

}

func NewService(rep *repository.Repository) *Service {
	return &Service{
		BankService: NewBankSvc(rep),
	}
}