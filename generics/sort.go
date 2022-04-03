package generics

import (
	"sort"

	"golang.org/x/exp/constraints"
)

func SortSlice[T constraints.Ordered](s []T) {
	sort.SliceStable(s, func(i, j int) bool {
		return s[i] < s[j]
	})
}
