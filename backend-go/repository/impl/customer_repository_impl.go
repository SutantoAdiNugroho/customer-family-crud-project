package impl

import (
	"customer-family-crud-backend/domain/model"
	"customer-family-crud-backend/repository"
	"database/sql"
	"fmt"
)

type CustomerRepositoryImpl struct {
	DB *sql.DB
}

func NewCustomerRepositoryImpl(db *sql.DB) repository.CustomerRepository {
	return &CustomerRepositoryImpl{DB: db}
}

func (r *CustomerRepositoryImpl) CreateCustomer(customer *model.Customer, familyList []*model.FamilyList) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	custQuery := `insert into customer (nationality_id, cst_name, cst_dob, cst_phoneNum, cst_email) values ($1, $2, $3, $4, $5) returning cst_id`
	err = tx.QueryRow(custQuery, customer.NationalityID, customer.CstName, customer.CstDob, customer.CstPhoneNum, customer.CstEmail).Scan(&customer.CstID)
	if err != nil {
		fmt.Printf("error when create customer: %v", err)
		return fmt.Errorf("failed to create customer")
	}

	familyListQuery := `insert into family_list (cst_id, fl_relation, fl_name, fl_dob) VALUES ($1, $2, $3, $4)`
	for _, fl := range familyList {
		fl.CstID = customer.CstID
		_, err := tx.Exec(familyListQuery, fl.CstID, fl.FlRelation, fl.FlName, fl.FlDob)
		if err != nil {
			fmt.Printf("error when create customer family list: %v", err)
			return fmt.Errorf("failed to create family list: %w", err)
		}
	}

	return tx.Commit()
}

func (r *CustomerRepositoryImpl) GetCustomerByIdOrEmail(id *int, email *string) (*model.Customer, error) {
	customer := &model.Customer{}
	var (
		query    string
		argParam []interface{}
	)

	if id != nil {
		query = `SELECT cst_id, nationality_id, cst_name, cst_dob, cst_phoneNum, cst_email 
		         FROM customer WHERE cst_id = $1`
		argParam = append(argParam, *id)
	} else if email != nil {
		query = `SELECT cst_id, nationality_id, cst_name, cst_dob, cst_phoneNum, cst_email 
		         FROM customer WHERE cst_email = $1`
		argParam = append(argParam, *email)
	} else {
		return nil, fmt.Errorf("no id or email provided")
	}

	err := r.DB.QueryRow(query, argParam...).Scan(
		&customer.CstID,
		&customer.NationalityID,
		&customer.CstName,
		&customer.CstDob,
		&customer.CstPhoneNum,
		&customer.CstEmail,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		fmt.Printf("error when get customer: %v", err)
		return nil, fmt.Errorf("failed to get customer: %w", err)
	}

	return customer, nil
}

func (r *CustomerRepositoryImpl) UpdateCustomer(customer *model.Customer, familyLists []*model.FamilyList) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	custQuery := `UPDATE customer SET nationality_id = $1, cst_name = $2, cst_dob = $3, cst_phoneNum = $4, cst_email = $5 WHERE cst_id = $6`
	_, err = tx.Exec(custQuery, customer.NationalityID, customer.CstName, customer.CstDob, customer.CstPhoneNum, customer.CstEmail, customer.CstID)
	if err != nil {
		return fmt.Errorf("failed to update customer: %w", err)
	}

	// delete old family list
	deleteFamilyList := `DELETE FROM family_list WHERE cst_id = $1`
	_, err = tx.Exec(deleteFamilyList, customer.CstID)
	if err != nil {
		return fmt.Errorf("failed to delete old family list: %w", err)
	}

	// insert new family list
	if len(familyLists) > 0 {
		familyListQuery := `INSERT INTO family_list (cst_id, fl_relation, fl_name, fl_dob) VALUES ($1, $2, $3, $4)`
		for _, fl := range familyLists {
			_, err := tx.Exec(familyListQuery, customer.CstID, fl.FlRelation, fl.FlName, fl.FlDob)
			if err != nil {
				return fmt.Errorf("failed to insert new family list: %w", err)
			}
		}
	}

	return tx.Commit()
}
