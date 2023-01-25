package fp_golang

type Stream[T any] interface {
	Iterator[T]
	Filter(func(T) bool) Stream[T]
	Map(func(T) T) Stream[T]
	FlatMap(func(T) Stream[T]) Stream[T]
	ForEach(func(T))
	ToSlice() []T
}

func Map[IT, OT any](input Stream[IT], mapper func(IT) OT) Stream[OT] {
	stream := NewList[OT]()
	for input.HasNext() {
		stream.Add(mapper(input.GetNext()))
	}

	return stream.ToStream()
}

func FlatMap[IT, OT any](input Stream[IT], mapper func(IT) Stream[OT]) Stream[OT] {
	stream := NewList[OT]()
	for input.HasNext() {
		val := mapper(input.GetNext())
		for val.HasNext() {
			stream.Add(val.GetNext())
		}
	}

	return stream.ToStream()
}

func (a *arrayList[T]) Filter(f func(T) bool) Stream[T] {
	filtered := arrayList[T]{}
	for a.HasNext() {
		val := a.GetNext()
		if f(val) {
			filtered.array = append(filtered.array, val)
		}

	}

	return &filtered
}

func (a *arrayList[T]) Map(f func(T) T) Stream[T] {
	return Map[T, T](a, f)
}

func (a *arrayList[T]) FlatMap(f func(T) Stream[T]) Stream[T] {
	return FlatMap[T, T](a, f)
}

func (a *arrayList[T]) ForEach(f func(T)) {
	for a.HasNext() {
		f(a.GetNext())
	}
}

func (a *arrayList[T]) ToSlice() []T {
	return a.array
}
