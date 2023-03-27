package main

import "fmt"
import "sync"

var wg sync.WaitGroup

func showMsgInt(i int){
	defer wg.Done() //基本等价于wp.Add(-1) 结束一个goroutine -1
    fmt.Printf("i: %v\n", i)
}

func TestWaitGroup() {
	for i := 0; i < 10; i++ {
		go showMsgInt(i)
		wg.Add(1)//启动一个goroutine +1
	}//主协程会等待group执行完
	wg.Wait()
	fmt.Println("end of wait")
}