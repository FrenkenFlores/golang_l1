package l1_11

func Intersect(a, b map[int]bool) map[int]bool {
	res := make(map[int]bool)
	if len(a) > len(b) {
		a, b = b, a
	}
	for k, _ := range a {
		if b[k] {
			res[k] = true
		}
	}
	return res
}
