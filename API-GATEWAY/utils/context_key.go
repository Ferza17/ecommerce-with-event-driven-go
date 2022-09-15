package utils

type ContextKey string

const (
	HeadersContextKey        ContextKey = "headers"
	HostInfoKey              ContextKey = "host_info"
	RabbitmqAmqpContextKey   ContextKey = "rabbitmq_context_key"
	TracerContextKey         ContextKey = "tracer_context_key"
	CassandraDBContextKey    ContextKey = "cassandra_db_context_key"
	APIGatewaySpanContextKey ContextKey = "api_gateway_span_key"
	TokenIdentityContextKey  ContextKey = "token_identity_context_key"

	UserServiceGrpcClientContextKey    ContextKey = "user_service_grpc_client_context_key"
	CartServiceGrpcClientContextKey    ContextKey = "cart_service_grpc_client_context_key"
	ProductServiceGrpcClientContextKey ContextKey = "product_service_grpc_client_context_key"

	UserSchemaConfigContextKey ContextKey = "user_schema_config_context_key"
	AuthSchemaConfigContextKey ContextKey = "auth_schema_config_context_key"
	CartSchemaConfigContextKey ContextKey = "cart_schema_config_context_key"

	UserUseCaseContextKey    ContextKey = "user_use_case_context_key"
	CartUseCaseContextKey    ContextKey = "cart_use_case_context_key"
	ProductUseCaseContextKey ContextKey = "product_use_case_context_key"
)
