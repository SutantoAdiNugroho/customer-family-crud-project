package impl

import (
	"customer-family-crud-backend/domain/model"
	"customer-family-crud-backend/repository"
	"customer-family-crud-backend/repository/dto"
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

func (s *customerServiceImpl) GetCustomerByID(id int) (*model.Customer, *service.ServiceError) {
	extCustomer, err := s.customerRepo.GetCustomerByIdOrEmail(&id, nil)
	if err != nil {
		return nil, service.NewServiceError("Failed to get existing customer", http.StatusInternalServerError)
	}

	log.Printf("extCustomer: %v", extCustomer)

	if extCustomer == nil {
		return nil, service.NewServiceError("Customer not found", http.StatusNotFound)
	}

	return extCustomer, nil
}

func (s *customerServiceImpl) Update(customer *model.Customer, familyLists []*model.FamilyList) *service.ServiceError {
	extCustomer, err := s.GetCustomerByID(customer.CstID)
	if err != nil {
		return service.NewServiceError(err.Message, err.StatusCode)
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

func (s *customerServiceImpl) GetCustomerDetailsByID(id int) (*model.Customer, []*model.FamilyList, *service.ServiceError) {
	customer, familyLists, err := s.customerRepo.GetCustomerDetailsyByID(id)
	if err != nil {
		if err.Error() == "customer not found" {
			return nil, nil, service.NewServiceError("Customer not found", http.StatusNotFound)
		}
		return nil, nil, service.NewServiceError("Failed to get customer details", http.StatusInternalServerError)
	}

	return customer, familyLists, nil
}

func (s *customerServiceImpl) Delete(id int) *service.ServiceError {
	extCustomer, errCust := s.GetCustomerByID(id)
	if errCust != nil {
		return service.NewServiceError(errCust.Message, errCust.StatusCode)
	}

	err := s.customerRepo.DeleteCustomer(extCustomer.CstID)
	if err != nil {
		return service.NewServiceError("Failed to delete customer", http.StatusInternalServerError)
	}
	return nil
}

func (s *customerServiceImpl) GetAllCustomers(page, limit int) ([]*dto.CustomerWithFamilyCount, int, *service.ServiceError) {
	offset := (page - 1) * limit
	customers, total, err := s.customerRepo.GetAllCustomers(limit, offset)
	if err != nil {
		return nil, 0, service.NewServiceError("Failed to get all customers", http.StatusInternalServerError)
	}

	return customers, total, nil
}
