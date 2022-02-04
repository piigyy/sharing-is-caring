package config

type Config struct {
	Production bool  `mapstructure:"production"`
	Port       Port  `mapstructure:"port"`
	Mongo      Mongo `mapstructure:"mongo"`
}

type Port struct {
	Auth string `mapstructure:"auth"`
}

type Mongo struct {
	URI  string `mapstructure:"uri"`
	User string `mapstructure:"user"`
	Pass string `mapstructure:"pass"`
	DB   struct {
		Auth string `mapstructure:"auth"`
	} `mapstructure:"db"`
}
