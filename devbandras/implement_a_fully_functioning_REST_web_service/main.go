package main

import (
	"bookstore/config"
	"bookstore/service"
	"database/sql"
	"log/slog"
	"os"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/lib/pq"
)

func main() {
	// default logger beállítása
	appLog := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	// config beolvasása
	config, err := config.LoadConfig("config.yml")
	if err != nil {
		appLog.Error("Hiba a konfiguráció betöltése közben", "error", err)
		// ha nics konfiguráció kilépünk a programból
		os.Exit(1)
	}

	// kapcsolódás az adatbázishoz
	appLog.Info("Kapcsolódás az adatbázishoz", "host", config.DBServer.Host, "port", config.DBServer.Port, "dbname", config.DBServer.DBName)
	connectionString := config.DBServer.GetConnectionString()
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		appLog.Error("Hiba az adatbázis kapcsolódása közben", "error", err)
		// ha nics kapcsolódás kilépünk a programból
		os.Exit(1)
	}
	defer db.Close()

	// adatbázis kapcsolat ellenőrzése
	err = db.Ping()
	if err != nil {
		appLog.Error("Hiba az adatbázis kapcsolat ellenőrzése közben", "error", err)
		// ha nics kapcsolat ellenőrzés kilépünk a programból
		os.Exit(1)
	}
	appLog.Info("Kapcsolódás az adatbázishoz sikeres")

	// sql builder létrehozása az adatbázisból
	sqlBuilder := goqu.New("postgres", db)

	// webservice indítása
	appLog.Info("Webservice indítása", "host", config.Server.Host, "port", config.Server.Port)
	websService := service.NewWebService(&config.Server, sqlBuilder, appLog)
	err = websService.Execute()
	if err != nil {
		appLog.Error("Hiba a webservice indítása közben", "error", err)
		// ha sikertelen a webservice indítása akkor kilépünk a programból
		os.Exit(1)
	}
}
