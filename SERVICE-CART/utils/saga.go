package utils

type (
	Queue      string
	SagaQueue  string
	SagaStatus string
)

const (
	SagaStatusSuccess SagaStatus = "SUCCESS"
	SagaStatusFailed  SagaStatus = "FAILED"

	NewUserQueue Queue = "NEW-USER-QUEUE"
	NewCartQueue Queue = "NEW-CART-QUEUE"

	NewUserSagaQueue SagaQueue = "NEW-USER-SAGA-QUEUE"
)
