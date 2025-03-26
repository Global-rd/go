package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

type Config struct {
	Db     DbCfg     `yaml:"db"`
	Server ServerCfg `yaml:"server"`
}

type DbCfg struct {
	DbName string `yaml:"dbname"`
	DbUser string `yaml:"user"`
	DbPass string `yaml:"pass"`
}

type ServerCfg struct {
	Address string `yaml:"address"`
	Port    int    `yaml:"port"`
}

func getExecutionDir() (string, error) {
	execPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return "", err
	}
	return execPath, nil
}

func LoadConfig() (*Config, error) {
	execDir, err := getExecutionDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get CWD: %w", err)
	}

	v := viper.New()
	v.SetConfigName("movie-ws") // name of config file (without extension)
	v.SetConfigType("yaml")     // YAML format
	v.AddConfigPath(execDir)    // path to look for the config file in

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config
	if err := v.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &config, nil

}
