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
}

func (conf *Config) Parse(path string) error {
	file, err := os.OpenFile(path, os.O_RDONLY, 0)
	if err != nil {
		return fmt.Errorf("failed to open config file: %w", err)
	}

	return yaml.NewDecoder(file).Decode(conf)
}
