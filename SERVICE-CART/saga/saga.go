package saga

import (
	"encoding/json"

	"github.com/RoseRocket/xerrs"

	"github.com/Ferza17/event-driven-cart-service/utils"
)

type saga struct {
}

type Step struct {
	TransactionId string           `json:"transactionId"`
	Counter       int              `json:"counter"`
	Status        utils.SagaStatus `json:"status"`
}

func NewSaga() SagaStore {
	return &saga{}
}

func (s *saga) ParseStringToStep(rawString string) (response Step, err error) {
	if err = json.Unmarshal([]byte(rawString), &response); err != nil {
		err = xerrs.Mask(err, utils.ErrInternalServerError)
	}
	return
}

func (s *saga) IsFailed(step *Step) bool {
	if step.Status == utils.SagaStatusFailed {
		return true
	}
	return false
}
