package balance

import "github.com/fabiopsouza/balance/internal/core/domain/balance"

type Repository interface {
	List(accountID string) ([]balance.Model, error)
	Save(payload balance.UpdateEventPayload) error
}
