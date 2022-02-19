package router

import (
	"net/http"

	"github.com/gerardocv1/go-and-compose/config"
	"github.com/gerardocv1/go-and-compose/server"
	"github.com/gin-gonic/gin"
)

//MapRoutes function
func MapRoutes(svr *server.Server) {

	svr.Router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, config.Get().AppName)
		c.Status(200)
	})
}
