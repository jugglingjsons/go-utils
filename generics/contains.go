package generics

func Contains[T comparable](elems []T, el T) bool {
	for _, v := range elems {
		if v == el {
			return true
		}
	}
	return false
}
