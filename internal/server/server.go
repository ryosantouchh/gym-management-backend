package server

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	app *gin.Engine
	db  *gorm.DB
}

// func Connect(dsn string) *gorm.DB {
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		panic(err)
// 	}
// 	return db
// }

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

	// mock port
	s.HttpListen(":8080")
}

func (s *Server) HttpListen(port string) {
	s.app.Run(port)
}

