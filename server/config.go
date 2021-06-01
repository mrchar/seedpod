package server

import "github.com/spf13/viper"

var defaultConfig = Config{
	Address: "localhost",
	Port:    8080,
}

type Config struct {
	Address string `mapstructure:"address"`
	Port    int    `mapstructure:"port"`
}

func (c *Config) Load() error {
	if err := viper.UnmarshalKey("server", c); err != nil {
		return err
	}
	return nil
}
