package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	DBConfig `yaml:"postgres"`
}

type DBConfig struct {
	Host         string `yaml:"host"`
	Port         int    `yaml:"port"`
	User         string `yaml:"user" env:"DB_USER" env-default:"postgres"`
	Password     string `yaml:"password" env:"DB_PASS" env-default:"postgres"`
	DatabaseName string `yaml:"db" env:"DB_NAME" env-default:"postgres"`
	TimeoutSec   int    `yaml:"db-timeout-sec"`
}

func ReadConfig() (*Config, error) {
	cfg := Config{}
	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("failed to get working directory: %v", err)
	}

	configPath := filepath.Join(dir, "config", "config.yaml")

	err = cleanenv.ReadConfig(configPath, &cfg)
	if err != nil && !os.IsNotExist(err) {
		return nil, fmt.Errorf("failed to read config file %s: %v", configPath, err)
	}
	err = cleanenv.ReadEnv(&cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to read env file: %v", err)
	}

	return &cfg, nil
}

