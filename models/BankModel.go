package models

type Bank struct {
	BankId      uint      `pg:"bank_id,pk"`
	BankName    string    `pg:"bank_name"`
	BankAddress string    `pg:"bank_address"`
	Branches    []*Branch `pg:"rel:has-many"`
}
