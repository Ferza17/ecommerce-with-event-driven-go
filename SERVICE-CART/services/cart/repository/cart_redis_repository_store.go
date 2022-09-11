package repository

type CartRedisRepositoryCommand interface {
}

type CartRedisRepositoryQuery interface {
}

type CartRedisRepositoryStore interface {
	CartRedisRepositoryCommand
	CartRedisRepositoryQuery
}
