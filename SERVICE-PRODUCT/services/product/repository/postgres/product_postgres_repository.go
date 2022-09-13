package postgres

import "github.com/jmoiron/sqlx"

type productPostgresRepository struct {
	dbRead  *sqlx.DB
	dbWrite *sqlx.DB
}

func NewProductPostgresRepository(dbRead *sqlx.DB, dbWrite *sqlx.DB) ProductPostgresRepositoryStore {
	return &productPostgresRepository{
		dbRead:  dbRead,
		dbWrite: dbWrite,
	}
}
