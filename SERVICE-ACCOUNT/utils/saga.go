package utils

type (
	Event      string
	EventSaga  string
	SagaStatus string
)

const (
	SagaStatusSuccess SagaStatus = "SUCCESS"
	SagaStatusFailed  SagaStatus = "FAILED"

	CreateUserEvent Event = "CREATE-USER-EVENT"
	CreateCartEvent Event = "CREATE-CART-EVENT"

	CreateCartEventSaga EventSaga = "CREATE-CART-EVENT-SAGA"
)
