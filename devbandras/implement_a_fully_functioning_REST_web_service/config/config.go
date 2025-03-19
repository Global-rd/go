package config

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/spf13/viper"
)

type ServerConfig struct {
	Address string
	Port    int
}

type DB struct {
	Host         string
	Port         int
	DatabaseName string `mapstructure:"database_name"`
	User         string
	Password     string
}

type Config struct {
	Server   ServerConfig
	DBServer DB
}

// A függvény validálja a ServerConfig struktúrát.
//
// Returns:
//   - error: nil ha az érvénysítés sikeres, egyébként a hibaüzenet
func (sc *ServerConfig) Validate() error {
	return validation.ValidateStruct(
		sc,
		validation.Field(&sc.Address, validation.Required),
		validation.Field(&sc.Port, validation.Required),
	)
}

// A függvény visszadja a connection string-et a DBServerConfig struktúrának alapján.
// A connection string formátuma:
// "host=<Host> port=<Port> user=<User> password=<Password> dbname=<DatabaseName> sslmode=disable"
//
// Returns:
//   - string: PostgreSQL kapcsolat formátumú connection stringet ad vissza.
func (db *DB) GetConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", db.Host, db.Port, db.User, db.Password, db.DatabaseName)
}

// A függvény validálja a DBServerConfig struktúrát.
// Ellenőrzi, hogy minden kötelező mező meg legyen adva, illetve a port 1-65535 között legyen.
//
// Returns:
//   - error: nil ha az érvénysítés sikeres, egyébként a hibaüzenet
func (db *DB) Validate() error {
	return validation.ValidateStruct(
		db,
		validation.Field(&db.Host, validation.Required),
		validation.Field(&db.Port, validation.Required, validation.Min(1), validation.Max(65535)),
		validation.Field(&db.DatabaseName, validation.Required),
		validation.Field(&db.User, validation.Required),
		validation.Field(&db.Password, validation.Required),
	)
}

// A függvény validálja a Config struktúrát.
//
// Returns:
//   - error: nil ha az érvénysítés sikeres, egyébként a hibaüzenet
func (c *Config) Validate() error {
	return validation.ValidateStruct(
		c,
		validation.Field(&c.Server, validation.Required),
		validation.Field(&c.DBServer, validation.Required),
	)
}

// A LoadConfig beolvassa és elemzi a konfigurációs fájlt, és visszaad egy konfigurációs struktúrát.
//
// Parameters:
//   - configFile: Konfigurációs fájl neve
//
// Returns:
//   - *Config: A beolvasott, validált, és struktúrált konfigurációs struktúra
//   - error: nil ha a config fájl beolvasása és validálása sikeres, egyébként a hibaüzenet
func LoadConfig(configFile string) (*Config, error) {
	if configFile == "" {
		return nil, fmt.Errorf("konfigurációs fájl megadása kötelező")
	}

	viper.AddConfigPath(".")
	viper.SetConfigName(configFile)
	viper.SetConfigType("yml")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("konfigurációs fájl (%s) beolvasása sikertelen :%w", configFile, err)
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, fmt.Errorf("konfiguráció beolvasása sikertelen :%w", err)
	}

	err = config.Validate()
	if err != nil {
		return nil, fmt.Errorf("konfiguráció validálás hiba: %w", err)
	}

	return &config, nil
}
