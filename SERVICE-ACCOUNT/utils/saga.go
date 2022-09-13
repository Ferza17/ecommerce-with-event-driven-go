package utils

type (
	Event      string
	EventSaga  string
	SagaStatus string
)

const (
	SagaStatusSuccess SagaStatus = "SUCCESS"
	SagaStatusFailed  SagaStatus = "FAILED"

	NewUserEvent Event = "NEW-USER-EVENT"
	NewCartEvent Event = "NEW-CART-EVENT"

	NewCartEventSaga EventSaga = "NEW-CART-EVENT-SAGA"
)
