package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Groups            map[string][]string `toml:"groups"`
	EnableGroups      bool                `toml:"enableGroups"`
	DirPrefix         string              `toml:"dirPrefix"`
	DefaultWorkingDir string              `toml:"workingDir"`
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
	tomlBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = toml.Unmarshal(tomlBytes, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

func locateConfigFile() (string, error) {
	// hardcoded for now, change with flag eventually (probably with app struct)

	configName := "config.toml"

	configDir, err := os.UserConfigDir()
	if err == nil {
		// only check for sift folder in userconfigdir
		configDir := filepath.Join(configDir, "sift")
		configPath := filepath.Join(configDir, configName)
		_, err = os.Stat(configPath)
		if err == nil {
			return configPath, nil
		}
	}

	configDir, err = os.Getwd()
	if err == nil {
		configPath := filepath.Join(configDir, configName)
		_, err = os.Stat(configPath)
		if err == nil {
			return configPath, nil
		}
	}

	return "", fmt.Errorf("Error locating configuration file")
}

func plantConfigFile() (path string, err error) {
	return path, nil
}
