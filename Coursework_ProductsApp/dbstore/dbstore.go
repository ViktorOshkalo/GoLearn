package dbstore

import (
	"database/sql"
	"fmt"
	"log"
	"main/configuration"
	r "main/repositories"
)

type DbStore struct {
	Products   r.ProductRepository
	Skus       r.SkuRepository
	Attributes r.AttributeRepository
}

func GetNewDbStore(connString string) DbStore {
	return DbStore{
		Products:   r.ProductRepository{BaseRepository: r.BaseRepository{ConnectionString: connString}},
		Skus:       r.SkuRepository{BaseRepository: r.BaseRepository{ConnectionString: connString}},
		Attributes: r.AttributeRepository{BaseRepository: r.BaseRepository{ConnectionString: connString}},
	}
}

func (dbs *DbStore) Ping() {
	db, err := sql.Open(configuration.SqlProvider, configuration.ConnectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("error pinging database:", err)
	}

	fmt.Println("MySQL server is reachable.")
}
