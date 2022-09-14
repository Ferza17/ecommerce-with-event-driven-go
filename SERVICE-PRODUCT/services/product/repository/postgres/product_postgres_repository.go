package postgres

import (
	"github.com/jmoiron/sqlx"
)

type (
	tableName                 string
	productPostgresRepository struct {
		dbRead  *sqlx.DB
		dbWrite *sqlx.DB
	}
)

const (
	tableProduct tableName = "products"
)

func NewProductPostgresRepository(dbRead *sqlx.DB, dbWrite *sqlx.DB) ProductPostgresRepositoryStore {
	return &productPostgresRepository{
		dbRead:  dbRead,
		dbWrite: dbWrite,
	}
}
