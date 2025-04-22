package service

import (
	"bookstore/book"
	"bookstore/config"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/doug-martin/goqu/v9"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// Webservice
type WebService struct {
	ServeMux *chi.Mux
	Cfg      *config.ServerConfig
	DB       *goqu.Database
	AppLog   *slog.Logger
}

type Option func(*WebService)

// A NewWebService létrehozza a WebService struktúra új példányát.
//
// Parameters:
// - cfg: A HTTP kiszolgáló config struktúrája
// - appLog: AppLogger példány mutatója
// - db: sqlBuilder példány mutatója
//
// Returns:
// - A WebService struktúra új példány mutatója
func NewWebService(db *goqu.Database, options ...Option) *WebService {
	ws := &WebService{
		ServeMux: chi.NewRouter(),
		Cfg:      &config.ServerConfig{Host: "localhost", Port: 8080}, // Alapértelmezett konfig ha nem adunk semmit
		DB:       db,
		AppLog:   slog.Default(), // alapértelmezett logger
	}

	// options patterns alkalmazása
	for _, option := range options {
		option(ws)
	}

	ws.configureService()

	return ws
}

func WithCfg(cfg *config.ServerConfig) Option {
	return func(ws *WebService) {
		ws.Cfg = cfg
	}
}

func WithLogger(logger *slog.Logger) Option {
	return func(ws *WebService) {
		ws.AppLog = logger
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
		s.AppLog.Info("Request received",
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

// configureService beállítja a HTTP-kezelőket és a middleware beállításokat a webszolgáltatáshoz.
//
// Parameters:
// -
//
// Returns:
// -
func (s *WebService) configureService() {

	// middlewares
	s.ServeMux.Use(middleware.Recoverer)
	s.ServeMux.Use(s.loggingMiddleware)
	s.ServeMux.Use(s.corsMiddleware)

	// handlers
	s.ServeMux.Get("/books", s.getAllBooks)
	s.ServeMux.Get("/book/{id}", s.getBookByID)
	s.ServeMux.Post("/book", s.createBook)
	s.ServeMux.Put(("/book"), s.updateBook)
	s.ServeMux.Delete("/book/{id}", s.deleteBook)
}

// Az Execute elindítja a HTTP-kiszolgálót a ServeMux és a megadott port használatával.
//
// Parameters:
// - s: WebService példány mutatója
//
// Returns:
// - error: Hiba esetén a hibaüzenetet tartalmazó érték.
func (s *WebService) Execute() error {
	// http server indítása
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", s.Cfg.Host, s.Cfg.Port), s.ServeMux)
	if err != nil {
		return err
	}
	return nil
}

// sendJSONEncodeError JSON kódolás esetén küldendő hibaválasz a kliens felé
//
// Parameters:
// - w: http.ResponseWriter
// - err: a hibát tartalmazó interface
//
// Returns:
// -
func (s *WebService) sendJSONEncodeError(w http.ResponseWriter, err error) {
	errorMessage := fmt.Sprintf("JSON encoding failed: %s", err.Error())
	s.AppLog.Error(errorMessage)
	http.Error(w, errorMessage, http.StatusInternalServerError)
}

// getIDFromRequest függvény kinyeri az ID a reqest url alapján
//
// Parameters:
// - r:  HTTP request
//
// Returns:
// - int: a kinyert ID értéke
// - error: nil ha nincs hiba a kinyerés során, különben a hibát tartalmazó érték
func (s *WebService) getIDFromRequest(r *http.Request) (int, error) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	return id, err
}

// sendHttpError HTTP hibaválaszt küld a megadott HTTP hibakóddal és hibaüzenettel.
//
// Parameters:
// - w: http.ResponseWriter
// - httpErrorCode: http hibakód
// - err: a hibát tartalmazó érték
//
// Returns:
// -
func (s *WebService) sendHttpError(w http.ResponseWriter, httpErrorCode int, err error) {
	errorMessage := fmt.Sprintf("%s: %d", err.Error(), httpErrorCode)
	s.AppLog.Error(errorMessage)
	http.Error(w, errorMessage, httpErrorCode)
}

// getAllBooks lekéri az összes könyvet s.BookStore-ból
//
// Parameters:
// - w: a http.ResponseWriter-ben adjuk visza a könyveket tartalmazó JSON-t
// - r: tartalmazza a http kérést
//
// Returns:
func (s *WebService) getAllBooks(w http.ResponseWriter, r *http.Request) {

	// lekérjük az adatbázisból az összes könyvet
	bookService := book.NewBookService(s.DB)
	books, err := bookService.GetAllBooks()
	if err != nil {
		s.sendHttpError(w, http.StatusInternalServerError, err)
		return
	}

	// elküldjük az összes könyvet
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(books)
	if err != nil {
		s.sendJSONEncodeError(w, err)
		return
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

	// Kinyerjük az útvonal alapján a könyv azonosítóját
	id, err := s.getIDFromRequest(r)
	if err != nil {
		s.sendHttpError(w, http.StatusBadRequest, err)
		return
	}

	// lekérjük az adatbázisból az id alapján a könyvet
	bookService := book.NewBookService(s.DB)
	book, err := bookService.GetBookByID(id)
	if err != nil {
		s.sendHttpError(w, http.StatusBadRequest, err)
		return
	}

	// elküldjük a megtalált könyvet
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(book)
	if err != nil {
		s.sendJSONEncodeError(w, err)
		return
	}
}

// A createBook létrehoz egy új bejegyzést a body-ban létrehozott bejegyzés alapján
//
// Parameters:
// - w: a http.ResponseWriter-ben adjuk visza a megtalát könyvet tartalmazó JSON-t
// - r: tartalmazza a http kérést
//
// Returns:
// -
func (s *WebService) createBook(w http.ResponseWriter, r *http.Request) {

	// body-ból lekérjük a book struktúrát a oneBook-ba
	var oneBook book.Book
	err := json.NewDecoder(r.Body).Decode(&oneBook)
	if err != nil {
		s.sendHttpError(w, http.StatusBadRequest, err)
		return
	}

	// új könyv létrehozása a oneBook alapján
	bookService := book.NewBookService(s.DB)
	newId, err := bookService.CreateBook(oneBook)
	if err != nil {
		s.sendHttpError(w, http.StatusInternalServerError, err)
		return
	}

	// visszaküldjük az új id-t
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	var resp book.BookResponseMessage
	resp.Message = fmt.Sprintf("sikeres felvitel az adatbazisba (ID: %d)", newId)
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		s.sendJSONEncodeError(w, err)
		return
	}
}

// Az updateBook módosít egy bejegyzést a body-ban létrehozott bejegyzés alapján
//
// Parameters:
// - w: a http.ResponseWriter-ben adjuk visza a megtalát könyvet tartalmazó JSON-t
// - r: tartalmazza a http kérést
//
// Returns:
// -
func (s *WebService) updateBook(w http.ResponseWriter, r *http.Request) {

	// body-ból lekérjük a book struktúrát a oneBook-ba
	var oneBook book.Book
	err := json.NewDecoder(r.Body).Decode(&oneBook)
	if err != nil {
		s.sendHttpError(w, http.StatusBadRequest, err)
		return
	}

	// módosítás az adatbázisban
	bookService := book.NewBookService(s.DB)
	err = bookService.UpdateBook(oneBook)
	if err != nil {
		s.sendHttpError(w, http.StatusInternalServerError, err)
		return
	}

	// módosítás sikerességének visszajelzése
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	var resp book.BookResponseMessage
	resp.Message = fmt.Sprintf("sikeres módosítás az adatbazisban (ID: %d)", oneBook.ID)
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		s.sendJSONEncodeError(w, err)
	}
}

// A deleteBook töröl egy könyvet az útvonalban megadott id azonosító alapján
//
// Parameters:
// - w: a http.ResponseWriter-ben adjuk visza a megtalát könyvet tartalmazó JSON-t
// - r: tartalmazza a http kérést
//
// Returns:
// -
func (s *WebService) deleteBook(w http.ResponseWriter, r *http.Request) {

	// kinyerjük az útvonal alapján a könyv azonosítóját
	id, err := s.getIDFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// törlés az adatbázisból
	bookService := book.NewBookService(s.DB)
	err = bookService.DeleteBook(id)
	if err != nil {
		s.sendHttpError(w, http.StatusInternalServerError, err)
		return
	}

	// törlés sikerességének visszajelzése
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	var resp book.BookResponseMessage
	resp.Message = fmt.Sprintf("sikeres törlés az adatbazisban (ID: %d)", id)
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		s.sendJSONEncodeError(w, err)
		return
	}
}
