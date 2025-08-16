package impl

import (
	"customer-family-crud-backend/domain/model"
	"customer-family-crud-backend/repository"
	"customer-family-crud-backend/service"
	"log"
	"net/http"
)

type customerServiceImpl struct {
	customerRepo repository.CustomerRepository
}

func NewCustomerService(customerRepo repository.CustomerRepository) service.CustomerService {
	return &customerServiceImpl{customerRepo: customerRepo}
}

func (s *customerServiceImpl) Create(customer *model.Customer, familyLists []*model.FamilyList) *service.ServiceError {
	extCustomer, err := s.customerRepo.GetCustomerByIdOrEmail(nil, &customer.CstEmail)
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

func (s *customerServiceImpl) Update(customer *model.Customer, familyLists []*model.FamilyList) *service.ServiceError {
	extCustomer, err := s.customerRepo.GetCustomerByIdOrEmail(&customer.CstID, nil)
	if err != nil {
		return service.NewServiceError("Failed to get existing customer", http.StatusInternalServerError)
	}

	log.Printf("extCustomer: %v", extCustomer)

	if extCustomer == nil {
		return service.NewServiceError("Customer not found", http.StatusNotFound)
	}

	if extCustomer.CstEmail != customer.CstEmail {
		extCustomerByEmail, err := s.customerRepo.GetCustomerByIdOrEmail(nil, &customer.CstEmail)
		if err != nil {
			return service.NewServiceError("Failed to check customer email", http.StatusInternalServerError)
		}
		if extCustomerByEmail != nil {
			return service.NewServiceError("This email is already user by another user", http.StatusConflict)
		}
	}

	if err := s.customerRepo.UpdateCustomer(customer, familyLists); err != nil {
		return service.NewServiceError("Gagal memperbarui data customer", http.StatusInternalServerError)
	}

	return nil
}
