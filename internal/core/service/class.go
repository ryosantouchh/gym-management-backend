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

func (s *ClassService) GetClassList() (*[]models.Class, error) {
	classes, err := s.repo.Get()
	if err != nil {
		return nil, err
	}

	classEntity := *classes
	classData := make([]models.Class, 0)
	for i := range *classes {
		data := classEntity[i]
		createData := models.Class{
			ID:        data.ID,
			ClassType: data.ClassType,
			Duration:  data.Duration,
		}
		classData = append(classData, createData)
	}
	return &classData, nil
}

func (s *ClassService) GetClassByID(id string) (*models.Class, error) {
	class, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	classEntity := *class
	classData := models.Class{
		ID:        classEntity.ID,
		ClassType: classEntity.ClassType,
		Duration:  classEntity.Duration,
	}
	return &classData, nil
}

func (s *ClassService) UpdateClass(id string, updateClass *models.UpdateClassRequest) error {
	updateData := make(map[string]interface{})
	classData := *updateClass

	if classData.ClassType != "" {
		updateData["ClassType"] = classData.ClassType
	}
	if classData.Duration != 0 {
		updateData["Duration"] = classData.Duration
	}

	err := s.repo.Update(id, classData)
	if err != nil {
		return err
	}
	return nil
}
