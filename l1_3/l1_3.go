package l1_3

import "sync"

func getSquare(a int, wg *sync.WaitGroup, ch chan<- int) {
	defer wg.Done()
	ch <- a * a
}

func GetSquareSum(a []int) int {
	var wg sync.WaitGroup
	ch := make(chan int, len(a))
	for _, i := range a {
		wg.Add(1)
		go getSquare(i, &wg, ch)
	}
	wg.Wait()
	var res int = 0
	for i := 0; i < len(a); i++ {
		res += <-ch
	}

	return res
}
