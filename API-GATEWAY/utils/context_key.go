package utils

type key string

const (
	HeadersContextKey        key = "headers"
	HostInfoKey              key = "host_info"
	RabbitmqAmqpContextKey   key = "rabbitmq_context_key"
	TracerContextKey         key = "tracer_context_key"
	CassandraDBContextKey    key = "cassandra_db_context_key"
	APIGatewaySpanContextKey key = "api_gateway_span_key"

	UserServiceGrpcClientContextKey    key = "user_service_grpc_client_context_key"
	CartServiceGrpcClientContextKey    key = "cart_service_grpc_client_context_key"
	ProductServiceGrpcClientContextKey key = "product_service_grpc_client_context_key"

	UserUseCaseContextKey    key = "user_use_case_context_key"
	CartUseCaseContextKey    key = "cart_use_case_context_key"
	ProductUseCaseContextKey key = "product_use_case_context_key"
)
