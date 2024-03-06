package queries

import (
	db "bankassignment/database"
	"bankassignment/models"
	"errors"

	"github.com/go-pg/pg/v10"
)

func CreateMapping(mapping *models.Mapping, CustomerId uint, AccountId uint, tx *pg.Tx) (uint, error) {
	account := new(models.Account)
	err := tx.Model(account).Where("account_id = ? ", AccountId).Select()
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	customer := new(models.Customer)
	err2 := tx.Model(customer).Where("customer_id = ? ", CustomerId).Select()
	if err2 != nil {
		tx.Rollback()
		return 0, err
	}
	if account.BranchId != customer.BranchId {
		tx.Rollback()
		return 0, errors.New("branch id does not match")
	}
	res, err := tx.Model(mapping).Insert()
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return uint(res.RowsAffected()), nil
}

func DeleteMapping(id string) (uint, error) {
	mapping := new(models.Mapping)
	res, err := db.DB.Model(mapping).Where("id = ?", id).Delete()
	if err != nil {
		return 0, err
	}
	return uint(res.RowsAffected()), nil
}

func UpdateMapping(mapping *models.Mapping) (uint, error) {
	res, err := db.DB.Model(mapping).Where("id = ?", mapping.Id).Update()
	if err != nil {
		return 0, nil
	}
	return uint(res.RowsAffected()), err
}
