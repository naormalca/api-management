package config
import (
	"github.com/jinzhu/configor"
)
type JWT struct {
	Secret string `required:"true"`
}

type Database struct {
	Dialect  string `default:"postgres"`
	Debug    bool   `default:"false"`
	Username string `required:"true"`
	Password string `required:"true"`
	Host     string `required:"true"`
	Port     int
	SSLMode bool
	Source	string
}

type Configuration struct {
	Database Database `required:"true"`
	JWT      JWT      `required:"true"`
}
var Main = (func() Configuration {
	//TODO: make the config path as a env variable
	var conf Configuration
	if err := configor.Load(&conf, "/home/nmalca/prvDev/Go/ApiManagement/config/config-local-container.ENV.json"); err != nil {
		panic(err.Error())
	}
	return conf
})()