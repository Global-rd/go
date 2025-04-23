package routes

import (
	"advrest/logger"
	"database/sql"
	"net/http"
)

func Helloroute(db *sql.DB, log *logger.Log) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := db.Ping()
		if err != nil {
			panic("DATABASE PING ERROR IN HELLOROUTE!")
		}
		w.Write([]byte("Hello World!"))
	}
}
