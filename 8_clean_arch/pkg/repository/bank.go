package repository

import "go_practice/8_clean_arch/models"

type BankRep struct {
	rep *InRuntimeRepository
}

func NewBankRep(mapRep *InRuntimeRepository) *BankRep {
	return &BankRep{
		rep: mapRep,
	}
}

func (brp *BankRep) GetBanks() ([]*models.Bank, error) {
	return brp.rep.Banks, nil
}

func (brp *BankRep) GetBankById(id int) (*models.Bank, error) {
	for i := 0; i < len(brp.rep.Banks); i++ {
		if brp.rep.Banks[i].Id == id {
			return brp.rep.Banks[i], nil
		}
	}
	return nil, nil
}

func (brp *BankRep) AddBank(bank models.BankInput) (int, error) {
	brp.rep.mu.Lock()
	defer brp.rep.mu.Unlock()
	brp.rep.Banks = append(brp.rep.Banks, &models.Bank{
		Id: brp.rep.nextId,
		Name: bank.Name,
	})
	brp.rep.nextId++
	return brp.rep.nextId+ - 1, nil
}

