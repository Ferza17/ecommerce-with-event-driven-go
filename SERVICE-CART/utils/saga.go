package utils

type (
	Event      string
	EventSaga  string
	SagaStatus string
)

const (
	SagaStatusSuccess SagaStatus = "SUCCESS"
	SagaStatusFailed  SagaStatus = "FAILED"

	CrateCartEvent       Event = "CREATE-CART-EVENT"
	CreateCartItemsEvent Event = "CREATE-CART-ITEMS-EVENT"
	DeleteCartItemsEvent Event = "DELETE-CART-ITEMS-EVENT"
	UpdateCartItemsEvent Event = "UPDATE-CART-ITEMS-EVENT"
	ReadCartNewState     Event = "READ-CART-NEW-STATE"

	CreateCartEventSaga      EventSaga = "CREATE-CART-EVENT-SAGA"
	CreateCartItemsEventSaga EventSaga = "INSERT-CART-ITEMS-EVENT-SAGA"
	DeleteCartItemsEventSaga EventSaga = "DELETE-CART-ITEMS-EVENT-SAGA"
	UpdateCartItemsEventSaga EventSaga = "UPDATE-CART-ITEMS-EVENT-SAGA"
)
