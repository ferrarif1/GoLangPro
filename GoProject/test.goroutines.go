package main

import (
	"fmt"
	"time"
)

/*
go 多线程-协程
*/
func showMsg(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("msg: %v\n", msg)
		time.Sleep(time.Millisecond * 100)
	}
}

func TestRoutine()  {
	go showMsg("test routine") //用go命令开启协程
	showMsg("test routine2")
}