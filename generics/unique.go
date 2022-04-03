package generics

func Unique[T comparable](s []T) []T {
	inResult := make(map[T]bool)
	var result []T
	for _, v := range s {
		if _, ok := inResult[v]; !ok {
			result = append(result, v)
			inResult[v] = true
		}
	}
	return result
}
