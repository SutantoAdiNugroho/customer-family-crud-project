package repository

import "customer-family-crud-backend/domain/model"

type CustomerRepository interface {
	CreateCustomer(customer *model.Customer, familyList []*model.FamilyList) error
	GetCustomerByEmail(email string) (*model.Customer, error)
}
