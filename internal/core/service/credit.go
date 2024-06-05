package service

import (
	"ryosantouchh/gym-management-backend/internal/core/entities"
	"ryosantouchh/gym-management-backend/internal/core/models"
	"ryosantouchh/gym-management-backend/internal/core/ports"
)

type CreditService struct {
	repo ports.CreditRepository
}

func NewCreditService(repo ports.CreditRepository) *CreditService {
	return &CreditService{repo: repo}
}

func (s *CreditService) CreateCredit(req models.Credit) error {
	creditType := entities.Credit{
		Type:           req.Type,
		IsMembership:   req.IsMembership,
		ExpireDuration: req.ExpireDuration,
	}

	err := s.repo.Create(&creditType)
	if err != nil {
		return err
	}
	return nil
}
