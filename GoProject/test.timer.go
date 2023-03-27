package main

import (
	"fmt"
	"time"
)

func TestTimer() {
	//以下几种方法都可以等待2s
	//1
	timer := time.NewTimer(time.Second * 2)
	fmt.Printf("time.Now(): %v\n", time.Now()) //2023-03-27 18:59:31
	t1 := <-timer.C                            //阻塞time.Second*2时长
	fmt.Printf("t1: %v\n", t1)                 //2023-03-27 18:59:33 差两秒

	//2
	timer2 := time.NewTimer(time.Second * 2)
	<-timer2.C
	fmt.Printf("time.Now(): %v\n", time.Now())

	//3
	time.Sleep(time.Second * 2)

	//4
	<-time.After(time.Second * 2)

	/*
		执行结果：
			func timer3...
			main end 1
	*/
	//5
	timer3 := time.NewTimer(time.Second * 2)
	go func() {
		<-timer3.C
		fmt.Println("func timer3...")
	}()
	time.Sleep(time.Second * 3) //主线程停3s，这样才能func才能先执行完
	fmt.Println("main end 1")

	/*
		执行结果：直接停止了计时，不会执行func
			Timer4 stopped
			main end 2
	*/
	timer4 := time.NewTimer(time.Second * 2)
	go func() {
		<-timer4.C
		fmt.Println("func timer4...")
	}()

	stop := timer4.Stop() //停止定时器
	if stop {             //阻止timer事件发生，当函数执行后，timer计时器停止，相应的事件不再执行
		fmt.Println("Timer4 stopped")
	}
	fmt.Println("main end 2")

}
