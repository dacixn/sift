package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/goccy/go-yaml"
)

type Config struct {
	Groups            map[string][]string `yaml:"groups"`
	EnableGroups      bool                `yaml:"enableGroups"`
	DirPrefix         string              `yaml:"dirPrefix"`
	DefaultWorkingDir string              `yaml:"workingDir"`
}

// put yaml in here
// var configFile []byte

func InitConfig() *Config {
	// locateConfig
	// if not exist plantConfig
	// then readConfig
	// wtf should this return??
	// take no args return config ptr
	path, err := locateConfigFile()
	if err != nil {
		path, err = plantConfigFile()
	}

	cfg, err := readConfig(path)
	return cfg
}

// functions local to this module

func readConfig(path string) (*Config, error) {
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

func locateConfigFile() (string, error) {
	// search user config dir for ./sift/config.yaml
	// otherwise search working dir for config.yaml
	// WTF AM I DOING
	// hardcoded for now, change with flag eventually (probably with app struct)

	var configDir string
	var configPath string
	var configName string = "config.yaml"

	// might be unidiomatic but what if i flip the conditions
	// if err == nil, else is easier to read

	configDir, err := os.UserConfigDir()
	if err != nil {
		configDir, err = os.Getwd()
		if err != nil {
			return "", fmt.Errorf("Error opening directories: %w", err)
		} else {
			configPath = filepath.Join(configDir, configName)
			_, err = os.Stat(configPath)
		}
	} else {
		configDir = filepath.Join(configDir, "sift") // only add sift subfolder if in .config/
		configPath = filepath.Join(configDir, configName)
		_, err := os.Stat(configPath)
		if err != nil {
			return "", fmt.Errorf("Error locating configuration file: %w", err)
		}
	}
	_, err = os.Stat(configPath)
	if err != nil {
		return "", fmt.Errorf("Error locating configuration file: %w", err)
	}

	return configPath, nil
}

func plantConfigFile() (path string, err error) {
	return path, nil
}
