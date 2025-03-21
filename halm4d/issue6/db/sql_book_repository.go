package db

import (
	"database/sql"
	"errors"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	"log/slog"
)

type SQLBookRepository struct {
	db         *sql.DB
	logger     *slog.Logger
	sqlBuilder *goqu.SelectDataset
}

func NewSQLBookRepository(db *sql.DB, logger *slog.Logger) *SQLBookRepository {
	return &SQLBookRepository{
		db:         db,
		logger:     logger,
		sqlBuilder: goqu.Dialect("mysql").From("books"),
	}
}

func (r *SQLBookRepository) Initialize() error {
	query := `CREATE TABLE IF NOT EXISTS books (
		id VARCHAR(36) PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		author VARCHAR(255) NOT NULL
	)`
	_, err := r.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (r *SQLBookRepository) GetAllBooks() (books []Book, err error) {
	query, _, err := r.sqlBuilder.Select("*").ToSQL()
	if err != nil {
		return books, err
	}
	r.logger.Info(query)
	rows, err := r.db.Query(query)
	if err != nil {
		return books, err
	}
	defer rows.Close()
	for rows.Next() {
		var book Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Author); err != nil {
			return books, err
		}
		books = append(books, book)
	}
	if err := rows.Err(); err != nil {
		return books, err
	}
	if len(books) == 0 {
		return books, NotFoundError
	}
	return books, nil
}

func (r *SQLBookRepository) GetBookByID(id string) (book Book, err error) {
	query, _, err := r.sqlBuilder.Where(goqu.Ex{"id": id}).ToSQL()
	if err != nil {
		return book, err
	}
	row := r.db.QueryRow(query)
	if err := row.Scan(&book.ID, &book.Title, &book.Author); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return book, NotFoundError
		}
		return book, err
	}
	return book, nil
}

func (r *SQLBookRepository) CreateBook(b Book) (book Book, err error) {
	exists, err := r.IsExists(b.ID)
	if err != nil {
		return book, err
	}
	if exists {
		return book, AlreadyExistsError
	}

	query, _, err := r.sqlBuilder.Insert().
		Rows(b).
		ToSQL()

	res, err := r.db.Exec(query)
	if err != nil {
		return book, err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return book, err
	}
	if affected == 0 {
		return book, FailedToCreateError
	}
	return b, nil
}

func (r *SQLBookRepository) UpdateBook(id string, b Book) (book Book, err error) {
	exists, err := r.IsExists(id)
	if err != nil {
		return book, err
	}
	if !exists {
		return book, NotFoundError
	}

	if b.ID == "" {
		b.ID = id
	}
	query, _, err := r.sqlBuilder.Update().
		Set(b).
		Where(goqu.Ex{"id": id}).
		ToSQL()

	res, err := r.db.Exec(query)
	if err != nil {
		return book, err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return book, err
	}
	if affected == 0 {
		return book, FailedToUpdateError
	}

	updatedBook, err := r.GetBookByID(id)
	if err != nil {
		return book, err
	}
	book = updatedBook
	return book, nil
}

func (r *SQLBookRepository) DeleteBook(id string) error {
	query, _, err := r.sqlBuilder.Delete().Where(goqu.Ex{"id": id}).ToSQL()
	_, err = r.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (r *SQLBookRepository) IsExists(id string) (bool, error) {
	_, err := r.GetBookByID(id)
	if errors.Is(err, NotFoundError) {
		return false, nil
	}
	return true, err
}
