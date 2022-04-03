package generics

func Map[T any, M any](s []T, f func(T) M) []M {
	newSlice := make([]M, len(s))
	for i, v := range s {
		newSlice[i] = f(v)
	}
	return newSlice
}
