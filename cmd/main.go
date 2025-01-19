package main

import (
	"fmt"

	"github.com/FrenkenFlores/golang_l1/l1_1"
)

func main() {
	{
		fmt.Println("TEST L1.1")
		a := l1_1.Action{Human: l1_1.Human{Name: "Adam"}}
		fmt.Println(a.GetName())
	}

}
