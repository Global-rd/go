package config

import (
	"github.com/spf13/viper"
)

type Cfg struct {
	Kafkaurl string `mapstructure:"kafkaurl"`
	Logtopic string `mapstructure:"logtopic"`
	Groupid  string `mapstructure:"groupid"`
}

func SetConfig() (*Cfg, error) {
	// Set the file name of the configurations file
	viper.SetConfigName("config")

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")
	var configuration Cfg

	if err := viper.ReadInConfig(); err != nil {
		return &configuration, err
	}

	// Set undefined variables
	viper.SetDefault("database.dbname", "books")

	err := viper.Unmarshal(&configuration)
	if err != nil {
		return &configuration, err
	}

	return &configuration, nil

}
