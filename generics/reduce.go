package generics

func Reduce[T, M any](s []T, f func(M, T) M, accStart M) M {
	acc := accStart
	for _, v := range s {
		acc = f(acc, v)
	}
	return acc
}
