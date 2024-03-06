package queries

import (
	db "bankassignment/database"
	"bankassignment/models"

	"github.com/go-pg/pg/v10"
)

func CreateTransaction(transaction *models.Transaction, tx *pg.Tx) (uint, error) {
	res, err := tx.Model(transaction).Insert()
	if err != nil {
		return 0, err
	}
	return uint(res.RowsAffected()), err
}

func GetTransactionByAccount(id string) ([]models.Transaction, error) {
	transactions := new([]models.Transaction)
	err := db.DB.Model(transactions).Where("account_id= ?", id).Select()
	if err != nil {
		return nil, err
	}
	return *transactions, nil
}

func GetTransactionDetailsById(id string) (*models.Transaction, error) {
	transaction := new(models.Transaction)
	err := db.DB.Model(transaction).Where("transaction_id = ?", id).Select()
	if err != nil {
		return nil, err
	}
	return transaction, err
}
