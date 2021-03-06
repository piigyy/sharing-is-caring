package model

type Config struct {
	Production bool   `mapstructure:"production"`
	Port       string `mapstructure:"port"`
	Mongo      Mongo  `mapstructure:"mongo"`
	JWTSecret  string `mapstructure:"jwtSsecret"`
}

type Mongo struct {
	URI  string `mapstructure:"uri"`
	User string `mapstructure:"user"`
	Pass string `mapstructure:"pass"`
	DB   struct {
		Auth string `mapstructure:"auth"`
	} `mapstructure:"db"`
}
