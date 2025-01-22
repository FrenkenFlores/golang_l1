package l1_25

import "time"

func Sleep(x int) {
	ch := time.After(time.Second * time.Duration(x))
	<-ch
}
