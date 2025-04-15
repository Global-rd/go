package movie

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"
)

const (
	dbFieldId          = "id"
	dbFieldTitle       = "title"
	dbFieldReleaseDate = "release_date"
	dbFieldImdbId      = "imdb_id"
	dbFieldDirector    = "director"
	dbFieldWriter      = "writer"
	dbFieldStars       = "stars"
)

var MoviesTable = goqu.S("movieapp").Table("movies")

type DB struct {
	db *sql.DB
}

func NewDB(db *sql.DB) *DB {
	return &DB{db: db}
}

func (d *DB) GetMovie(id string) (Movie, error) {
	sqlBuilder := goqu.New("postgres", d.db)
	query := sqlBuilder.From(MoviesTable).Select(dbFieldTitle, dbFieldReleaseDate, dbFieldImdbId, dbFieldDirector, dbFieldWriter, dbFieldStars).Where(goqu.C(dbFieldId).Eq(id))
	sqlString, _, _ := query.ToSQL()

	row := d.db.QueryRow(sqlString)
	movie := Movie{ID: id}

	err := row.Scan(
		&movie.Title,
		&movie.ReleaseDate,
		&movie.ImdbId,
		&movie.Director,
		&movie.Writer,
		&movie.Stars,
	)
	if err != nil {
		return Movie{}, err
	}

	return movie, nil
}

func (d *DB) CreateMovie(movie Movie) (string, error) {
	movie.ID = uuid.New().String()
	sqlBuilderDB := goqu.New("postgres", d.db)

	ds := sqlBuilderDB.
		Insert(MoviesTable).
		Rows(movie)

	//fmt.Println(ds.ToSQL())
	result, err := ds.Executor().Exec()
	if err != nil {
		return "", err
	}

	affected, _ := result.RowsAffected()
	if affected < 1 {
		return "", errors.New("failed to create new movie")
	}

	return movie.ID, nil
}

func (d *DB) DeleteMovie(id string) error {
	sqlBuilderDB := goqu.New("postgres", d.db)
	ds := sqlBuilderDB.
		Delete(MoviesTable).
		Where(goqu.C(dbFieldId).Eq(id))

	result, err := ds.Executor().Exec()
	if err != nil {
		return err
	}

	affected, _ := result.RowsAffected()
	if affected < 1 {
		return fmt.Errorf("failed to delete the movie with id %s", id)
	}

	return nil
}

func (d *DB) UpdateMovie(movie Movie) error {
	sqlBuilderDB := goqu.New("postgres", d.db)
	ds := sqlBuilderDB.
		Update(MoviesTable).
		Set(goqu.Record{
			dbFieldTitle:       movie.Title,
			dbFieldReleaseDate: movie.ReleaseDate,
			dbFieldImdbId:      movie.ImdbId,
			dbFieldDirector:    movie.Director,
			dbFieldWriter:      movie.Writer,
			dbFieldStars:       movie.Stars,
		}).Where(goqu.C(dbFieldId).Eq(movie.ID))

	result, err := ds.Executor().Exec()
	if err != nil {
		return err
	}
	affected, _ := result.RowsAffected()
	if affected < 1 {
		return fmt.Errorf("failed to update the movie with id %s", movie.ID)
	}

	return nil

}
