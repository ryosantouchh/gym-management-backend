package api

import (
	"encoding/json"
	"net/http"
	"ryosantouchh/gym-management-backend/internal/core/models"
	"ryosantouchh/gym-management-backend/internal/core/ports"
	"ryosantouchh/gym-management-backend/internal/core/service"
)

type ClassHandler struct {
	service service.ClassService
}

func NewClassHandler(service service.ClassService) *ClassHandler {
	return &ClassHandler{service: service}
}

func (h *ClassHandler) CreateClass(ctx ports.HTTPContext) {
	var classRequest models.Class
	req := ctx.Request()

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&classRequest)
	if err != nil {
		res := models.ResponseError{Error: "Invalid request body"}
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	err = h.service.CreateClass(classRequest)
	if err != nil {
		res := models.ResponseError{Error: "failed to create class"}
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := models.ResponseMessage{
		Message: "Class created successfully!",
	}
	ctx.JSON(http.StatusCreated, res)
}

func (h *ClassHandler) GetClasses(ctx ports.HTTPContext) {

	classes, err := h.service.GetClasses()
	if err != nil {
		res := models.ResponseError{Error: "failed to get class"}
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := models.ResponseSuccess{}
	res.Message = "Get classes successfully!"
	res.Data = *classes
	ctx.JSON(http.StatusOK, res)
}
