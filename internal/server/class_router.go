package server

import (
	"ryosantouchh/gym-management-backend/internal/adapter/handler/api"
	"ryosantouchh/gym-management-backend/internal/adapter/storage/repository"
	"ryosantouchh/gym-management-backend/internal/core/service"

	"gorm.io/gorm"
)

func (s *Server) initClassRouter(db *gorm.DB) {
	classRouter := s.app.Group("/v1/class")

	classRepository := repository.NewClassReposityImpl(db)
	classService := service.NewClassService(classRepository)
	classHandler := api.NewClassHandler(*classService)

	classRouter.POST("", api.GinHandler(classHandler.CreateClass))
	classRouter.GET("", api.GinHandler(classHandler.GetClasses))
}
