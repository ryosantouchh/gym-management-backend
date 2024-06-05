package repository

import (
	"fmt"
	"ryosantouchh/gym-management-backend/internal/core/entities"

	"gorm.io/gorm"
)

type CreditReposityImpl struct {
	db *gorm.DB
}

func NewCreditReposityImpl(db *gorm.DB) *CreditReposityImpl {
	return &CreditReposityImpl{db: db}
}

func (r *CreditReposityImpl) Create(creditType *entities.Credit) error {
	result := r.db.Create(creditType)

	if result.Error != nil {
		err := fmt.Errorf("Failed to create credit type : %w\n", result.Error)
		return err
	}
	return nil
}

// func (r *CreditTypeReposityImpl) Get() (*[]entities.CreditType, error) {
// }
