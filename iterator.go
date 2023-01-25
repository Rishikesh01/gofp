package fp_golang

type iteratorType string

const (
	ArrayList  iteratorType = "ArrayList"
	LinkedList iteratorType = "LinkedList"
)

type Iterator[T any] interface {
	HasNext() bool
	GetNext() T
	Type() iteratorType
}
