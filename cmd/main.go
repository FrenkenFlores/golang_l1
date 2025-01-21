package main

import (
	"fmt"
	"os"

	"github.com/FrenkenFlores/golang_l1/l1_1"
	"github.com/FrenkenFlores/golang_l1/l1_2"
	"github.com/FrenkenFlores/golang_l1/l1_3"
	"github.com/FrenkenFlores/golang_l1/l1_4"
	"github.com/FrenkenFlores/golang_l1/l1_5"
	"github.com/FrenkenFlores/golang_l1/l1_6"
)

func main() {
	switch os.Args[1] {
	case "1":
		{
			fmt.Println("TEST L1.1")
			a := l1_1.Action{Human: l1_1.Human{Name: "Adam"}}
			fmt.Println(a.GetName())
		}
	case "2":
		{
			fmt.Println("TEST L1.2")
			a := l1_2.GetSquareList([]int{2, 4, 6, 8, 10})
			fmt.Println(a)
		}
	case "3":
		{
			fmt.Println("TEST L1.3")
			a := l1_3.GetSquareSum([]int{2, 4, 6, 8, 10})
			fmt.Println(a)
		}
	case "4":
		{
			fmt.Println("TEST L1.4")
			var workers_n int
			fmt.Printf("Enter the number of workers: ")
			fmt.Scan(&workers_n)
			l1_4.ReadWriteThread(workers_n)
		}
	case "5":
		{
			fmt.Println("TEST L1.5")
			var n int = 5
			l1_5.WriteAndReadChan(n)
		}
	case "6":
		{
			fmt.Println("TEST L1.6")
			fmt.Println("Stop by channel")
			l1_6.StopByChan()
			fmt.Println("Stop by context")
			l1_6.StopByContext()
			fmt.Println("Stop by timer")
			l1_6.StopByTimer()
			fmt.Println("Stop by Goexit")
			l1_6.StopByGoexit()
		}
	}

}
