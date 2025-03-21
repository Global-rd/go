package config

import (
	"fmt"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/spf13/viper"
)

type Cfg struct {
	Server Server
	DB     DB
}

type Server struct {
	Address string
	Port    int
}

type DB struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string `mapstructure:"db_name"`
}

func (db DB) ConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		db.Host, db.Port, db.User, db.Password, db.DBName)
}

func (d DB) Validate() error {
	return validation.ValidateStruct(&d,
		validation.Field(&d.Host, validation.Required),
		validation.Field(&d.Port, validation.Required),
		validation.Field(&d.User, validation.Required),
		validation.Field(&d.Password, validation.Required),
		validation.Field(&d.DBName, validation.Required),
	)
}

func (s Server) Validate() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.Address, validation.Required),
		validation.Field(&s.Port, validation.Required),
	)
}

func (c Cfg) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Server, validation.Required),
		validation.Field(&c.DB, validation.Required),
	)
}

func Parse() (result *Cfg, err error) {
	viper.SetConfigName("config")
	viper.AddConfigPath("configs")
	viper.SetEnvPrefix("WEBSERVICE")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetConfigType("yml")

	if err = viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file, %s", err)
	}

	err = viper.Unmarshal(&result)
	if err != nil {
		return nil, fmt.Errorf("unable to decode into worker, %v", err)
	}

	err = result.Validate()
	if err != nil {
		return nil, fmt.Errorf("invalid configuration provided %v", err)
	}

	return
}
