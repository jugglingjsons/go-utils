package slices

func Contains(ss []string, el string) bool {
	for _, v := range ss {
		if v == el {
			return true
		}
	}
	return false
}
