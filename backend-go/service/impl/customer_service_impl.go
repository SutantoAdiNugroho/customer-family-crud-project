package impl

import (
	"customer-family-crud-backend/domain/model"
	"customer-family-crud-backend/repository"
	"customer-family-crud-backend/service"
	"net/http"
)

type customerServiceImpl struct {
	customerRepo repository.CustomerRepository
}

func NewCustomerService(customerRepo repository.CustomerRepository) service.CustomerService {
	return &customerServiceImpl{customerRepo: customerRepo}
}

func (s *customerServiceImpl) Create(customer *model.Customer, familyLists []*model.FamilyList) *service.ServiceError {
	extCustomer, err := s.customerRepo.GetCustomerByEmail(customer.CstEmail)
	if err != nil {
		return service.NewServiceError(err.Error(), http.StatusInternalServerError)
	}

	if extCustomer != nil {
		return service.NewServiceError("Email already registered", http.StatusBadRequest)
	}

	if errSave := s.customerRepo.CreateCustomer(customer, familyLists); errSave != nil {
		return service.NewServiceError("Failed to save customer", http.StatusInternalServerError)
	}

	return nil
}
