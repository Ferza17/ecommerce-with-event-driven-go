package repository

type ProductRedisRepositoryCommand interface {
}

type ProductRedisRepositoryQuery interface {
}

type ProductRedisRepositoryStore interface {
	ProductRedisRepositoryCommand
	ProductRedisRepositoryQuery
}
