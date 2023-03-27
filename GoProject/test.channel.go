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

func send()  {
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
	 value := <- values
	 fmt.Printf("value: %v\n", value)
     fmt.Println("end...")
}
