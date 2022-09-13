package usecase

import (
	"github.com/Ferza17/event-driven-product-service/saga"
	"github.com/Ferza17/event-driven-product-service/services/product/repository/cassandradb"
	"github.com/Ferza17/event-driven-product-service/services/product/repository/elasticsearch"
	"github.com/Ferza17/event-driven-product-service/services/product/repository/postgres"
	"github.com/Ferza17/event-driven-product-service/services/product/repository/redis"
)

type productUseCase struct {
	productElasticsearchRepository elasticsearch.ProductElasticsearchRepositoryStore
	productCassandraDBRepository   cassandradb.ProductCassandraDBRepositoryStore
	productPostgresRepository      postgres.ProductPostgresRepositoryStore
	productRedisRepository         redis.ProductRedisRepositoryStore
	productSaga                    saga.SagaStore
}

func NewProductUseCase(
	productElasticsearchRepository elasticsearch.ProductElasticsearchRepositoryStore,
	productCassandraDBRepository cassandradb.ProductCassandraDBRepositoryStore,
	productPostgresRepository postgres.ProductPostgresRepositoryStore,
	productCacheRepository redis.ProductRedisRepositoryStore,
	productSaga saga.SagaStore,
) ProductUseCaseStore {
	return &productUseCase{
		productElasticsearchRepository: productElasticsearchRepository,
		productCassandraDBRepository:   productCassandraDBRepository,
		productPostgresRepository:      productPostgresRepository,
		productRedisRepository:         productCacheRepository,
		productSaga:                    productSaga,
	}
}
