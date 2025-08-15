package service

import "customer-family-crud-backend/domain/model"

type CustomerService interface {
	Create(customer *model.Customer, familyLists []*model.FamilyList) *ServiceError
}
