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

func ConvertYamlToStruct(path string) (*Config, error) {
	var config Config
	yamlBytes, err := os.ReadFile(path)
	if err != nil {
		return &config, err
	}
	err = yaml.Unmarshal(yamlBytes, &config)
	if err != nil {
		return &config, err
	}
	return &config, nil
}
