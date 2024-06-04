package server

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	app *gin.Engine
	db  *gorm.DB
}

func NewServerApp(db *gorm.DB) *Server {
	ginApp := gin.Default()

	return &Server{
		app: ginApp,
		db:  db,
	}
}

func (s *Server) Start() {
	// register all route here
	s.initUserRouter(s.db)
	s.initClassRouter(s.db)

	// mock port
	s.HttpListen(":8080")
}

func (s *Server) HttpListen(port string) {
	s.app.Run(port)
}
