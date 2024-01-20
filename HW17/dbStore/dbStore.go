package dbStore

import (
	r "main/dbStore/repositories"
)

type DbStore struct {
	Products   r.ProductRepository
	Skus       r.SkuRepository
	Attributes r.AttributeRepository
}

var storage DbStore = DbStore{
	Products:   r.ProductRepository{},
	Skus:       r.SkuRepository{},
	Attributes: r.AttributeRepository{},
}

func GetNewDbStore(connString string) DbStore {
	return DbStore{
		Products:   r.ProductRepository{BaseRepository: r.BaseRepository{ConnectionString: connString}},
		Skus:       r.SkuRepository{BaseRepository: r.BaseRepository{ConnectionString: connString}},
		Attributes: r.AttributeRepository{BaseRepository: r.BaseRepository{ConnectionString: connString}},
	}
}
