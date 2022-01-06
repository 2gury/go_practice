package service

import (
	"go_practice/8_clean_arch/models"
	"go_practice/8_clean_arch/pkg/repository"
)

type BankSvc struct {
	rep *repository.Repository
}

func NewBankSvc(repo *repository.Repository) *BankSvc {
	return &BankSvc{
		rep: repo,
	}
}

func (bsv *BankSvc) GetBanks() ([]*models.Bank, error) {
	return bsv.rep.GetBanks()
}

func (bsv *BankSvc) GetBankById(id int) (*models.Bank, error) {
	return bsv.rep.GetBankById(id)
}

func (bsv *BankSvc) AddBank(bank models.BankInput) (int, error) {
	return bsv.rep.AddBank(bank)
}
