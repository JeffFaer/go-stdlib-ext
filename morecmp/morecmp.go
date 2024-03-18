// morecmp implements extensions to the standard cmp library.
package morecmp

import "cmp"

// CmpFunc is a convience that represents the comparison functions used by
// APIs like slices.SortFunc.
type CmpFunc[T any] func(T, T) int

// AndThen returns a new CmpFunc that first checks this CmpFunc, and if the
// comparison is zero, checks the next CmpFuncs in order.
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

// Reversed returns a new CmpFunc that orders the opposite of this one.
func (f CmpFunc[T]) Reversed() CmpFunc[T] {
	return func(a, b T) int {
		return -f(a, b)
	}
}

// LessFunc converts this CmpFunc into a LessFunc.
func (f CmpFunc[T]) LessFunc() LessFunc[T] {
	return func(a, b T) bool {
		return f(a, b) < 0
	}
}

// LessFunc is a convenience that repesents the comparison functions used by
// older APIs like sort.Slice.
type LessFunc[T any] func(T, T) bool

// Comparing creates a CmpFunc that compares Ts by first converting them into
// Es.
func Comparing[T any, E cmp.Ordered](f func(T) E) CmpFunc[T] {
	return ComparingFunc(f, cmp.Compare[E])
}

// Comparing creates a CmpFunc that compares Ts by first converting them into
// Es, then then checking those Es with the provided CmpFunc.
func ComparingFunc[T, E any](f func(T) E, g CmpFunc[E]) CmpFunc[T] {
	return func(a, b T) int {
		e, f := f(a), f(b)
		return g(e, f)
	}
}

// TrueFirst returns a CmpFunc[bool] that sorts true before false.
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

// FalseFirst returns a CmpFunc[bool] that sorts false before true.
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
