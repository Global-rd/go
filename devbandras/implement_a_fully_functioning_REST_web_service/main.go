package main

import (
	"bookstore/config"
	"fmt"
	"log/slog"
	"os"
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

	s := fmt.Sprintf("ServerConfig: %+v\nDBServerConfig: %+v", config.Server, config.DBServer)
	fmt.Println(s)

}
