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
		err := fmt.Errorf("Failed to create class : %w\n", result.Error)
		return err
	}
	return nil
}

func (r *ClassReposityImpl) Get() (*[]entities.Class, error) {
	var classes []entities.Class
	result := r.db.Find(&classes)

	if result.Error != nil {
		err := fmt.Errorf("Cannot get classes : %w\n", result.Error)
		return nil, err
	}
	return &classes, nil
}

func (r *ClassReposityImpl) GetByID(id string) (*entities.Class, error) {
	var class entities.Class
	result := r.db.Where("id = ?", id).Find(&class)

	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	if result.Error != nil {
		err := fmt.Errorf("Cannot get class : %w\n", result.Error)
		return nil, err
	}
	return &class, nil
}

func (r *ClassReposityImpl) Update(id string, data interface{}) error {
	result := r.db.Model(entities.Class{}).Where("id = ?", id).Updates(data)

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	if result.Error != nil {
		err := fmt.Errorf("Cannot update class : %w\n", result.Error)
		return err
	}
	return nil
}
