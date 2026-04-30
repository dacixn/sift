package main

import (
	"os"

	"github.com/goccy/go-yaml"
)

type Config struct {
	Groups            map[string][]string `yaml:"groups"`
	EnableGroups      bool                `yaml:"enableGroups"`
	DirPrefix         string              `yaml:"dirPrefix"`
	DefaultWorkingDir string              `yaml:"workingDir"`
}

func ReadConfig(path string) (*Config, error) {
	var cfg Config
	yamlBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(yamlBytes, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
