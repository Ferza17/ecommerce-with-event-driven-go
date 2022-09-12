package saga

type SagaStore interface {
	ParseStringToStep(rawString string) (response Step, err error)
	IsFailed(step *Step) bool
}
