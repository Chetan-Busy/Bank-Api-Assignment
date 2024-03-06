package models

type Branch struct {
	BranchId      uint        `pg:"branch_id,pk"`
	IFSCCode      string      `pg:"ifsc_code"`
	BranchAddress string      `pg:"branch_address"`
	Bank          *Bank       `pg:"rel:has-one"`
	BankID        uint        `pg:"fk:bank_id,on_delete:CASCADE,on_update:CASCADE"`
	Account       []*Account  `pg:"rel:has-many"`
	Customers     []*Customer `pg:"rel:has-many"`
}
