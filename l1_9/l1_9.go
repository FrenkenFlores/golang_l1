package l1_9

import "fmt"

func gen(nums []int) chan int {
	out := make(chan int)
	go func() {
		for _, num := range nums {
			out <- num
		}
		close(out)
	}()
	return out

}

func sq(ch chan int) chan int {
	out := make(chan int)
	go func() {
		for num := range ch {
			out <- num * num
		}
		close(out)
	}()
	return out

}

func Pipeline(nums []int) {
	nums_gen := gen(nums)
	nums_sq := sq(nums_gen)
	for num := range nums_sq {
		fmt.Println(num)
	}
}
