package eq

type ConfigMap map[string]any

type Message struct {
	Topic string
	Value []byte
}

type Producer[T any] interface {
	Inner() *T
	Close()
	Produce(*Message) error
}

type Consumer[T any] interface {
	Inner() *T
	Close()
	Subscribe([]string, chan bool)
}
