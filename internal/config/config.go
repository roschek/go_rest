package config

import (
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/roschek/go_rest.git/pkg/logging"
)

type Config struct {
	IsDebug *bool `yaml:"is_debug"`
	Listen  struct {
		Type   string `yaml:"type"`
		BindIp string `yaml:"bind_ip"`
		Port   string `yaml:"port"`
	} `yaml:"listen"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		logger := logging.GetLogger()
		logger.Info("Read app config")
		instance = &Config{}
		err := cleanenv.ReadConfig("pkg/config.yml", instance)
		if err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logger.Info(help)
			logger.Fatal(err)
		}
	})
	return instance
}
