package main

import "fmt"

func Judge(aim int, compare int, max int) bool {
	sub := aim - compare
	if sub < 0 {
		sub = -sub
	}
	return sub > max
}

func Test(n int, x int, a []int) int {
	max := 0
	min := 0
	for _, v := range a {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	if Judge(max, min, x) {
		return 0
	}
	num := 0
	initNum := (max + min) / 2
	for _, v := range a {
		if Judge(initNum, v, x) {
			num++
		}
	}
	return num
}

func TestShen() {
	var n int
	var x int
	var a []int
	fmt.Scan(&n, &x)
	if n > 0 {
		a = make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&a[i])
		}
	}
	Test(n, x, a)

}
