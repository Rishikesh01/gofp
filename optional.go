package gofp

import "reflect"

type Optional[T any] interface {
	IsPresent() bool
	IsEmpty() bool
	OrElse(something T) T
	OrElseGet(func() T) func() T
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

func (o *optional[T]) OrElse(something T) T {
	return something
}

func (o *optional[T]) OrElseGet(f func() T) func() T {
	return f
}

func (o *optional[T]) OrELseError(err error) (T, error) {
	return o.val, err
}

func (o *optional[T]) Get() T {
	return o.val
}
