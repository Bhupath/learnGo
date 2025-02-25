package main

import "fmt"

func cls_test(number int) func(string) string {
	temp := func(msg string) string {
		return "hello World" + msg
	}
	return temp
}
func main() {
	a := cls_test(10)
	fmt.Println(a("Bhupath"))
}
