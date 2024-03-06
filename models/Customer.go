package models

type Customer struct {
	CustomerId      uint       `pg:"customer_id,pk"`
	CustomerName    string     `pg:"customer_name"`
	CustomerAddress string     `pk:"customer_address"`
	CustomerAge     int        `pg:"customer_age"`
	CustomerPan     uint       `pg:"customer_pan"`
	Branch          *Branch    `pg:"rel:has-one"`
	BranchId        uint       `pg:"fk:branch-id,on_delete:CASCADE,on_update:CASCADE"`
	Mapping         []*Mapping `pg:"rel:has-many"`
}
