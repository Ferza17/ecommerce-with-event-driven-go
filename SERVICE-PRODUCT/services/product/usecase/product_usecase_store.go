package usecase

type ProductUseCaseCommand interface {
}

type ProductUseCaseQuery interface {
}

type ProductUseCaseStore interface {
	ProductUseCaseCommand
	ProductUseCaseQuery
}
