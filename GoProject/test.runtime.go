package main

import (
	"fmt"
	"runtime"
	"time"
)

func a()  {
	for i := 0; i < 10; i++ {
		fmt.Printf("a_i: %v\n", i)
	}
}

func b()  {
	for i := 0; i < 10; i++ {
		fmt.Printf("b_i: %v\n", i)
	}
}
func TestRunTime() {
	fmt.Printf("runtime.NumCPU(): %v\n", runtime.NumCPU())
	runtime.GOMAXPROCS(2)//设定最高2核 

	go a()
	go b()

	time.Sleep(time.Second)
}