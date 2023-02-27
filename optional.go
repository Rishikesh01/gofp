package gofp

import "reflect"

type Optional[T any] interface {
	IsPresent() bool
	IsEmpty() bool
	IfPresent() Stream[T]
	OrElse(something T) T
	OrElseGet(func() T) T
	OrELseError(err error) (T, error)
	Get() T
}

type optional[T any] struct {
	val T
}

func OptionalOf[T any](index T) Optional[T] {
	return &optional[T]{val: index}
}

func (o *optional[T]) IsPresent() bool {
	if v := reflect.ValueOf(o.val); (v.Kind() == reflect.Ptr ||
		v.Kind() == reflect.Interface ||
		v.Kind() == reflect.Slice ||
		v.Kind() == reflect.Map ||
		v.Kind() == reflect.Chan ||
		v.Kind() == reflect.Func) && v.IsNil() {
		return false
	} else if v := reflect.ValueOf(o.val); v.IsZero() {
		return false
	}

	return true
}

func (o *optional[T]) IsEmpty() bool {
	return !o.IsPresent()
}

func (o *optional[T]) IfPresent() Stream[T] {
	//TODO implement me
	panic("implement me")
}

func (o *optional[T]) OrElse(something T) T {
	//TODO implement me
	panic("implement me")
}

func (o *optional[T]) OrElseGet(f func() T) T {
	//TODO implement me
	panic("implement me")
}

func (o *optional[T]) OrELseError(err error) (T, error) {
	//TODO implement me
	panic("implement me")
}

func (o *optional[T]) Get() T {
	//TODO implement me
	panic("implement me")
}
