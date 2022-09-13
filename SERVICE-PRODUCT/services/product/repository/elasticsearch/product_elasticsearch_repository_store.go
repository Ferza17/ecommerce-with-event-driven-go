package elasticsearch

type ProductElasticsearchTransactionCreator interface {
}

type ProductElasticsearchRepositoryCommand interface {
}

type ProductElasticsearchRepositoryQuery interface {
}

type ProductElasticsearchRepositoryStore interface {
	ProductElasticsearchTransactionCreator
	ProductElasticsearchRepositoryCommand
	ProductElasticsearchRepositoryQuery
}
