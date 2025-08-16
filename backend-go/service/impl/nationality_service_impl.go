package impl

import (
	"customer-family-crud-backend/domain/model"
	"customer-family-crud-backend/repository"
	"customer-family-crud-backend/service"
	"net/http"
)

type NationalityService interface {
	GetAllNationalities() ([]*model.Nationality, *service.ServiceError)
}

type nationalityServiceImpl struct {
	nationalityRepo repository.NationalityRepository
}

func NewNationalityService(nationalityRepo repository.NationalityRepository) NationalityService {
	return &nationalityServiceImpl{nationalityRepo: nationalityRepo}
}

func (s *nationalityServiceImpl) GetAllNationalities() ([]*model.Nationality, *service.ServiceError) {
	nationalities, err := s.nationalityRepo.GetAllNationalities()
	if err != nil {
		return nil, service.NewServiceError("Failed to get nationalities", http.StatusInternalServerError)
	}
	return nationalities, nil
}
