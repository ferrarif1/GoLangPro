package main

import (
	"fmt"
	"sync"
	"time"
)

/*
除了用channel外 可用mutex实现线程间同步

通过lock，其他线程就无法操作资源
*/

var m int = 100

var lock sync.Mutex

var wt sync.WaitGroup

func add() {
	defer wt.Done()
	lock.Lock()
	m += 1
    time.Sleep(time.Millisecond*10)
	lock.Unlock()
}

func sub()  {
	defer wt.Done()
	lock.Lock()
	m-=1
	time.Sleep(time.Millisecond*2)
	lock.Unlock()
}

func TestMutex() {
    for i := 0; i < 100; i++ {
		go add()
		wt.Add(1)
		go sub()
		wt.Add(1)
	}
	wt.Wait()
	fmt.Printf("m: %v\n", m)
}
