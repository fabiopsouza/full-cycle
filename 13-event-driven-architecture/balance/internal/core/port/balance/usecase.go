package balance

import "github.com/fabiopsouza/balance/internal/core/domain/balance"

type UseCase interface {
	List() ([]balance.Model, error)
	Save(msg []byte) error
}
