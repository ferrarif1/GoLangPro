package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
channel 用于在协程之间通信
分为两种：有缓存，无缓存
有缓存的，缓存满之前双方都可以收发
无缓存的受方接收前发送方不可再发
*/
var values = make(chan int)

func send() {
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(10)
	fmt.Printf("send value: %v\n", value)
	values <- value
}

func TestChannel() {
	//    unbuffered := make(chan int) //无缓冲
	//    buffered := make(chan int, 10) //有缓冲
	defer close(values)
	go send()
	fmt.Println("wait...")
	value := <-values
	fmt.Printf("value: %v\n", value)
	fmt.Println("end...")

	//如果向一个已经关闭的通道（channel）发送消息，就会导致panic异常的发生。
	//如果从一个已经关闭的通道读取消息，可以读取到已经被发送到通道中但尚未被读取的消息。但是，一旦通道中的所有消息都被读取完毕，再从通道中读取消息就会返回零值，并且通道的状态会变成已关闭。
	ch := make(chan int, 10)
	close(ch)

	select {
	case ch <- 1: //panic: send on closed channel
		// 不会执行到这里
	default:
		fmt.Println("通道已关闭，无法发送消息")
	}

}
