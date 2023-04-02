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

	//控制台输入数据
	// var name string
	// var age string
	// fmt.Println("Please Input Name & age")
	// fmt.Scan(&name, &age)
	// fmt.Printf("name: %v\n", name)
	// fmt.Printf("age: %v\n", age)

	// //slice crud
	//TestSlice()

	// //map
	// TestMap()

	// //func
	// TestFunc()

	// //point array
	// TestPointArr()

	// //struct
	// TestStruct()

	// ProjTest()

	// InterfaceTest()

	// TestNewFunc()

	/*
	   并发编程：start***********
	*/
	// TestRoutine()

	// TestChannel()

	// TestWaitGroup()

	// TestRunTime()

	// TestMutex()

	// TestChannelIterate()

	// TestSelectSwitch()

	// TestTimer()

	//TestTicker()

	// TestAtomic()
	//TestAtomic2()

	/*
	   并发编程：end***********
	   golang标准库os模块：start***********
	*/

	//TestOSFile()
	//TestFile()
	// TestWriteFile()

	// TestIO()
	/*
	   golang标准库os模块：end***********
	*/

	/*
		牛客网题目：
	*/
	TestNewCoder1()

}
