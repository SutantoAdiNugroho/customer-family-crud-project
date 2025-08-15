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

func (r *CustomerRepositoryImpl) GetCustomerByEmail(email string) (*model.Customer, error) {
	customer := &model.Customer{}

	query := `SELECT cst_id, nationality_id, cst_name, cst_dob, cst_phoneNum, cst_email FROM customer WHERE cst_email = $1`
	err := r.DB.QueryRow(query, email).Scan(
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
		fmt.Printf("error when get customer by email: %v", err)
		return nil, fmt.Errorf("failed to get customer by email: %w", err)
	}

	return customer, nil
}
