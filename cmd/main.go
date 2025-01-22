package main

import (
	"fmt"
	"os"
	"time"

	"github.com/FrenkenFlores/golang_l1/l1_1"
	"github.com/FrenkenFlores/golang_l1/l1_10"
	"github.com/FrenkenFlores/golang_l1/l1_11"
	"github.com/FrenkenFlores/golang_l1/l1_12"
	"github.com/FrenkenFlores/golang_l1/l1_13"
	"github.com/FrenkenFlores/golang_l1/l1_14"
	"github.com/FrenkenFlores/golang_l1/l1_15"
	"github.com/FrenkenFlores/golang_l1/l1_2"
	"github.com/FrenkenFlores/golang_l1/l1_3"
	"github.com/FrenkenFlores/golang_l1/l1_4"
	"github.com/FrenkenFlores/golang_l1/l1_5"
	"github.com/FrenkenFlores/golang_l1/l1_6"
	"github.com/FrenkenFlores/golang_l1/l1_7"
	"github.com/FrenkenFlores/golang_l1/l1_8"
	"github.com/FrenkenFlores/golang_l1/l1_9"
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
	case "7":
		{
			fmt.Println("TEST L1.7")
			smap := l1_7.MutexMap{
				Data: make(map[string]any),
			}
			for i := 0; i < 10; i++ {
				go func() {
					smap.Write(fmt.Sprintf("key_%d", i), i)
				}()
			}
			time.Sleep(time.Second * time.Duration(1))
			for i := 0; i < 10; i++ {
				v, _ := smap.Read(fmt.Sprintf("key_%d", i))
				fmt.Println(v)
			}
		}
	case "8":
		{
			fmt.Println("TEST L1.8")
			var n int64 = 0b11111111111111111111111111111111
			var m int64 = l1_8.SetBit(n, 0, false)
			fmt.Printf("%b\n", m)
		}
	case "9":
		{
			fmt.Println("TEST L1.9")
			var nums []int = []int{1, 2, 3, 4, 5}
			l1_9.Pipeline(nums)
		}
	case "10":
		{
			fmt.Println("TEST L1.10")
			var nums []float32 = []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, 335.5, 333.7}
			group := l1_10.GroupTemperature(nums)
			fmt.Println(group)
		}
	case "11":
		{
			fmt.Println("TEST L1.11")
			var a map[int]bool = map[int]bool{1: true, 2: true, 3: true, 4: true, 5: true}
			var b map[int]bool = map[int]bool{10: true, 8: true, 3: true, 4: true, 5: true}

			intersect := l1_11.Intersect(a, b)
			fmt.Println(intersect)
		}
	case "12":
		{
			fmt.Println("TEST L1.12")
			var a []string = []string{"cat", "dog", "cat", "zebra"}
			set := l1_12.MakeSet(a)
			for k := range set {
				fmt.Println(k)
			}
		}
	case "13":
		{
			fmt.Println("TEST L1.13")
			var a, b int = 0, 1
			fmt.Println(a, b)
			l1_13.Swip(&a, &b)
			fmt.Println(a, b)
		}
	case "14":
		{
			fmt.Println("TEST L1.14")
			fmt.Println(l1_14.FindType(1))
			fmt.Println(l1_14.FindType(true))
			fmt.Println(l1_14.FindType("Hello"))
			fmt.Println(l1_14.FindType(1.0))
			fmt.Println(l1_14.FindType(nil))
		}
	case "15":
		{
			fmt.Println("TEST L1.15")
			l1_15.SomeFunc()
			l1_15.SomeFuncFix()
		}
	}

}
