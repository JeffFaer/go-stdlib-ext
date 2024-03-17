// moremaps impements extensions to the standard maps library.
package moremaps

import (
	"cmp"
	"slices"

	"golang.org/x/exp/maps"
)

type SortableKey interface {
	comparable
	cmp.Ordered
}

func SortedKeys[K SortableKey, V any](m map[K]V) []K {
	return SortedKeysFunc(m, cmp.Compare[K])
}

func SortedKeysFunc[K comparable, V any](m map[K]V, f func(K, K) int) []K {
	keys := maps.Keys(m)
	slices.SortFunc[[]K](keys, f)
	return keys
}
