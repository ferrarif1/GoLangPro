package main

import (
	"fmt"
)

type USBHandler interface {
	read() string
	write2(string)
}

type pointUSBHandler interface {
	read() string
	write(string)//这个write接收指针参数，可更改外部变量
	write2(string)//这个write2接收值参数，不可修改外部变量
}

type Computer struct {
	name string
}

func (com Computer) read() string {
	fmt.Printf("read com.name: %v\n", com.name)
	return com.name
}

func (com *Computer) write(name string) {
	com.name = name
}

func (com Computer) write2(name string) {
	com.name = name
}

func InterfaceTest() {
	// c := new(Computer)
	c := Computer{
		name: "Mac",
	}
	c.write2("Tommy2")
	s := c.read()
	fmt.Printf("s: %v\n", s)
	c.write("Tommy")
	s2 := c.read()
	fmt.Printf("s2: %v\n", s2)

	//interface赋值为结构体类型，无法调用write()
	var cs USBHandler = Computer{
		name: "Windows",
	}
	cs.read()
	cs.write2("Windows update")
	cs.read()

	//interface赋值为指针类型 才能调用以指针为接收者的方法write()，对于read()和write2()，指针会自动解析，所以可以调用这两个方法
	var cs2 pointUSBHandler = &Computer{
		name: "po Windows",
	}
	cs2.write("po Windows update")
	cs2.read()
	cs2.write2("fake update")
	cs2.read()
}
