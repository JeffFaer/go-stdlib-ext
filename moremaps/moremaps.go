// moremaps impements extensions to the standard maps library.
package moremaps

import (
	"cmp"
	"slices"

	"golang.org/x/exp/maps"
)

// SortableKey is a type constraint that captures all map keys that are
// naturally sortable.
type SortableKey interface {
	comparable
	cmp.Ordered
}

// SortableKeys returns a new slice of the map's keys in their natural order.
func SortedKeys[K SortableKey, V any](m map[K]V) []K {
	return SortedKeysFunc(m, cmp.Compare[K])
}

// SortableKeysFunc returns a new slice of the map's keys ordered according to
// the provided function.
func SortedKeysFunc[K comparable, V any](m map[K]V, f func(K, K) int) []K {
	keys := maps.Keys(m)
	slices.SortFunc[[]K](keys, f)
	return keys
}
