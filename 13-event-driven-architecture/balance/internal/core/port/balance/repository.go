package balance

import "github.com/fabiopsouza/balance/internal/core/domain/balance"

type Repository interface {
	Save(payload balance.UpdateEventPayload) error
}
