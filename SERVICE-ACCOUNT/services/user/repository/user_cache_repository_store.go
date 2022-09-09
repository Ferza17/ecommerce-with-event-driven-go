package repository

type UserCacheRepositoryWriter interface {
}

type UserCacheRepositoryReader interface {
}

type UserCacheRepositoryStore interface {
	UserCacheRepositoryWriter
	UserCacheRepositoryReader
}
