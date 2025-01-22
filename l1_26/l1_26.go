package l1_26

import "strings"

func CheckUnique(s string) bool {
	s = strings.ToLower(s)
	m := make(map[rune]bool)
	for _, v := range s {
		if m[v] {
			return false
		}
		m[v] = true

	}
	return true
}
