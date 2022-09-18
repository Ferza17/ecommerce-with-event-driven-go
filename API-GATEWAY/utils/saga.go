package utils

type (
	Event      string
	SagaStatus string
)

const (
	CreateUserEvent Event = "CREATE-USER-EVENT"
	UserNewState    Event = "USER-NEW-STATE"
)
