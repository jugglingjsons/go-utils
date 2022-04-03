package main

import (
	"fmt"
	"jugglingjsons/go-utils/generics"
	"math"
	"strings"
)

func main() {
	testData := []int{1, 2, 3, 4, 5}
	min := generics.Min(testData)
	max := generics.Max(testData)
	fmt.Println(min, max)

	floatSlice := []float64{8.9, 1.1, 2.2, 3.3, 4.4, 5.5}
	generics.SortSlice(floatSlice)
	fmt.Println(floatSlice)

	stringSlice := []string{"a", "b", "c", "d", "e", "f", "a", "a", "a", "a", "a"}
	uniqueStringSlice := generics.Unique(stringSlice)
	fmt.Println(uniqueStringSlice)

	intSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum := generics.Reduce(intSlice, func(acc, current int) int {
		return acc + current
	}, 0)

	fmt.Println(sum)
	divided := generics.Reduce(intSlice, func(acc float64, current int) float64 {
		return acc + float64(current)/10
	}, 0.0)
	fmt.Println(divided)

	sliceFloat := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	newSliceFloat := generics.Map(sliceFloat, math.Sqrt)
	fmt.Println(newSliceFloat)

	toPower2 := generics.Map(sliceFloat, func(v float64) float64 {
		return math.Pow(v, 2)
	})
	fmt.Println(toPower2)

	filtered := generics.Filter(sliceFloat, func(v float64) bool {
		return v > 4
	})
	fmt.Println(filtered)
	websites := []string{"http://foo.com", "https://bar.com", "https://gosamples.dev"}
	filteredWebsites := generics.Filter(websites, func(s string) bool {
		return strings.HasPrefix(s, "https://")
	})
	fmt.Println(filteredWebsites)
}
