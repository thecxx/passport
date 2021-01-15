package restful

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/thecxx/passport/pkg/passport/domain/restful/handlers/v1"
)

type Server struct {
	engine *gin.Engine
}

// New a server.
func NewServer() *Server {
	s := &Server{
		engine: gin.New(),
	}
	s.init()
	return s
}

func (s *Server) init() {
	v1g := s.engine.Group("/v1")
	{
		v1g.POST("/passport/login", v1.LoginHandler)
		v1g.POST("/passport/register", v1.RegisterHandler)
	}
}
