package postgres

type ProductPostgresRepositoryCommand interface {
}

type ProductPostgresRepositoryQuery interface {
}

type ProductPostgresRepositoryStore interface {
	ProductPostgresRepositoryCommand
	ProductPostgresRepositoryQuery
}
