package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"ryosantouchh/gym-management-backend/internal/core/models"
	"ryosantouchh/gym-management-backend/internal/core/ports"
	"ryosantouchh/gym-management-backend/internal/core/service"

	"gorm.io/gorm"
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

  defer req.Body.Close()

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&classRequest)
	if err != nil {
		res := models.ResponseError{Error: "Invalid request body!"}
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	err = h.service.CreateClass(classRequest)
	if err != nil {
		res := models.ResponseError{Error: "Failed to create class!"}
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := models.ResponseMessage{
		Message: "Class created successfully!",
	}
	ctx.JSON(http.StatusCreated, res)
}

func (h *ClassHandler) GetClassList(ctx ports.HTTPContext) {
	classes, err := h.service.GetClassList()
	if err != nil {
		res := models.ResponseError{Error: "Failed to get classes!"}
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := models.ResponseSuccess{}
	res.Message = "Get classes successfully!"
	res.Data = *classes
	ctx.JSON(http.StatusOK, res)
}

func (h *ClassHandler) GetClassByID(ctx ports.HTTPContext) {
	id := ctx.Param("id")

  class, err := h.service.GetClassByID(id)
  if err != nil {
    var res models.ResponseError
    if errors.Is(err, gorm.ErrRecordNotFound) {
      res.Error = "Not found class!"
      ctx.JSON(http.StatusNotFound, res)
    } else {
      res.Error = "Failed to get class by ID!"
      ctx.JSON(http.StatusInternalServerError, res)
    }
		return
  }

	res := models.ResponseSuccess{}
	res.Message = "Get class by ID successfully!"
	res.Data = *class
	ctx.JSON(http.StatusOK, res)
}

func (h *ClassHandler) UpdateClassByID(ctx ports.HTTPContext) {
  id := ctx.Param("id")
	req := ctx.Request()

  defer req.Body.Close()

	var updateRequest models.UpdateClassRequest
  decoder := json.NewDecoder(req.Body)
  err := decoder.Decode(&updateRequest)

	if err != nil {
		res := models.ResponseError{Error: "Invalid request body!"}
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

  err = h.service.UpdateClass(id, &updateRequest)
}
