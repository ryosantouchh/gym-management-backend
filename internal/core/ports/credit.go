package ports

import "ryosantouchh/gym-management-backend/internal/core/entities"

type CreditRepository interface {
	Create(*entities.Credit) error
	// Get() (*[]entities.Credit, error)
}
