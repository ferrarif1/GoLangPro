package Utils

import (
	"fmt"
	"unsafe" //一般不用 不安全
)

func Defaultvalue(){
	fmt.Printf("default value\n")
}

func Scope(){
	var a int
	var pointer unsafe.Pointer = unsafe.Pointer(&a)
	var p uintptr = uintptr(pointer)
	var ptr *int = &a
	fmt.Printf("pointer:%p \n p:%d \n p:%x \n ptr:%p\n",pointer,p,p,ptr)
    
}