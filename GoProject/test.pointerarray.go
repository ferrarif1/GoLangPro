package main

import "fmt"

func TestPointArr() {
	a := [3]int{1, 2, 3}
	var pa [3]*int

	fmt.Printf("pa: %v\n", pa)
	for i := 0; i < len(a); i++ {
		pa[i] = &a[i]
	}
	fmt.Printf("pa: %v\n", pa)

	for i := 0; i < len(pa); i++ {
		fmt.Printf("%v ",*pa[i])
	}
}
