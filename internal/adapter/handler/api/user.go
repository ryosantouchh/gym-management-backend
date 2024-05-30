package api

import (
	"fmt"
	"ryosantouchh/gym-management-backend/internal/core/ports"
	"ryosantouchh/gym-management-backend/internal/core/service"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) GetUserById(ctx ports.HTTPContext) {
	fmt.Println("this is the log")
}
