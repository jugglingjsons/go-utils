package slices

func Filter(ss []string, f func(string) bool) []string {
	vs := make([]string, 0)
	for _, v := range ss {
		if f(v) {
			vs = append(vs, v)
		}
	}
	return vs
}

func RemoveDuplicates(ss []string) []string {
	vs := make([]string, 0)
	for _, v := range ss {
		if !Contains(vs, v) {
			vs = append(vs, v)
		}
	}
	return vs
}
