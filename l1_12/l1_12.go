package l1_12

func MakeSet(strs []string) map[string]bool {
	set := make(map[string]bool)
	for _, str := range strs {
		set[str] = true
	}
	return set
}
