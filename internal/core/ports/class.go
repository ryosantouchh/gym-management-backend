package ports

import "ryosantouchh/gym-management-backend/internal/core/entities"

type ClassRepository interface {
	Create(*entities.Class) error
	Get() (*[]entities.Class, error)
	GetByID(string) (*entities.Class, error)
	Update(string, *entities.Class) error
}
