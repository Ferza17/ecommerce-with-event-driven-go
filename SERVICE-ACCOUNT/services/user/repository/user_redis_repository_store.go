package repository

type UserRedisRepositoryCommand interface {
}

type UserRedisRepositoryQuery interface {
}

type UserRedisRepositoryStore interface {
	UserRedisRepositoryCommand
	UserRedisRepositoryQuery
}
