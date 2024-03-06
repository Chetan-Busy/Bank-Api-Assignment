package queries

import (
	db "bankassignment/database"
	"bankassignment/models"
)

func CreateBank(bank *models.Bank) (uint, error) {
	_, err := db.DB.Model(bank).Returning("bank_id").Insert()
	if err != nil {
		return 0, err
	}

	return bank.BankId, nil
}

func GetAllBanks() ([]*models.Bank, error) {
	var banks []*models.Bank
	err := db.DB.Model(&banks).Select()
	if err != nil {
		return nil, err
	}
	return banks, nil
}

func UpdateBank(bank *models.Bank) (uint, error) {
	res, err := db.DB.Model(bank).Where("bank_id =?", bank.BankId).Update()
	if err != nil {
		return 0, err
	}
	return uint(res.RowsAffected()), nil
}

func DeleteAllBanks() (uint, error) {
	var bank *models.Bank
	res, err := db.DB.Model(bank).Where("true").Delete()
	if err != nil {
		return 0, err
	}
	return uint(res.RowsAffected()), nil
}

func DeleteBankById(id string) (uint, error) {
	var bank *models.Bank
	res, err := db.DB.Model(bank).Where("bank_id = ?", id).Delete()
	if err != nil {
		return 0, err
	}
	return uint(res.RowsAffected()), nil
}

func GetAllBanksWithBranches() ([]*models.Bank, error) {
	var banks []*models.Bank
	err := db.DB.Model(&banks).Relation("Branches").Select()
	if err != nil {
		return nil, err
	}
	return banks, err
}

func GetBankDetailsById(id string) (*models.Bank, error) {
	bank := new(models.Bank)
	err := db.DB.Model(bank).Relation("Branches").Where("bank_id = ?", id).Select()
	if err != nil {
		return nil, err
	}
	return bank, nil
}
