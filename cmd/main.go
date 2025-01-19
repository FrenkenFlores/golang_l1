package main

import (
	"fmt"

	"github.com/FrenkenFlores/golang_l1/l1_1"
	"github.com/FrenkenFlores/golang_l1/l1_2"
)

func main() {
	{
		fmt.Println("TEST L1.1")
		a := l1_1.Action{Human: l1_1.Human{Name: "Adam"}}
		fmt.Println(a.GetName())
	}
	{
		fmt.Println("TEST L1.2")
		a := l1_2.GetSquareList([]int{2, 4, 6, 8, 10})
		fmt.Println(a)
	}

}
