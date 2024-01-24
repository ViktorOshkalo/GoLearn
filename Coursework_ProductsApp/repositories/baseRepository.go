package repositories

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const provider string = "mysql"

type BaseRepository struct {
	ConnectionString string
}

// interface to combine sql.Db and sql.Tx Exec and Query calls
type IDbExecutable interface {
	Exec(query string, attributes ...any) (sql.Result, error)
	Query(query string, attributes ...any) (*sql.Rows, error)
	QueryRow(query string, attributes ...any) *sql.Row
}

// helper methods
func Execute(br BaseRepository, handler func(sqldb IDbExecutable) error) error {
	db, err := sql.Open("mysql", br.ConnectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	return handler(db)
}

func ExecuteWithResult[T any](br BaseRepository, handler func(sqldb IDbExecutable) (T, error)) (T, error) {
	db, err := sql.Open(provider, br.ConnectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	return handler(db)
}

func ExecuteTransactWithResult[T any](br BaseRepository, handler func(sqldb IDbExecutable) (T, error)) (T, error) {
	var resultDefault T

	db, err := sql.Open(provider, br.ConnectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return resultDefault, err
	}

	result, err := handler(tx)
	if err != nil {
		tx.Rollback()
		return resultDefault, err
	} else {
		tx.Commit()
	}

	return result, nil
}
