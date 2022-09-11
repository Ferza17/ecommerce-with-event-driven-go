package utils

type key string

const (
	RabbitmqAmqpContextKey key = "rabbitmq_context_key"
	RedisDBContextKey      key = "redis_db_context_key"
	MongodbContextKey      key = "mongo_db_context_key"
	TracerContextKey       key = "tracer_context_key"
	CassandraDBContextKey  key = "cassandra_db_context_key"
	CartSpanContextKey     key = "cart_service_span_key"
	CartUseCaseContextKey  key = "Cart_use_case_context_key"
)
