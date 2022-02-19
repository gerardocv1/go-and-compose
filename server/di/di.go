package di

import (
	"github.com/gerardocv1/go-and-compose/config"
	"go.uber.org/dig"
)

var container = dig.New()

//BuildContainer provide all injectable services
func BuildContainer() *dig.Container {

	// config
	container.Provide(config.Get)

	// DBs
	// container.Provide(db.NewDb)
	// container.Provide(awsclient.NewS3ClientImplementation)

	return container
}

//@Invoke(i interface{})
//@return error
func Invoke(i interface{}) error {
	return container.Invoke(i)
}
