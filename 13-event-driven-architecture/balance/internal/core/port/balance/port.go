package balance

type UseCase interface {
	Save(msg []byte) error
}
