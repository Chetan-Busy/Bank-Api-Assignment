package models

type Mapping struct {
	Id         uint      `pg:"id,pk"`
	Customer   *Customer `pg:"rel:has-one"`
	CustomerId uint      `pg:"fk:customer_id"`
	Account    *Account  `pg:"rel:has-one"`
	AccountId  uint      `pg:"fk:account_id"`
}
