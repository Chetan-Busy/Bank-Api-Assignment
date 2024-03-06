package queries

import (
	db "bankassignment/database"
	"bankassignment/models"
	"errors"

	"github.com/go-pg/pg/v10"
)

func CreateAccount(account *models.Account) (uint, error) {
	res, err := db.DB.Model(account).Insert()
	if err != nil {
		return 0, err
	}
	return uint(res.RowsAffected()), nil
}

func DeleteAccountById(id string) (uint, error) {
	account := new(models.Account)
	res, err := db.DB.Model(account).Where("account_id = ?", id).Delete()
	if err != nil {
		return 0, err
	}
	return uint(res.RowsAffected()), nil
}

func GetAccountDetailsById(id string) (*models.Account, error) {
	account := new(models.Account)
	err := db.DB.Model(account).Where("account_id =?", id).Relation("Branch").Relation("Transaction").Relation("Mapping").Relation("Mapping.Customer").Select()
	if err != nil {
		return nil, err
	}
	return account, nil
}

func UpdateAccount(account *models.Account) (uint, error) {
	res, err := db.DB.Model(account).Where("account_id = ?", account.AccountId).Update()
	if err != nil {
		return 0, err
	}
	return uint(res.RowsAffected()), nil
}

func Credit(accountId float32, amount float32, tx *pg.Tx) (uint, error) {
	account := new(models.Account)
	err := tx.Model(account).Where("account_id =?", accountId).Select()
	account.Balance += float32(amount)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	res, err2 := tx.Model(account).Where("account_id =?", accountId).Update()

	if err2 != nil {
		tx.Rollback()
		return 0, err2
	}
	return uint(res.RowsAffected()), nil
}

func Debit(accountId float32, amount float32, tx *pg.Tx) (uint, error) {
	account := new(models.Account)
	err := tx.Model(account).Where("account_id =?", accountId).Select()
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	if account.Balance < amount {
		tx.Rollback()
		return 0, errors.New("insufficient balance")
	}
	account.Balance -= float32(amount)

	res, err2 := db.DB.Model(account).Where("account_id =?", accountId).Update()

	if err2 != nil {
		tx.Rollback()
		return 0, err2
	}
	return uint(res.RowsAffected()), nil
}
