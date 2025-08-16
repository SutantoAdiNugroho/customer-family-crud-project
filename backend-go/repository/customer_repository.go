package repository

import (
	"customer-family-crud-backend/domain/model"
	"customer-family-crud-backend/repository/dto"
)

type CustomerRepository interface {
	CreateCustomer(customer *model.Customer, familyList []*model.FamilyList) error
	GetCustomerByIdOrEmail(id *int, email *string) (*model.Customer, error)
	UpdateCustomer(customer *model.Customer, familyLists []*model.FamilyList) error
	DeleteCustomer(id int) error
	GetAllCustomers(limit, offset int) ([]*dto.CustomerWithFamilyCount, int, error)
}
