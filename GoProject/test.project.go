package main

import "fmt"

type Integer int
func (a Integer)Add(b Integer) Integer {
	return a+b
}

func ProjTest()  {
	var a Integer = 1
	var b Integer = 2
	var i interface{} = a
	sum := i.(Integer).Add(b)
	fmt.Printf("sum: %v\n", sum)
}