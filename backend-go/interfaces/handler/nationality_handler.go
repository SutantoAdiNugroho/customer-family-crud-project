package handler

import (
	"customer-family-crud-backend/service"
	"customer-family-crud-backend/utils"
	"net/http"
)

type NationalityHandler struct {
	nationalityService service.NationalityService
}

func NewNationalityHandler(nationalityService service.NationalityService) *NationalityHandler {
	return &NationalityHandler{nationalityService: nationalityService}
}

func (h *NationalityHandler) GetAllNationalities(w http.ResponseWriter, r *http.Request) {
	nationalities, svcErr := h.nationalityService.GetAllNationalities()
	if svcErr != nil {
		utils.ErrorResponse(w, svcErr.StatusCode, svcErr.Message, nil)
		return
	}

	utils.SuccessResponse(w, http.StatusOK, "Nationalities fetched successfully", nationalities)
}
