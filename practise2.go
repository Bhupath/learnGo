package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Hello")
	a := person{name: "Bhupath", age: 28}
	a.display()
	fmt.Println(a.name)
	b := number(10)
	b.display1()
	data := Person{
		name: "Bhupath",
		//age:      28,
		Location: "Bangalore",
	}
	fmt.Println(data)
	temp := map[string]int{
		"Bhupath": 1,
	}
	fmt.Println(temp["Bhupath"])
	arr := []int{1, 2, 3, 3}
	arr[0] = 10
	arr[1] = 20
	arr = append(arr, 12)
	fmt.Println(arr)
	myslice := make([]int, 4)
	//make(type, len, capacity i.e max)
	fmt.Println(cap(myslice))
	fmt.Println(strconv.Atoi("33"))
}

// Create a method with struct receiver
func (p person) display() {
	fmt.Println(p.name)
	fmt.Println(p.age)
}

type person struct {
	name string
	age  int
}
type Person struct {
	name     string
	age      int
	Location string
}

// create a method nonstruct receiver
type number int

func (n number) display1() {
	fmt.Println(n)
}

// Supports nestes structs
type Address struct {
	addr1 string
	addr2 string
	addr3 addr2
	//can access like addr3.location
}
type addr2 struct {
	location string
	hno      string
}
