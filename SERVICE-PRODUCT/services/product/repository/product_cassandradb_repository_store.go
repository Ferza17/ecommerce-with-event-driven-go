package repository

type ProductCassandraDBTransactionCreator interface {
}

type ProductCassandraDBRepositoryCommand interface {
}

type ProductCassandraDBRepositoryQuery interface {
}

type ProductCassandraDBRepositoryStore interface {
	ProductCassandraDBTransactionCreator
	ProductCassandraDBRepositoryCommand
	ProductCassandraDBRepositoryQuery
}
