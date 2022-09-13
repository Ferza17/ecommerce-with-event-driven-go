package utils

type (
	Event      string
	EventSaga  string
	SagaStatus string
)

const (
	SagaStatusSuccess SagaStatus = "SUCCESS"
	SagaStatusFailed  SagaStatus = "FAILED"

	UpdateProductStock Event = "UPDATE-PRODUCT-STOCK-EVENT"

	UpdateProductStockEventSaga EventSaga = "UPDATE-PRODUCT-STOCK-EVENT-SAGA"
)
