// morecmp implements extensions to the standard cmp library.
package morecmp

import "cmp"

type CmpFunc[T any] func(T, T) int

func (f CmpFunc[T]) AndThen(next ...CmpFunc[T]) CmpFunc[T] {
	return func(a, b T) int {
		if c := f(a, b); c != 0 {
			return c
		}
		for _, n := range next {
			if c := n(a, b); c != 0 {
				return c
			}
		}
		return 0
	}
}

func (f CmpFunc[T]) LessFunc() LessFunc[T] {
	return func(a, b T) bool {
		return f(a, b) < 0
	}
}

type LessFunc[T any] func(T, T) bool

func Comparing[T any, E cmp.Ordered](f func(T) E) CmpFunc[T] {
	return ComparingFunc(f, cmp.Compare[E])
}

func ComparingFunc[T, E any](f func(T) E, g CmpFunc[E]) CmpFunc[T] {
	return func(a, b T) int {
		e, f := f(a), f(b)
		return g(e, f)
	}
}

func TrueFirst() CmpFunc[bool] {
	return func(a, b bool) int {
		switch {
		case a == b:
			return 0
		case a:
			return -1
		default:
			return 1
		}
	}
}

func FalseFirst() CmpFunc[bool] {
	return func(a, b bool) int {
		switch {
		case a == b:
			return 0
		case !a:
			return -1
		default:
			return 1
		}
	}
}
