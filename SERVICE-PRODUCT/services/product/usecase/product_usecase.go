package usecase

import (
	"github.com/Ferza17/event-driven-product-service/services/product/repository"
)

type productUseCase struct {
	productElasticsearchRepository repository.ProductElasticsearchRepositoryStore
	productCassandraDBRepository   repository.ProductCassandraDBRepositoryStore
	productRedisRepository         repository.ProductRedisRepositoryStore
}

func NewProductUseCase(
	productElasticsearchRepository repository.ProductElasticsearchRepositoryStore,
	productCassandraDBRepository repository.ProductCassandraDBRepositoryStore,
	productCacheRepository repository.ProductRedisRepositoryStore,
) ProductUseCaseStore {
	return &productUseCase{
		productElasticsearchRepository: productElasticsearchRepository,
		productCassandraDBRepository:   productCassandraDBRepository,
		productRedisRepository:         productCacheRepository,
	}
}
