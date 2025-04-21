package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Cfg struct {
	Server Server
	DB     DB
}

type Server struct {
	Port         int
	ReadTimeout  int `mapstructure:"read_timeout"`
	WriteWimeout int `mapstructure:"write_timeout"`
	PORT         string
	READTIMEOUT  string
	WRITETIMEOUT string
}

type DB struct {
	Dialect          string
	Host             string
	Port             int
	User             string
	Password         string
	DBName           string `mapstructure:"db_name"`
	ConnectionString string `mapstructure:"connection_string"`
}

func BuildConnectionString(db DB) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		db.Host, db.Port, db.User, db.Password, db.DBName)
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
