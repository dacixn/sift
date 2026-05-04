package config

import (
	_ "embed"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

const (
	DefaultConfigName = ".sift.toml"
)

type Config struct {
	DirPrefix  string              `toml:"dirPrefix"`
	Groups     map[string][]string `toml:"groups"`
	IgnoreDirs bool                `toml:"ignoreDirs"`
}

//go:embed .sift.toml
var configFile []byte

func (c *Config) Init() error {
	path, err := locateConfigFile()
	if err != nil {
		path, err = plantConfigFile()
		if err != nil {
			return fmt.Errorf("failed to plant config: %w", err)
		}
	}

	cfg, err := readConfig(path)
	if err != nil {
		return fmt.Errorf("failed to read config: %w", err)
	}
	*c = *cfg
	return nil
}

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

func appConfigDir() (string, error) {
	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(userConfigDir, "sift"), nil
}

func locateConfigFile() (string, error) {
	var errs []error
	var dirs []string

	wd, err := os.Getwd()
	if err != nil {
		errs = append(errs, err)
	} else {
		dirs = append(dirs, wd)
	}

	appConfigDir, err := appConfigDir()
	if err != nil {
		errs = append(errs, err)
	} else {
		dirs = append(dirs, appConfigDir)
	}

	for _, dir := range dirs {
		path := filepath.Join(dir, DefaultConfigName)
		_, err = os.Stat(path)
		if err == nil {
			return path, nil
		} else {
			errs = append(errs, err)
		}
	}

	return "", fmt.Errorf("error locating configuration file: %w", errors.Join(errs...))
}

func plantConfigFile() (string, error) {
	appConfigDir, err := appConfigDir()
	if err != nil {
		return "", err
	}
	err = os.MkdirAll(appConfigDir, 0755)
	if err != nil {
		return "", err
	}
	path := filepath.Join(appConfigDir, DefaultConfigName)
	err = os.WriteFile(path, configFile, 0600)
	if err != nil {
		return "", err
	}

	return path, nil
}
