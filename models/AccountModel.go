package models

type Account struct {
	AccountId         uint           `pg:"account_id,pk"`
	AccountNumber     uint           `pg:"account_number"`
	AccountHolderName string         `pg:"account_holder_name"`
	Balance           float32        `pg:"balance"`
	AccountType       string         `pg:"account_type"`
	OpeningDate       string         `pg:"account_opening_date"`
	Branch            *Branch        `pg:"rel:has-one"`
	BranchId          uint           `pg:"fk:branch_id,on_delete:CASCADE,on_update:CASCADE"`
	Transaction       []*Transaction `pg:"rel:has-many"`
	Mapping           []*Mapping     `pg:"rel:has-many"`
}
