// package main

// import (
// 	"fmt"
// )

// var NAME = "BHUPATH KUMAR"

// func main() {
// 	fmt.Printf("%v\n", NAME)
// 	var name = "BK"
// 	fmt.Println(name)
// 	var c_test = complex(1, 2)
// 	fmt.Println(real(c_test))
// 	fmt.Println(imag(c_test))
// 	fmt.Println(name == "BK")
// 	fmt.Print(name[1])
// 	const NAME = "Bhupath"
// 	fmt.Print(NAME)
// }

// import (
// 	"fmt"
// 	"os"
// 	"strings"
// )

//	func main() {
//		fmt.Print("enter any number\n")
//		var number int
//		fmt.Scanln(&number)
//		fmt.Printf("Entered Number is : %d", number)
//	}
// func justreturn() int {
// 	return 777
// }
// func testing() {
// 	array := [10]int{}
// 	fmt.Printf("Array: %d\n", array[0:9])
// 	for i, j := range array {
// 		fmt.Println(i, j)
// 	}
// 	fmt.Println(len(array))
// 	fmt.Print(os.Args[1:])
// 	temp_string := "Bhupath Kumar"
// 	fmt.Println(strings.HasPrefix(temp_string, "B"))
// 	fmt.Println(strings.HasSuffix(temp_string, "r"))
// 	temp := map[int]int{1: 123}
// 	i, ok := temp[1]
// 	fmt.Print(i, ok)

// }
// func main() {
// 	output := justreturn()
// 	fmt.Printf("Returned value: %d\n", output)
// 	testing()
// }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func display(msg string) {
// 	for i := 0; i <= 3; i++ {
// 		fmt.Printf("%v %v\n", msg, i)
// 	}
// }
// func main() {
// 	fmt.Println("Hello World")
// 	go display("Bhupath1")
// 	display("Bhupath")
// 	time.Sleep(10 * time.Second)
// 	go func(s string) {
// 		for i := 0; i <= 2; i++ {
// 			fmt.Printf("%v %v\n", i, s)
// 			time.Sleep(time.Millisecond)
// 		}
// 	}("ME GO ROUTINE")
// 	time.Sleep(10 * time.Second)
// 	fmt.Println("concurency func completed.")
// }

// Pointers
package main

import (
	"fmt"
)

var N_NAME string = "Bhupath1"

const (
	NAME = "Bhupath"
)

func main() {
	defer count1(111111111)
	defer count1(111111111223)
	fmt.Println("Hello World..GO")
	// NAME := "Test"
	fmt.Print(N_NAME)
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	if NAME == "Bhupath" && N_NAME == "Bhupath1" {
		fmt.Print(true)
	} else {
		fmt.Print(false)
	}
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
	temp := []int{1, 2, 3, 3}
	for _, j := range temp {
		fmt.Println(j)
	}
	temp1 := map[int]string{
		1: "first",
	}
	fmt.Print(temp1)
	mmap := map[int]string{
		22: "Geeks",
		33: "GFG",
		44: "GeeksforGeeks",
	}
	for key, value := range mmap {
		fmt.Println(key, value)
	}
	chn1 := make(chan int)
	go func() {
		chn1 <- 10
		chn1 <- 1
		close(chn1)
	}()
	for i := range chn1 {
		fmt.Println(i)
	}
	switch day1 := "Sunday"; day1 {
	case "Sunday":
		fmt.Println("It's Sunday")
	case "0":
		fmt.Println("It;s ZERO")
	default:
		fmt.Print("Am Default")
	}
	one := 2
	two := 3
	fmt.Println(sumOfTwoNumbers(one, two))
	fmt.Println(multiply(&two, &one))
	fmt.Println("How many arguments are passed", count(1, 2, 3))

	fmt.Println("How many arguments are passed", count(1, 2))
	fmt.Println("How many arguments are passed", count(1, 2, 3, 3, 4, 5, 6))
	defer count1(222222222)
	//defer function executed LIFO

}

// call by value
func sumOfTwoNumbers(a int, b int) int {
	return a + b
}

// call by reference
func multiply(a *int, b *int) int {
	*a = *a + 1
	return *a * *b
}

// variadic function
func count(nums ...int) int {
	temp := 0
	for i := range nums {
		temp = i
	}
	return temp + 1
	// Limitation: Can have variadic parameter at only one & if more then parameter were there to place at last
}
func count1(nums ...int) {

	for _, j := range nums {
		fmt.Println(j)
	}
	// Limitation: Can have variadic parameter at only one & if more then parameter were there to place at last
}

// init function which would be default intiated at runtime before main function calls
// Can have n no of init functions in a program
func init() {
	fmt.Print("Helo inint")
}
func init() {
	fmt.Println("hello init")
}

//anonymous function can be used inside of the function. Not outside
//Syntax func(*args){
// }(*args) or func(){
// var_ := func(){}return var_}
