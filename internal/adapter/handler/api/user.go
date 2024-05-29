package api

import "ryosantouchh/gym-management-backend/internal/core/service"

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{service: service}
}
