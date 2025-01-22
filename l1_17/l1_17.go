package l1_17

// 1 2 3 4 5 6 7 8 9 10

func BinarySearch(a []int, b int) (int, bool) {
	if len(a) == 1 && a[0] != b {
		return -1, false
	}
	mid := len(a) / 2
	if a[mid] == b {
		return a[mid], true
	} else if b > a[mid] {
		return BinarySearch(a[mid:], b)
	} else {
		return BinarySearch(a[:mid+1], b)
	}
}
