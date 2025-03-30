package movie

type Movie struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	ReleaseDate string `json:"release_date" db:"release_date"`
	ImdbId      string `json:"imdb_id" db:"imdb_id"`
	Director    string `json:"director"`
	Writer      string `json:"writer"`
	Stars       string `json:"stars"`
}
