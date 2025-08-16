package repository

import "customer-family-crud-backend/domain/model"

type NationalityRepository interface {
	GetAllNationalities() ([]*model.Nationality, error)
}
