package db

type Database[T any] interface {
    Inner() T
    Migrate(any) error
}
