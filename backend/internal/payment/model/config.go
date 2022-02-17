package model

type Config struct {
	Port     string   `mapstructure:"port"`
	Payment  Payment  `mapstructure:"payment"`
	Database Database `mapstructure:"database"`
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
