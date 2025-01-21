package l1_8

func SetBit(x int64, i int, v bool) int64 {
	if !v {
		return x &^ (1 << i)
	}
	return x | (1 << i)
}
