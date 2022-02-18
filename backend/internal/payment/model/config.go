package model

type Config struct {
	Port     string   `mapstructure:"port"`
	Payment  Payment  `mapstructure:"payment"`
	Database Database `mapstructure:"database"`
	Certfile string   `mapstructure:"certfile"`
	Keyfile  string   `mapstructure:"keyfile"`
}

type Payment struct {
	Method string `mapstructure:"method"`
	URL    string `mapstructure:"url"`
	Key    string `mapstructure:"key"`
}

type Database struct {
	URI string `mapstructure:"uri"`
	DB  string `mapstructure:"db"`
}
