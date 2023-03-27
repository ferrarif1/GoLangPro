package main

import (
	"fmt"
	// "sync"
	"sync/atomic"
	"time"
)

var i int32 = 100

//方法1:加锁
//加锁后就不会出错，否则可能是99、101...
// var locki sync.Mutex

// func addI() {
// 	locki.Lock()
// 	i++
// 	locki.Unlock()
// }

// func subI() {
// 	locki.Lock()
// 	i--
// 	locki.Unlock()
// }

//方法2:atomic
func addI2() {
	atomic.AddInt32(&i, 1)
}

func subI2() {
	atomic.AddInt32(&i, -1)
}

func TestAtomic() {
	for i := 0; i < 100; i++ {
		// go addI()
		// go subI()
        go addI2()
		go subI2()
	}
	time.Sleep(time.Second * 2)
	fmt.Printf("i: %v\n", i)
}
