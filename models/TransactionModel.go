package models

type Transaction struct {
	TransactionId     uint     `pg:"transaction_id,pk"`
	TransactionType   string   `pg:"transaction_type"`
	TransactionAmount int      `pg:"transaction_amount"`
	TransactionMode   string   `pg:"transaction_mode"`
	Account           *Account `pg:"rel:has-one"`
	AccountId         uint     `pg:"fk:account_id,on_delete:CASCADE,on_update:CASCADE"`
	RecieverAccountId uint     `pg:"fk:account_id,on_delete:CASCADE,on_update:CASCADE"`
}
