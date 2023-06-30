package linear

type Linear[T any] interface {
	Len() int
	Insert(index int, element T) bool
	Remove(index int) bool
	Get(index int) T
	Empty() bool
}
