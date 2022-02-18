package config

import (
	"os"
	"path/filepath"
	"runtime"
	"sync"

	"github.com/jinzhu/configor"
)

type Config struct {
	AppName string `yaml:"appname"`
	Port    int32  `yaml:"port"`
}

var once sync.Once
var instance *Config

func Get() *Config {
	once.Do(func() {

		basepath := os.Getenv("PROJECT_PATH")
		if basepath == "" {
			_, file, _, _ := runtime.Caller(0)
			basepath = filepath.Dir(file) + "/../"
		}
		instance = &Config{}

		//this gets the path to the folder of the file and then move one level to find the .env.yaml file
		configor.Load(instance, basepath+".env.yaml")
	})
	return instance
}
