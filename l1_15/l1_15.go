package l1_15

import "fmt"

func createHugeString(size int64) string {
	var s []rune = make([]rune, size)
	var i int64
	for i = 0; i < size; i++ {
		s[i] = rune(int64('a') + i%26)
	}
	return string(s)
}

var justString string

func SomeFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100] // The string v will keep alive even after destroying the function
	_ = justString       // because justString is a slice of it.
	fmt.Println(justString)
}

func SomeFuncFix() {
	v := createHugeString(1 << 10)
	tmp := make([]rune, 100)
	copy(tmp, []rune(v[:100]))
	justString = string(tmp)
	_ = justString
	fmt.Println(justString)
}
