package main

import (
	"fmt"
	"time"
)

/*
time只执行一次，ticker可周期执行
*/
func TestTicker() {
	//周期执行
	ticker := time.NewTicker(time.Second)
	counter := 0
	for range ticker.C { //每秒执行一次
		fmt.Println("ticker...")
		counter++
		if counter > 3 {
			ticker.Stop()
			break
		}
	}

	//周期发送给chanint
	ticker2 := time.NewTicker(time.Second)
	chanInt := make(chan int)
	
	go func ()  {//启动一个协程往channel里放元素
		for range ticker2.C {
			select {
			case chanInt <- 1:
			case chanInt <- 2:
			case chanInt <- 3:
			}
		}
	}()
	sum := 0
	for v := range chanInt {//在主协程一直循环等着收，查到就输出
		fmt.Printf("receive v: %v\n", v)
        sum+=v
		if sum>10 {
			break;
		}
	}
}
