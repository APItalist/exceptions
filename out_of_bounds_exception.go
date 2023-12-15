package exceptions

import "fmt"

// OutOfBoundsException indicates that the caller attempted to access an index of an array, slice, map, or other
// indexable item with an index that was not present.
type OutOfBoundsException[T any] interface {
	Exception

	Index() T
}

// NewOutOfBoundsException create a new OutOfBoundsException
func NewOutOfBoundsException[T any](index T) OutOfBoundsException[T] {
	return &outOfBoundsException[T]{
		Inherit(nil, 1, DefaultExceptionStackDepth),
		index,
	}
}

type outOfBoundsException[T any] struct {
	BaseException

	index T
}

func (o outOfBoundsException[T]) Error() string {
	return fmt.Sprintf("index out of bounds: %v", o.index)
}

func (o outOfBoundsException[T]) Index() T {
	return o.index
}
