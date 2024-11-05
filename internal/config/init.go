package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Addr       string `yaml:"addr"`        // * 0.0.0.0:8086 by default
	Spatref    int    `yaml:"spatref"`     // * 4326 by default
	WorkerType string `yaml:"worker_type"` // * ord by default, also have mock, inventions in gdal/structs
	Logdest    string `yaml:"logdest"`     // * logs/service_poly_intersection.log by default
	Loglevel   string `yaml:"loglevel"`    // * info by default
	Verbose    string `yaml:"verbose"`     // * true by default, provides to see swagger docs
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
