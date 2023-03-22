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

func TestSlice()  {
	Add()
	Del()
	Update()
	Update2()
}