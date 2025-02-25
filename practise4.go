package main

import (
	"fmt"
	"time"
)

func first() {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}
func second() {
	for j := 10; j < 20; j++ {
		time.Sleep(2 * time.Second)
		fmt.Println(j)
	}
}
func main() {
	go first()
	go second()
	time.Sleep(10 * time.Second)
	fmt.Println("Go concurrency finished")
	var test chan string
}
