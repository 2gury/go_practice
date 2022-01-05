package repository

import "go_practice/8_clean_arch/models"

type BankRep struct {
	mp *MapRepository
}

func NewBankRep(mapRep *MapRepository) *BankRep {
	return &BankRep{
		mp: mapRep,
	}
}

func (brp *BankRep) GetBanks() ([]*models.Bank, error) {
	return brp.mp.Banks, nil
}

