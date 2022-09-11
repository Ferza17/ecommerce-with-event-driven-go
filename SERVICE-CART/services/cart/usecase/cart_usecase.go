package usecase

import (
	"github.com/Ferza17/event-driven-cart-service/services/cart/repository"
)

type cartUseCase struct {
	cartMongoDBRepository     repository.CartMongoDBRepositoryStore
	cartCassandraDBRepository repository.CartCassandraDBRepositoryStore
	cartRedisRepository       repository.CartRedisRepositoryStore
}

func NewCartUseCase(
	cartMongoDBRepository repository.CartMongoDBRepositoryStore,
	cartCassandraDBRepository repository.CartCassandraDBRepositoryStore,
	cartRedisRepository repository.CartRedisRepositoryStore,
) CartUseCaseStore {
	return &cartUseCase{
		cartMongoDBRepository:     cartMongoDBRepository,
		cartCassandraDBRepository: cartCassandraDBRepository,
		cartRedisRepository:       cartRedisRepository,
	}
}
