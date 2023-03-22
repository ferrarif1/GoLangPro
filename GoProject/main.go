package main

import (
	"GoProject/user"
	"fmt"
)

func main() {
	user.Hello()
	a := 200
	p := &a
	fmt.Printf("p: %v\n", p)
	TestString()
	// var name string
	// var age string
	// fmt.Println("Please Input Name & age")
	// fmt.Scan(&name, &age)
	// fmt.Printf("name: %v\n", name)
	// fmt.Printf("age: %v\n", age)

	Add()
	Del()
	Update()
}
