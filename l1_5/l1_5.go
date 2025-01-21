package l1_5

import "time"

func WriteAndReadChan(n int) {
	ch := make(chan string)
	go func() {
		for {
			ch <- "Hello, world!"
			time.Sleep(time.Second * time.Duration(1))
		}
	}()
	go func() {
		for {
			s := <-ch
			println(s)
		}
	}()
	time.Sleep(time.Second * time.Duration(n))
}
