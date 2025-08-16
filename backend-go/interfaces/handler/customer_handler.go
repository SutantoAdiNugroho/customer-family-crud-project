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
