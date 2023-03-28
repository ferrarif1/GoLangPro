package main

import "fmt"

func Add() {
	s := []int{}
	s = append(s, 100)
}

func Del() {
	var s1 = []int{1, 2, 3, 4, 5}
	//删除第2个元素
	s1 = append(s1[:2], s1[3:]...)
	fmt.Printf("s1: %v\n", s1)
}

func Update() {
	var s1 = []int{1, 2, 3, 4, 5}
	s1[2] = 100
	fmt.Printf("s1: %v\n", s1)
}

func Update2() {
	var s1 = []int{1, 2, 3, 4, 5}
	s2 := s1
	var s3 = make([]int, 5) //注意用var s3 = []int{}将无法拷贝，因为没有分配内存空间
	copy(s3, s1)
	s2[2] = 66
	fmt.Printf("s1: %v\t %p\n", s1, s1)
	fmt.Printf("s2: %v\t %p\n", s2, s2)
	fmt.Printf("s3: %v\t %p\n", s3, s3)
}

// 数组与切片转换
/*
1.在 Golang 中，b 是一个切片（slice），b[:] 是对切片 b 的一个切片操作，它表示获取 b 中的所有元素，所以 b 和 b[:] 在大多数情况下是等价的。
2. 不过，如果 b 是一个数组，则 b 和 b[:] 有一些不同之处。数组是一个固定长度的序列，而切片是对数组的一个引用，它可以动态调整大小。
当对一个数组进行切片操作时，得到的结果是一个切片，而不是一个数组。因此，b 和 b[:] 在这种情况下是不同的。

通过 a[:] 来获取一个切片 b，这个切片引用了整个数组 a。我们还通过将数组 a 赋值给变量 c 来创建了另一个数组。
b 的第一个元素被设置为了 0。这是因为 b 是一个对 a 的切片，对 b 的修改会反映到原始数组 a 上，而 c 是一个独立的数组，对 c 的修改不会影响 a。
*/
func convert() {
	a := [3]int{1, 2, 3}
	b := a[:]
	c := a

	b[0] = 0
	c[0] = 10

	fmt.Println("a:", a) // [0 2 3]
	fmt.Println("b:", b) // [0 2 3]
	fmt.Println("c:", c) // [10 2 3]
}

func TestSlice() {
	Add()
	Del()
	Update()
	Update2()
	convert()
}
