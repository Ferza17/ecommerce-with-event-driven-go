package usecase

type ProductUseCaseCommand interface {
}

type ProductUseCaseQuery interface {
}

type ProductUseCaseCompensate interface {
}

type ProductUseCaseStore interface {
	ProductUseCaseCompensate
	ProductUseCaseCommand
	ProductUseCaseQuery
}
