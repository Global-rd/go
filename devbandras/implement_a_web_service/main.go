package main

import (
	"fmt"
	"os"
	"web_service/logger"
	"web_service/service"
)

const (
	defaultBookFileName = "books.json"
	defaultPort         = 8181
)

func main() {

	// logger létrehozása
	appLog := logger.NewAppLogger()

	// bookStore létrehozása
	bookStore := service.NewBookStore()
	err := bookStore.LoadBooksFromFile(defaultBookFileName)
	if err != nil {
		appLog.Logger.Error(fmt.Sprintf("Hiba a book store inicializálása közben: %s", defaultBookFileName), "error", err)
		// ha nincs adatbázis kilépünk
		os.Exit(1)
	}

	// http server létrehozása
	webService := service.NewWebService(defaultPort, bookStore, appLog)

	// http server indítása
	appLog.Logger.Info("HTTP szerver indítása", "port", defaultPort)
	err = webService.Execute()
	if err != nil {
		appLog.Logger.Error("Hiba az HTTP szerver indítása közben", "error", err)
		os.Exit(1)
	}
}
