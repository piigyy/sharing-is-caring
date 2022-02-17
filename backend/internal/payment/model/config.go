package model

type Config struct {
	Port    string  `mapstructure:"port"`
	Payment Payment `mapstructure:"payment"`
}

type Payment struct {
	Method string `mapstructure:"method"`
	URL    string `mapstructure:"url"`
	Key    string `mapstructure:"key"`
}
