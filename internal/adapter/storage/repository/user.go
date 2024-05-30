package repository

import (
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
	// db interface{}
}

func NewUserRepositoryImpl(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) Get() {

}
