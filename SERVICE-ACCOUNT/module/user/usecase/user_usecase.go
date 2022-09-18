package usecase

import (
	"github.com/Ferza17/event-driven-account-service/module/user/publisher"
	"github.com/Ferza17/event-driven-account-service/module/user/repository/cassandradb"
	"github.com/Ferza17/event-driven-account-service/module/user/repository/mongodb"
	"github.com/Ferza17/event-driven-account-service/module/user/repository/redis"
	"github.com/Ferza17/event-driven-account-service/saga"
)

type userUseCase struct {
	userMongoDBRepository     mongodb.UserMongoDBRepositoryStore
	userCassandraDBRepository cassandradb.UserCassandraDBRepositoryStore
	userRedisRepository       redis.UserRedisRepositoryStore
	userPub                   publisher.UserPublisherStore
	sagaStore                 saga.SagaStore
}

func NewUserUseCase(
	userNOSQLRepository mongodb.UserMongoDBRepositoryStore,
	userCassandraDBRepository cassandradb.UserCassandraDBRepositoryStore,
	userCacheRepository redis.UserRedisRepositoryStore,
	userPublisher publisher.UserPublisherStore,
	sagaStore saga.SagaStore,
) UserUseCaseStore {
	return &userUseCase{
		userMongoDBRepository:     userNOSQLRepository,
		userCassandraDBRepository: userCassandraDBRepository,
		userRedisRepository:       userCacheRepository,
		userPub:                   userPublisher,
		sagaStore:                 sagaStore,
	}
}
