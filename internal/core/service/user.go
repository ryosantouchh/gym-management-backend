package service

import "ryosantouchh/gym-management-backend/internal/core/ports"

type UserService struct {
	repo ports.UserRepository
}

func NewUserService(repo ports.UserRepository) *UserService {
	return &UserService{repo: repo}
}
