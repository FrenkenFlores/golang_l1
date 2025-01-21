package l1_10

func GroupTemperature(temperatures []float32) map[int]map[float32]bool {
	var group map[int]map[float32]bool = make(map[int]map[float32]bool)
	for i := 0; i < len(temperatures); i++ {
		_, ok := group[int(temperatures[i])/10*10]
		if !ok {
			group[int(temperatures[i])/10*10] = map[float32]bool{temperatures[i]: true}
		} else {
			group[int(temperatures[i])/10*10][temperatures[i]] = true
		}
	}
	return group
}
