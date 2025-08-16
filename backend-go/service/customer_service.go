package service

import (
	"customer-family-crud-backend/domain/model"
	"customer-family-crud-backend/repository/dto"
)

type CustomerService interface {
	Create(customer *model.Customer, familyLists []*model.FamilyList) *ServiceError
	Update(customer *model.Customer, familyLists []*model.FamilyList) *ServiceError
	GetAllCustomers(page, limit int) ([]*dto.CustomerWithFamilyCount, int, *ServiceError)
	Delete(id int) *ServiceError
}
