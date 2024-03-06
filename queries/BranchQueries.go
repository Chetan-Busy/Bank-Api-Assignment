package queries

import (
	db "bankassignment/database"
	"bankassignment/models"
)

func CreateBranch(branch *models.Branch) (uint, error) {
	res, err := db.DB.Model(branch).Insert()
	if err != nil {
		return 0, err
	}
	return uint(res.RowsAffected()), nil
}

func DeleteBranchById(id string) (uint, error) {
	var branch *models.Branch
	res, err := db.DB.Model(branch).Where("branch_id = ?", id).Delete()
	if err != nil {
		return 0, err
	}
	return uint(res.RowsAffected()), nil
}

func GetBranchDetailWithAllAccountAndCustomerDetails(id string) (*models.Branch, error) {
	branch := new(models.Branch)
	err := db.DB.Model(branch).Relation("Account").Relation("Customers").Relation("Bank").Where("branch_id = ?", id).Select()
	if err != nil {
		return nil, err
	}
	return branch, nil
}

func UpdateBranch(branch *models.Branch) (uint, error) {
	res, err := db.DB.Model(branch).Where("branch_id = ?", branch.BranchId).Update()
	if err != nil {
		return 0, err
	}
	return uint(res.RowsAffected()), nil
}
