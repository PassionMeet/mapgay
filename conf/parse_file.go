package conf

import (
	"fmt"

	"github.com/spf13/viper"
)

var _config = Config{}

type Server struct {
	Host string
	Port int
}

type Config struct {
	Wx struct {
	} `yaml:"wx"`
	MySQL struct {
		Server   `yaml:"server"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	} `yaml:"my_sql"`
	Redis struct {
		Server `yaml:"server"`
		Auth   string `yaml:"auth"`
	} `yaml:"redis"`
	NSQ struct {
		Producer []Server `yaml:"producer"`
		Consumer []Server `yaml:"consumer"`
	} `yaml:"nsq"`
}

func ParseJipengConf() {
	viper.SetConfigFile("config.yaml")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&_config)
	if err != nil {
		panic(err)
	}

	fmt.Println(_config)
}
