package main

import (
	"fmt"
	"strings"
)

// 普通函数
func sum(a int, b int) (ret int) {
	ret = a + b
	return ret
}

// 匿名函数
func anoymous() {
	max := func(a int, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	i := max(1, 2)
	fmt.Printf("i: %v\n", i)

}

//结构体函数

// 函数类型
func FuncType() {
	type f1 func(int, int) int
	var ff f1
	ff = sum
	ret := ff(1, 4)
	fmt.Printf("ret: %v\n", ret)
}

// 高阶函数
func say(name string) {
	fmt.Printf("hello, %s\n", name)
}

func test(name string, f func(string)) {
	f(name)
}

// 闭包：定义在一个函数内部的函数 = 函数+引用环境
// 返回一个函数
func testPKG() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name+suffix
		}
		return name
	}
}

func TestFunc() {
	ret := sum(2, 5)
	fmt.Printf("ret: %v\n", ret)
	FuncType()
	test("tom", say)
	anoymous()

	var f = testPKG()
	fmt.Printf("f: %v\n", f(10))
	fmt.Printf("f: %v\n", f(20))
	fmt.Printf("f: %v\n", f(30))
	f1 := testPKG() //此时新的x为0
	fmt.Printf("f1: %v\n", f1(40))
	fmt.Printf("f1: %v\n", f1(50))
    
	jpgFunc := makeSuffixFunc(".jpg")
	txtFunc := makeSuffixFunc(".txt")
    fmt.Println(jpgFunc("test"))
	fmt.Println(txtFunc("test"))
}
