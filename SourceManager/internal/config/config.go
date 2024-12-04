package config

import (
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env          string     `yaml:"env" env-default:"local"`
	DatabasePass string     `yaml:"database_pass"`
	GRPC         GRPCConfig `yaml:"grpc"`
	HTTP         HTTPConfig `yaml:"http"`
}

type GRPCConfig struct {
	Port    string        `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

type HTTPConfig struct {
	Port    string        `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

func MustLoad() *Config {

	path := "config/prod.yaml"

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("Path does not exist:" + path)
	}

	var cfg Config
	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("Failed to read config:" + err.Error())
	}
	return &cfg
}
