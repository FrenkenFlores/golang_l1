package l1_6

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func StopByChan() {
	stopChan := make(chan bool)
	go func() {
		for loop := true; loop; {
			select {
			case <-stopChan:
				loop = false
			default:
				fmt.Println("Working")
				time.Sleep(time.Second * time.Duration(1))
			}
		}
	}()
	time.Sleep(time.Second * time.Duration(3))
	stopChan <- true
	close(stopChan)
}

func StopByContext() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		for loop := true; loop; {
			select {
			case <-ctx.Done():
				loop = false
			default:
				fmt.Println("Working")
				time.Sleep(time.Second * time.Duration(1))
			}
		}
	}()
	time.Sleep(time.Second * time.Duration(3))
	cancel()
}

func StopByTimer() {
	timer := time.NewTimer(time.Second * time.Duration(3))
	go func() {
		for loop := true; loop; {
			select {
			case <-timer.C:
				loop = false
			default:
				fmt.Println("Working")
				time.Sleep(time.Second * time.Duration(1))

			}
		}
	}()
	time.Sleep(time.Second * time.Duration(4))
}

func StopByGoexit() {
	go func() {
		for {
			fmt.Println("Working")
			time.Sleep(time.Second * time.Duration(3))
			runtime.Goexit()
		}
	}()
	time.Sleep(time.Second * time.Duration(4))
}
