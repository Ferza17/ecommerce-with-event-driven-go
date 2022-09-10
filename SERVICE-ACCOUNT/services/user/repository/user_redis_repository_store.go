package repository

type UserCacheRepositoryCommand interface {
}

type UserCacheRepositoryQuery interface {
}

type UserCacheRepositoryStore interface {
	UserCacheRepositoryCommand
	UserCacheRepositoryQuery
}
