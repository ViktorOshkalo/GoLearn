package repositories

import (
	"database/sql"
	"log"
)

var connectionString string

func SetConnectionString(cs string) {
	connectionString = cs
}

// interface to combine sql.Db and sql.Tx Exec and Query call
type IDbExecutable interface {
	Exec(query string, attributes ...any) (sql.Result, error)
	Query(query string, attributes ...any) (*sql.Rows, error)
}

func Execute(handler func(sqldb IDbExecutable) error) error {
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	return handler(db)
}

func ExecuteWithResult[T any](handler func(sqldb IDbExecutable) (*T, error)) (*T, error) {
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	return handler(db)
}

func ExecuteTransactWithResult[T any](handler func(sqldb IDbExecutable) (*T, error)) (*T, error) {
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}

	res, err := handler(tx)
	if err != nil {
		tx.Rollback()
		return nil, err
	} else {
		tx.Commit()
	}

	return res, nil
}
