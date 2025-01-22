package l1_16

func QuickSort(a []int) {
	if len(a) < 2 {
		return
	}

	left, right := 0, len(a)-1
	pivot := right
	for i := 0; i < len(a); i++ {
		if a[i] < a[pivot] {
			a[i], a[left] = a[left], a[i]
			left++
		}
	}
	a[left], a[pivot] = a[pivot], a[left]
	QuickSort(a[:left])
	QuickSort(a[left+1:])
}
