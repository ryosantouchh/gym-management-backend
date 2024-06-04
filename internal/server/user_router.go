package server

import (
	"ryosantouchh/gym-management-backend/internal/adapter/handler/api"
	"ryosantouchh/gym-management-backend/internal/adapter/storage/repository"
	"ryosantouchh/gym-management-backend/internal/core/service"

	"gorm.io/gorm"
)

func (s *Server) initUserRouter(db *gorm.DB) {
	userRouter := s.app.Group("/v1/user")

	userRepository := repository.NewUserRepositoryImpl(db)
	userService := service.NewUserService(userRepository)
	userHanlder := api.NewUserHandler(*userService)

	userRouter.GET("", api.GinHandler(userHanlder.GetUserById))
}
