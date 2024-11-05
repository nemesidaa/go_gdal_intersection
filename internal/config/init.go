package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Addr       string `yaml:"addr"`
	Spatref    int    `yaml:"spatref"`
	WorkerType string `yaml:"worker_type"`
	Logdest    string `yaml:"logdest"`
	Loglevel   string `yaml:"loglevel"`
	Verbose    string `yaml:"verbose"`
}

func (conf *Config) Parse(path string) error {
	file, err := os.OpenFile(path, os.O_RDONLY, 0)
	if err != nil {
		return fmt.Errorf("failed to open config file: %w", err)
	}

	if err := yaml.NewDecoder(file).Decode(conf); err != nil {
		return fmt.Errorf("failed to decode config file: %w", err)
	}

	return nil
}
