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
	Host   string
	Port   int
	User   string
	Pass   string
	DBName string `mapstructure:"db_name"`
}

func (db DB) ConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		db.Host, db.Port, db.User, db.Pass, db.DBName)
}

func (db DB) Validate() error {
	return validation.ValidateStruct(&db,
		validation.Field(&db.DBName, validation.Required),
		validation.Field(&db.User, validation.Required),
		validation.Field(&db.Pass, validation.Required),
		validation.Field(&db.Port, validation.Required),
	)
}

func (cfg Cfg) Validate() error {
	return validation.ValidateStruct(&cfg,
		validation.Field(&cfg.Server, validation.Required),
		validation.Field(&cfg.DB, validation.Required),
	)
}

func (s Server) Validate() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.Address, validation.Required),
		validation.Field(&s.Port, validation.Required),
	)
}

func ReadConfig() (result *Cfg, err error) {
	viper.SetConfigName("config")
	viper.AddConfigPath("configs")
	viper.SetEnvPrefix("WEBSERVICE")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetConfigType("yml")

	if err = viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file, %s", err)
	}
	//fmt.Println("Loaded configuration:", viper.AllSettings())

	err = viper.Unmarshal(&result)
	if err != nil {
		return nil, fmt.Errorf("unable to decode into worker, %v", err)
	}
	//fmt.Println("Unmarshalled struct:", result)

	err = result.Validate()
	if err != nil {
		return nil, fmt.Errorf("invalid configuration provided %v", err)
	}

	return
}
