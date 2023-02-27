package gofp

type Stream[T any] struct {
	val  T
	next func() *Stream[T]
}

func Append[T any](stream1 *Stream[T], stream2 *Stream[T]) *Stream[T] {
	if stream1 == nil {
		return stream2
	}
	return &Stream[T]{
		val: stream1.val,
		next: func() *Stream[T] {
			return Append(stream1.next(), stream2)
		},
	}
}

func FlatMap[T, U any](stream *Stream[T], f func(T) *Stream[U]) *Stream[U] {
	if stream == nil {
		return nil
	}
	return Append(f(stream.val), FlatMap(stream.next(), f))
}
func Map[T, U any](s *Stream[T], f func(T) U) *Stream[U] {
	if s == nil {
		return nil
	}
	return &Stream[U]{
		val: f(s.val),
		next: func() *Stream[U] {
			return Map(s.next(), f)
		},
	}
}
func (s *Stream[T]) Filter(f func(T) bool) *Stream[T] {
	for ; s != nil; s = s.next() {
		if f(s.val) {
			return &Stream[T]{
				val: s.val,
				next: func() *Stream[T] {
					return s.next().Filter(f)
				},
			}
		}
	}
	return s
}

func (s *Stream[T]) ForEach(f func(T)) {
	for ; s != nil; s = s.next() {
		f(s.val)
	}
}
func (s *Stream[T]) ToSlice() []T {
	var result []T
	s.ForEach(func(value T) {
		result = append(result, value)
	})
	return result
}

func NewStreamFromSlice[T any](arr []T) *Stream[T] {
	if len(arr) == 0 {
		return nil
	}
	return &Stream[T]{
		val: arr[0],
		next: func() *Stream[T] {
			return NewStreamFromSlice(arr[1:])
		},
	}
}
