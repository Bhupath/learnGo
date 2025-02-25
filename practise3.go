package main

import (
	"fmt"
	"time"
)

func task1(ch chan string) {
	ch <- "task1 is completed"
}
func task2(ch chan int) {
	ch <- 2
}
func main() {
	ch1 := make(chan string)
	ch2 := make(chan int)
	go task1(ch1)
	go task2(ch2)
	time.Sleep(10 * time.Second)
	select {
	case msg := <-ch1:
		fmt.Println(msg)
	case msg := <-ch2:
		fmt.Print(msg)

	}
}
