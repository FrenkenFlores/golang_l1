package l1_4

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func sigHandler() {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT)
	<-sigCh
	os.Exit(0)
}

func outPutThreads(inp chan string) {
	for {
		fmt.Println(<-inp)
	}
}

func ReadWriteThread(workers int) {
	go sigHandler()
	inp := make(chan string)
	for i := 0; i < workers; i++ {
		go outPutThreads(inp)
	}
	var s string
	for {
		fmt.Scanln(&s)
		inp <- s
	}
}
