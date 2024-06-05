package api

import (
	"encoding/json"
	"net/http"
	"ryosantouchh/gym-management-backend/internal/core/models"
	"ryosantouchh/gym-management-backend/internal/core/ports"
	"ryosantouchh/gym-management-backend/internal/core/service"
)

type CreditHandler struct {
	service service.CreditService
}

func NewCreditTypeHandler(service service.CreditService) *CreditHandler {
	return &CreditHandler{service: service}
}

func (h *CreditHandler) CreateCredit(ctx ports.HTTPContext) {
	var creditRequest models.Credit
	req := ctx.Request()

	defer req.Body.Close()

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&creditRequest)
	if err != nil {
		res := models.ResponseError{Error: "Invalid request body!"}
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	err = h.service.CreateCredit(creditRequest)
	if err != nil {
		res := models.ResponseError{Error: "Failed to create credit!"}
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := models.ResponseMessage{
		Message: "Class created credit!",
	}
	ctx.JSON(http.StatusCreated, res)
}
