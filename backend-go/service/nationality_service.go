package service

import "customer-family-crud-backend/domain/model"

type NationalityService interface {
	GetAllNationalities() ([]*model.Nationality, *ServiceError)
}
