package main

import "fmt"

func TestChannelIterate() {
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	//遍历方法1
	for {
		data, ok := <-c
		if ok {
			fmt.Printf("data: %v\n", data)
		} else {
			break
		}
	}

	//如果通道关闭，读多写少，没有了就是默认值，如果没有关闭，就会死锁
	//遍历方法2
	for v := range c {
		fmt.Printf("v: %v\n", v)
	}

}
