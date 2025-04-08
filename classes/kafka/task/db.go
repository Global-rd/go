package task

import (
	"database/sql"
	"errors"
	"fmt"

	. "github.com/doug-martin/goqu/v9"
)

var TaskTable = S("app").Table("task")

type DB struct {
	db *sql.DB
}

func (d DB) Get() (result []Task, err error) {
	rows, err := d.db.Query(`SELECT id, name, description FROM app.task;`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var task Task

		rows.Scan(&task.ID, &task.Name, &task.Description)
		result = append(result, task)
	}

	return result, nil
}

func (d DB) Create(task Task) error {
	sqlBuilderDB := New("postgres", d.db)

	ds := sqlBuilderDB.
		Insert(TaskTable).
		Rows(task)

	sql, _, _ := ds.ToSQL()
	fmt.Println("sql", sql)

	result, err := ds.Executor().Exec()

	affected, _ := result.RowsAffected()
	if affected < 1 {
		return errors.New("task creation did not happen")
	}

	return err
}
