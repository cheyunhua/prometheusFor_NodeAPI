package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	ch := make(chan os.Signal, 1)
	stop := make(chan struct{}, 1)

	signal.Notify(ch, os.Interrupt, os.Kill)
	select {
	case <-ch:
		fmt.Println("re is a data closeding........")
		close(ch)
	case <-stop:
		return

	}

}
