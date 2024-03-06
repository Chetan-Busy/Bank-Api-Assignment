package models

type Bank struct {
	BankId      uint      `pg:"bank_id,pk"`
	BankName    string    `pg:"bank_name" validate:"required"`
	BankAddress string    `pg:"bank_address" validate:"required"`
	Branches    []*Branch `pg:"rel:has-many"`
}
