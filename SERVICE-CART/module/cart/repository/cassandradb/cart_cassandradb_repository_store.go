package cassandradb

type CartCassandraDBTransactionCreator interface {
}

type CartCassandraDBRepositoryCommand interface {
}

type CartCassandraDBRepositoryQuery interface {
}

type CartCassandraDBRepositoryStore interface {
	CartCassandraDBTransactionCreator
	CartCassandraDBRepositoryCommand
	CartCassandraDBRepositoryQuery
}
