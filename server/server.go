package server

import (
	"github.com/gerardocv1/go-and-compose/server/di"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

type Server struct {
	Router    *gin.Engine
	Container *dig.Container
}

// NewServer instance
func NewServer(e *gin.Engine, c *dig.Container) *Server {
	return &Server{
		Router:    e,
		Container: c,
	}
}

func Run(call func(*Server)) error {

	g := gin.Default()
	d := di.BuildContainer()

	svr := NewServer(g, d)

	call(svr)

	return nil
}
