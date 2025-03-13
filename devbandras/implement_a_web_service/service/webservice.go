package service

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
	"web_service/logger"
)

// Webservice
type WebService struct {
	ServeMux  *http.ServeMux
	Port      int
	BookStore *BookStore
	AppLog    *logger.AppLogger
}

// A NewWebService létrehozza a WebService struktúra új példányát.
//
// Parameters:
// - port: A HTTP kiszolgáló portja
// - appLog: AppLogger példányra mutatója
// - bookStore: BookStore példányra mutatója
//
// Returns:
// - A WebService struktúra új példány mutatója
func NewWebService(port int, bookStore *BookStore, appLog *logger.AppLogger) *WebService {
	return &WebService{
		ServeMux:  http.NewServeMux(),
		Port:      port,
		BookStore: bookStore,
		AppLog:    appLog,
	}
}

// A loggingMiddleware naplózza a HTTP kéréseket.
//
// Parameters:
// - next: http.Handler
//
// Returns:
// - http.Handler: Az új http.Handler amely már tartalmazza a naplózási funkciókat.
func (s *WebService) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Logoljuk a request adatokat
		s.AppLog.Logger.Info("Request received",
			slog.String("method", r.Method),
			slog.String("path", r.URL.Path),
			slog.String("remote_addr", r.RemoteAddr),
		)
		next.ServeHTTP(w, r)
	})
}

// corsMiddleware CORS beállításokat ad a kéréseknek.
//
// Parameters:
// - next: http.Handler
//
// Returns:
// - http.Handler: Az új http.Handler amely már tartalmazza a CORS ferjléceket
func (s *WebService) corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// getAllBooks lekéri az összes könyvet s.BookStore-ból
//
// Parameters:
// - w: a http.ResponseWriter-ben adjuk visza a könyveket tartalmazó JSON-t
// - r: tartalmazza a http kérést
//
// Returns:
func (s *WebService) getAllBooks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		errorMessage := "Method not allowed"
		http.Error(w, errorMessage, http.StatusMethodNotAllowed)
		s.AppLog.Logger.Error(fmt.Sprintf("%s: %s", errorMessage, r.Method))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(s.BookStore.GetAllBooks())
	if err != nil {
		errorMessage := "JSON encoding failed"
		s.AppLog.Logger.Error(errorMessage)
		http.Error(w, fmt.Sprintf("%s: %s", "Internal server error", errorMessage), http.StatusInternalServerError)
	}
}

// A getBookByID lekér egy könyvet a köny ID azonosítója alapján az s.BookStore-ból.
//
// Parameters:
// - w: a http.ResponseWriter-ben adjuk visza a megtalát könyvet tartalmazó JSON-t
// - r: tartalmazza a http kérést
//
// Returns:
// -
func (s *WebService) getBookByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		errorMessage := "Method not allowed"
		http.Error(w, errorMessage, http.StatusMethodNotAllowed)
		s.AppLog.Logger.Error(fmt.Sprintf("%s: %s", errorMessage, r.Method))
		return
	}

	// Kinyerjük az utvoal alapján a könyv azonosítóját
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	// Convert to integer
	id, err := strconv.Atoi(parts[2])
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	// lekérjük az adatbázisból az id alapján a könyvet
	book, err := s.BookStore.GetBookByID(id)
	if err != nil {
		warningMessage := err.Error()
		http.Error(w, warningMessage, http.StatusNotFound)
		s.AppLog.Logger.Warn(warningMessage)
		return
	}

	// elküldjük a megtalált könyvet
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(book)
	if err != nil {
		errorMessage := "JSON encoding failed"
		s.AppLog.Logger.Error(errorMessage)
		http.Error(w, fmt.Sprintf("%s: %s", "Internal server error", errorMessage), http.StatusInternalServerError)
	}
}

// configureService beállítja a HTTP-kezelőket és a middleware beállításokat a webszolgáltatáshoz.
//
// Parameters:
// -
//
// Returns:
// - http.Handler: Egy http.Handler, mely tartalmazza a regisztrált handlereket és middleware beállításokat
func (s *WebService) configureService() http.Handler {

	// handlers
	s.ServeMux.HandleFunc("/books", s.getAllBooks)
	s.ServeMux.HandleFunc("/books/", s.getBookByID)

	// middlewares
	returHandler := s.loggingMiddleware(s.ServeMux)
	return s.corsMiddleware(returHandler)
}

// Az Execute elindítja a HTTP-kiszolgálót a ServeMux és a megadott port használatával.
//
// Parameters:
// - s: WebService példány mutatója
//
// Returns:
// - error: Hiba esetén a hibaüzenetet tartalmazó érték.
func (s *WebService) Execute() error {

	// wrappedMux létrehozása, mely tartaslmazza a middleware-t és a handler-t
	wrappedMux := s.configureService()

	// http server indítása
	err := http.ListenAndServe(fmt.Sprintf(":%d", s.Port), wrappedMux)
	if err != nil {
		return err
	}
	return nil
}
