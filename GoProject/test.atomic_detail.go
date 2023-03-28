package main

import (
	"fmt"
	"sync/atomic"
)

func TestAtomic2() {
	//read
	var i int32 = 100
	atomic.LoadInt32(&i)
	fmt.Printf("i: %v\n", i)

    //write
	atomic.StoreInt32(&i,102)
	fmt.Printf("i: %v\n", i)

	//cas
	b := atomic.CompareAndSwapInt32(&i,102,200) //把200赋值给i，如果这个过程中有其他线程过来把i改成其他值使其不等于102，则修改失败 不进行交换
    fmt.Printf("b: %v\n", b)//修改成功返回true
	fmt.Printf("i: %v\n", i)
}