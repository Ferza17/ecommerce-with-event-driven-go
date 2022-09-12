package cassandradb

type UserCassandraDBTransactionCreator interface {
}

type UserCassandraDBRepositoryCommand interface {
}

type UserCassandraDBRepositoryQuery interface {
}

type UserCassandraDBRepositoryStore interface {
	UserCassandraDBTransactionCreator
	UserCassandraDBRepositoryCommand
	UserCassandraDBRepositoryQuery
}
