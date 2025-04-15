package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

type ServerCfg struct {
	Address string `yaml:"address"`
	Port    int    `yaml:"port"`
}

func getExecutionDir() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	var execPath string
	execPath, err = filepath.Abs(wd)
	if err != nil {
		return "", err
	}
	return execPath, nil
}

func LoadConfig() (*ServerCfg, error) {
	execDir, err := getExecutionDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get CWD: %w", err)
	}

	v := viper.New()
	v.SetConfigName("patterns") // name of config file (without extension)
	v.SetConfigType("yaml")     // YAML format
	v.AddConfigPath(execDir)    // path to look for the config file in

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config ServerCfg
	if err := v.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &config, nil

}
