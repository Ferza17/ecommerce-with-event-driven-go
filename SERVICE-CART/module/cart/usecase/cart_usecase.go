package usecase

import (
	cartPublisher "github.com/Ferza17/event-driven-cart-service/module/cart/publisher"
	"github.com/Ferza17/event-driven-cart-service/module/cart/repository/cassandradb"
	"github.com/Ferza17/event-driven-cart-service/module/cart/repository/mongodb"
	"github.com/Ferza17/event-driven-cart-service/module/cart/repository/redis"
	"github.com/Ferza17/event-driven-cart-service/saga"
)

type cartUseCase struct {
	cartMongoDBRepository     mongodb.CartMongoDBRepositoryStore
	cartCassandraDBRepository cassandradb.CartCassandraDBRepositoryStore
	cartRedisRepository       redis.CartRedisRepositoryStore
	cartPublisher             cartPublisher.CartPublisherStore
	cartSaga                  saga.SagaStore
}

func NewCartUseCase(
	cartMongoDBRepository mongodb.CartMongoDBRepositoryStore,
	cartCassandraDBRepository cassandradb.CartCassandraDBRepositoryStore,
	cartRedisRepository redis.CartRedisRepositoryStore,
	cartPublisher cartPublisher.CartPublisherStore,
	cartSaga saga.SagaStore,
) CartUseCaseStore {
	return &cartUseCase{
		cartMongoDBRepository:     cartMongoDBRepository,
		cartCassandraDBRepository: cartCassandraDBRepository,
		cartRedisRepository:       cartRedisRepository,
		cartPublisher:             cartPublisher,
		cartSaga:                  cartSaga,
	}
}
