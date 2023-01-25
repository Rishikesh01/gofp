package fp_golang

import "fmt"

type List[T any] interface {
	Add(x T)
	Remove(x int)
	Get(x int) T
	Length() int
	ToStream() Stream[T]
}

func ArrayToList[T any](arr ...T) List[T] {
	return &arrayList[T]{array: arr}
}
func NewList[T any]() List[T] {
	return &arrayList[T]{}
}

func (a *arrayList[T]) ToIterator() Iterator[T] {
	return a
}

func (a arrayList[T]) Type() iteratorType {
	return ArrayList
}
func (a *arrayList[T]) HasNext() bool {
	return a.index < len(a.array)
}

func (a *arrayList[T]) GetNext() T {
	if a.HasNext() {
		val := a.array[a.index]
		a.index++
		return val
	}
	fmt.Println("v")
	a.index = 0

	return *new(T)
}

type arrayList[T any] struct {
	index int
	array []T
}

func (a arrayList[T]) ToStream() Stream[T] {
	return &a
}

func (a *arrayList[T]) Length() int {
	return len(a.array)
}

func (a *arrayList[T]) Get(x int) T {
	if len(a.array) <= x {
		panic("Invalid index")
	}
	return a.array[x]
}

func (a *arrayList[T]) Add(element T) {
	a.array = append(a.array, element)
}

func (a *arrayList[T]) Remove(index int) {
	a.array = append(a.array[:index], a.array[index+1:]...)
}
