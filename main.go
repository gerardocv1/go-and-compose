package main

import (
	"fmt"

	"github.com/gerardocv1/go-and-compose/config"
	"github.com/gerardocv1/go-and-compose/router"
	"github.com/gerardocv1/go-and-compose/server"
)

func main() {
	fmt.Println(config.Get().AppName)

	server.Run(func(svr *server.Server) {
		router.MapRoutes(svr)
		svr.Router.Run(fmt.Sprintf(":%v", config.Get().Port))
	})

	fmt.Scanln()
}
