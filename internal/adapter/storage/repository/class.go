package repository

import (
	"fmt"
	"ryosantouchh/gym-management-backend/internal/core/entities"

	"gorm.io/gorm"
)

type ClassReposityImpl struct {
	db *gorm.DB
}

func NewClassReposityImpl(db *gorm.DB) *ClassReposityImpl {
	return &ClassReposityImpl{db: db}
}

func (r *ClassReposityImpl) Create(class *entities.Class) error {
	result := r.db.Create(class)

	if result.Error != nil {
		err := fmt.Errorf("failed to create class : %w\n", result.Error)
		return err
	}
	return nil
}

func (r *ClassReposityImpl) Get() (*[]entities.Class, error) {
	var classes []entities.Class
	result := r.db.Find(&classes)

	if result.Error != nil {
		err := fmt.Errorf("failed to create class : %w\n", result.Error)
		return nil, err
	}
	return &classes, nil
}
