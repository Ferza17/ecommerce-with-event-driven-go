package usecase

type CartUseCaseCommand interface {
}

type CartUseCaseQuery interface {
}

type CartUseCaseStore interface {
	CartUseCaseCommand
	CartUseCaseQuery
}
