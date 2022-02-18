package main

import (
	"fmt"

	"github.com/gerardocv1/go-and-compose/config"
)

func main() {
	fmt.Println(config.Get().AppName)
	fmt.Println(config.Get().AppName)
	fmt.Println(config.Get().Port)

	fmt.Scanln()
}
