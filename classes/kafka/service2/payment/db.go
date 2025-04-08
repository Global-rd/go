package payment

import (
	"database/sql"
	"errors"
	"fmt"

	. "github.com/doug-martin/goqu/v9"
)

var PaymentTable = S("app").Table("payment")

type DB struct {
	db *sql.DB
}

func (d DB) Get() (result []Payment, err error) {
	rows, err := d.db.Query(`SELECT id, name, description FROM app.payment;`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var payment Payment

		rows.Scan(&payment.ID, &payment.Name, &payment.Description)
		result = append(result, payment)
	}

	return result, nil
}

func (d DB) Create(payment Payment) error {
	sqlBuilderDB := New("postgres", d.db)

	ds := sqlBuilderDB.
		Insert(PaymentTable).
		Rows(payment)

	sql, _, _ := ds.ToSQL()
	fmt.Println("sql", sql)

	result, err := ds.Executor().Exec()

	affected, _ := result.RowsAffected()
	if affected < 1 {
		return errors.New("payment creation did not happen")
	}

	return err
}
