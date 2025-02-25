package main

import "fmt"

type Number interface {
	int | int8 | int16 | int32 | int64
}

func sumOfNumbers[T Number](number []T) T {
	var sum T
	for i := range number {
		sum += number[i]
	}
	return sum
}
func main() {
	num1 := []int8{1, 2, 3}
	num2 := []int16{4, 5, 6}
	num3 := []int32{7, 8, 9}
	num4 := []int64{10}
	fmt.Println(sumOfNumbers(num1))
	fmt.Println(sumOfNumbers(num3))
	fmt.Println(sumOfNumbers(num2))
	fmt.Println(sumOfNumbers(num4))
}
