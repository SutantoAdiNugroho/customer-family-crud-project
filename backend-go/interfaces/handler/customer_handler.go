package handler

import (
	"customer-family-crud-backend/domain/model"
	"customer-family-crud-backend/service"
	"customer-family-crud-backend/utils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type CustomerHandler struct {
	customerService service.CustomerService
}

func NewCustomerHandler(customerService service.CustomerService) *CustomerHandler {
	return &CustomerHandler{customerService: customerService}
}

func (h *CustomerHandler) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var reqBody struct {
		Customer   *model.Customer     `json:"customer"`
		FamilyList []*model.FamilyList `json:"family_list"`
	}

	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	if reqBody.FamilyList == nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Family list must be included", nil)
		return
	}

	errSave := h.customerService.Create(reqBody.Customer, reqBody.FamilyList)
	if errSave != nil {
		utils.ErrorResponse(w, errSave.StatusCode, errSave.Message, nil)
		return
	}

	utils.SuccessResponse(w, http.StatusCreated, "Customer successfully created", nil)
}

func (h *CustomerHandler) UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	idCustomer, err := strconv.Atoi(idStr)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid customer id", nil)
		return
	}

	var requestBody struct {
		Customer   *model.Customer     `json:"customer"`
		FamilyList []*model.FamilyList `json:"family_list"`
	}

	err = json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	requestBody.Customer.CstID = idCustomer
	serviceErr := h.customerService.Update(requestBody.Customer, requestBody.FamilyList)
	if serviceErr != nil {
		utils.ErrorResponse(w, serviceErr.StatusCode, serviceErr.Message, nil)
		return
	}

	utils.SuccessResponse(w, http.StatusOK, "Customer updated successfully", nil)
}

func (h *CustomerHandler) GetCustomerByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid customer id", nil)
		return
	}

	customer, familyLists, svcErr := h.customerService.GetCustomerDetailsByID(id)
	if svcErr != nil {
		utils.ErrorResponse(w, svcErr.StatusCode, svcErr.Message, nil)
		return
	}

	responseData := map[string]interface{}{
		"customer":    customer,
		"family_list": familyLists,
	}

	utils.SuccessResponse(w, http.StatusOK, "Customer details", responseData)
}

func (h *CustomerHandler) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	idCustomer, err := strconv.Atoi(idStr)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid customer id", nil)
		return
	}

	serviceErr := h.customerService.Delete(idCustomer)
	if serviceErr != nil {
		utils.ErrorResponse(w, serviceErr.StatusCode, serviceErr.Message, nil)
		return
	}

	utils.SuccessResponse(w, http.StatusOK, "Customer deleted successfully", nil)
}

func (h *CustomerHandler) GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	pageQuery := r.URL.Query().Get("page")
	limitQuery := r.URL.Query().Get("limit")

	page, err := strconv.Atoi(pageQuery)
	if err != nil || page < 1 {
		// default page is 1
		page = 1
	}

	limit, err := strconv.Atoi(limitQuery)
	if err != nil || limit < 1 {
		// default limit is 10
		limit = 10
	}

	allCustomers, total, serviceErr := h.customerService.GetAllCustomers(page, limit)
	if serviceErr != nil {
		utils.ErrorResponse(w, serviceErr.StatusCode, serviceErr.Message, nil)
		return
	}

	utils.PaginationResponse(w, http.StatusOK, "Successfully get customers", allCustomers, total, page, limit)
}
