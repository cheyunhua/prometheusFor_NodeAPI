package main

import (
	"fmt"
	"sync"
)

func main() {
	var ch1, ch2, ch3 = make(chan struct{}), make(chan struct{}), make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(3)
	go func(s string) {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			<-ch1
			fmt.Print(s)
			ch2 <- struct{}{}
		}
		<-ch1
	}("A")
	go func(s string) {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			<-ch2
			fmt.Print(s)
			ch3 <- struct{}{}
		}
	}("B")
	go func(s string) {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			<-ch3
			fmt.Println(s)
			ch1 <- struct{}{}
		}
	}("C")
	ch1 <- struct{}{}
	wg.Wait()
}
