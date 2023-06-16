package eq

type EventQueue[T, P, C any] interface {
    NewProducer(T) (*P, error)
    NewConsumer(T) (*C, error)
}
