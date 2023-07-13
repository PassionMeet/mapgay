package conf

import (
	"github.com/spf13/viper"
)

var _config = Config{}

func Get() *Config {
	return &_config
}

type Server struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type MySQL struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DB       string `yaml:"db"`
}

type Wx struct {
	AppID     string `yaml:"AppID"`
	AppSecret string `yaml:"AppSecret"`
}

type Redis struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	Auth string `yaml:"auth"`
}

type Config struct {
	Wx    *Wx    `yaml:"wx"`
	MySQL *MySQL `yaml:"mysql"`
	Redis *Redis `yaml:"redis"`
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

}
