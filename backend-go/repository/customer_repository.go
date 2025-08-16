package repository

import "customer-family-crud-backend/domain/model"

type CustomerRepository interface {
	CreateCustomer(customer *model.Customer, familyList []*model.FamilyList) error
	GetCustomerByIdOrEmail(id *int, email *string) (*model.Customer, error)
	UpdateCustomer(customer *model.Customer, familyLists []*model.FamilyList) error
}
