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

func Update()  {
	var s1 = []int{1, 2, 3, 4, 5}
	s1[2] = 100
	fmt.Printf("s1: %v\n", s1)
}