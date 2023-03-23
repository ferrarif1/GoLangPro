package main

import "fmt"

// init()是自动执行的，不能被其他函数调用，执行顺序：变量初始化-》init（）-》main（）
func init() {
	fmt.Println("init")
}
