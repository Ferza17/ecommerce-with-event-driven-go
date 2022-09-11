package repository

type ProductCacheRepositoryCommand interface {
}

type ProductCacheRepositoryQuery interface {
}

type ProductCacheRepositoryStore interface {
	ProductCacheRepositoryCommand
	ProductCacheRepositoryQuery
}
