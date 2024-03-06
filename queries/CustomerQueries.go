package queries

import (
	db "bankassignment/database"
	"bankassignment/models"
)

func CreateCustomer(customer *models.Customer) (uint, error) {
	res, err := db.DB.Model(customer).Insert()
	if err != nil {
		return 0, err
	}
	return uint(res.RowsAffected()), nil
}

func GetCustomerDetailsById(id string) (*models.Customer, error) {
	customer := new(models.Customer)
	err := db.DB.Model(customer).Where("customer_id = ?", id).Relation("Branch").Relation("Mapping").Relation("Mapping.Account").Select()
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func DeleteCustomerById(id string) (uint, error) {
	customer := new(models.Customer)
	res, err := db.DB.Model(customer).Where("customer_id = ?", id).Delete()
	if err != nil {
		return 0, err
	}
	return uint(res.RowsAffected()), nil
}

func UpdateCustomer(customer *models.Customer) (uint, error) {
	res, err := db.DB.Model(customer).Where("customer_id =?", customer.CustomerId).Update()
	if err != nil {
		return 0, err
	}
	return uint(res.RowsAffected()), nil
}
