package utils

type (
	Event      string
	EventSaga  string
	SagaStatus string
)

const (
	SagaStatusSuccess SagaStatus = "SUCCESS"
	SagaStatusFailed  SagaStatus = "FAILED"

	//USER EVENT
	UserNewState    Event = "USER-NEW-STATE"
	CreateUserEvent Event = "CREATE-USER-EVENT"
	UpdateUserEvent Event = "UPDATE-USER-EVENT"

	//CART EVENT
	CreateCartEvent Event = "CREATE-CART-EVENT"

	CreateCartEventSaga EventSaga = "CREATE-CART-EVENT-SAGA"
)
