package config

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/ilyakaznacheev/cleanenv"
)


type (
	Config struct {
		App  `yaml:"app"`
		HTTP `yaml:"http"`
		MySQL   `yaml:"mysql"`
	}

	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}

	HTTP struct {
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}

	MySQL struct {
		URL   string `env:"MYSQL_URL"`
	}
)

func init() {
    // loads values from .env into the system
    if err := godotenv.Load(); err != nil {
        log.Print("No .env file found")
    }
}


func NewConfig() (*Config, error) {
	config := &Config{}

	err := cleanenv.ReadConfig("./config/config.yml", config)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(config)
	if err != nil {
		return nil, err
	}


	return config, nil
}