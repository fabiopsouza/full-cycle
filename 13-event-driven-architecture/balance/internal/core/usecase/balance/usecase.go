package balance

import (
	"encoding/json"
	"fmt"
	"github.com/fabiopsouza/balance/internal/core/domain/balance"
	balancePort "github.com/fabiopsouza/balance/internal/core/port/balance"
)

type UseCaseHandler struct {
	repository balancePort.Repository
}

func NewUseCase(repository balancePort.Repository) balancePort.UseCase {
	return &UseCaseHandler{
		repository,
	}
}

func (h UseCaseHandler) Save(msg []byte) error {
	var balanceEvent balance.UpdateEvent

	err := json.Unmarshal(msg, &balanceEvent)
	if err != nil {
		return err
	}

	err = h.repository.Save(balanceEvent.Payload)
	if err != nil {
		return err
	}

	fmt.Println("Balance saved successfully")
	return nil
}
