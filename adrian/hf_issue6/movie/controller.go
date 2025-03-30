package movie

import (
	"database/sql"
	"encoding/json"
	"github.com/go-chi/chi"
	"net/http"
)

const (
	CTypeHeader = "Content-Type"
	JSONType    = "application/json"
)

type MovieController struct {
	movieDB *DB
}

func NewMovieController(db *sql.DB) *MovieController {
	return &MovieController{movieDB: NewDB(db)}
}

func (mc *MovieController) GetMovie(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	movie, err := mc.movieDB.GetMovie(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	response, err := json.Marshal(movie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set(CTypeHeader, JSONType)
	w.Write(response)

}

func (mc *MovieController) CreateMovie(w http.ResponseWriter, r *http.Request) {
	var movie Movie

	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	var newId string
	newId, err = mc.movieDB.CreateMovie(movie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Location", "/movies/"+newId)
	w.WriteHeader(http.StatusCreated)

}

func fillEmptyFields(movie *Movie, origMovie Movie) {
	if movie.Title == "" {
		movie.Title = origMovie.Title
	}
	if movie.ReleaseDate == "" {
		movie.ReleaseDate = origMovie.ReleaseDate
	}
	if movie.ImdbId == "" {
		movie.ImdbId = origMovie.ImdbId
	}

	if movie.Director == "" {
		movie.Director = origMovie.Director
	}

	if movie.Writer == "" {
		movie.Writer = origMovie.Writer
	}

	if movie.Stars == "" {
		movie.Stars = origMovie.Stars
	}
}

func (mc *MovieController) UpdateMovie(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var movie Movie
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var origMovie Movie
	origMovie, err = mc.movieDB.GetMovie(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	fillEmptyFields(&movie, origMovie)
	movie.ID = id

	err = mc.movieDB.UpdateMovie(movie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (mc *MovieController) DeleteMovie(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := mc.movieDB.DeleteMovie(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
