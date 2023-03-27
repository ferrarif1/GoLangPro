package main

import (
	"fmt"
	"time"
)

/*
select是go中一个控制结构，用于异步处理IO操作
select会监听channel的读写操作，当case中channel读写操作为非阻塞状态（能读能写）时，将会触发相应的动作
* select中的case语句必须是一个channel操作
* select中的default总是可以运行的

如果多个case都可以运行，会随机公平选择一个执行
如果没有可以执行的case语句，且有default，会执行default
如果没有可运行的case语句，且没有default，select将阻塞，直到某个case通信可以运行
*/

var chanInt = make(chan int)
var chanString = make(chan string)

func TestSelectSwitch() {
	go func() {
		chanInt <- 100
		chanString <- "hello"
		defer close(chanInt) //只注释掉两个close后，会执行default
		defer close(chanString)
	}()

	for {
		select {
		case r := <-chanInt:
			fmt.Printf("chanInt: %v\n", r)
		case r := <-chanString:
			fmt.Printf("chanString: %v\n", r)
		default: //注释掉close和default后，将会死锁：fatal error: all goroutines are asleep - deadlock!
			fmt.Println("default...")
		}
		time.Sleep(time.Second)
	}
}
