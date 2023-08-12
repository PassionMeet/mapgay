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

type CosBucket struct {
	BucketName string `yaml:"bucketName"`
	Region     string `yaml:"region"`
	Appid      string `yaml:"appid"`
}

type Cos struct {
	SecretID  string    `yaml:"secretID"`
	SecretKey string    `yaml:"secretKey"`
	Avatar    CosBucket `yaml:"avatar"`
}

type Config struct {
	Wx     Wx     `yaml:"wx"`
	MySQL  MySQL  `yaml:"mysql"`
	Redis  Redis  `yaml:"redis"`
	Cos    Cos    `yaml:"cos"`
	Server Server `yaml:"server"`
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
