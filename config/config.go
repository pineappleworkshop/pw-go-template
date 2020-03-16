package config

import (
	"os"
)

var Conf *Config

func InitConf() {
	Conf = new(Config)
	Conf.SetEnv()
}

type Config struct {
	Env string
}

func (c *Config) SetEnv() {
	env := os.Getenv("ENV")
	if env == DEV {
		c.Env = DEV
	} else if env == STAGE {
		c.Env = STAGE
	} else if env == PROD {
		c.Env = PROD
	} else {
		c.Env = WORKSTATION
	}
}

func (c *Config) GetEnv() string {
	return c.Env
}
