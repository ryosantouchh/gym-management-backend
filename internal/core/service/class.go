package service

import (
	"ryosantouchh/gym-management-backend/internal/core/entities"
	"ryosantouchh/gym-management-backend/internal/core/models"
	"ryosantouchh/gym-management-backend/internal/core/ports"
)

type ClassService struct {
	repo ports.ClassRepository
}

func NewClassService(repo ports.ClassRepository) *ClassService {
	return &ClassService{repo: repo}
}

func (s *ClassService) CreateClass(req models.Class) error {
	class := entities.Class{
		ClassType: req.ClassType,
		Duration:  req.Duration,
	}

	err := s.repo.Create(&class)
	if err != nil {
		return err
	}
	return nil
}

func (s *ClassService) GetClasses() (*[]models.Class, error) {
	classes, err := s.repo.Get()
	if err != nil {
		return nil, err
	}

	classEntity := *classes
	classData := make([]models.Class, 0)
	for i := range *classes {
		data := classEntity[i]
		jsonData := models.Class{
			ID:        data.ID,
			ClassType: data.ClassType,
			Duration:  data.Duration,
		}
		classData = append(classData, jsonData)
	}
	return &classData, nil
}
