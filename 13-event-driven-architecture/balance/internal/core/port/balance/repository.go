package balance

import "github.com/fabiopsouza/balance/internal/core/domain/balance"

type Repository interface {
	List() ([]balance.Model, error)
	Save(payload balance.UpdateEventPayload) error
}
