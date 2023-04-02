package main

import (
	// "flag"
	"fmt"
)

//构造函数

func NewPerson(name string, age int) (*Person, error) {
	if name == "" {
		return nil, fmt.Errorf("name cannot be empty")
	}
	if age < 0 {
		return nil, fmt.Errorf("age cannot be negative")
	}
	return &Person{name: name, age: age}, nil
}

/*
make和new都是golang用来分配内存的內建函数，且在堆上分配内存，make 即分配内存，也初始化内存。new只是将内存清零，并没有初始化内存。
make返回的还是引用类型本身；而new返回的是指向类型的指针。
make只能用来分配及初始化类型为slice，map，channel的数据；new可以分配任意类型的数据。
*/
func testNew() {
	fmt.Println("testNew start********************************")
	var v *int
	///fmt.Printf("(*v): %v\n", (*v)) //runtime error: invalid memory address or nil pointer dereference
	fmt.Printf("v: %v\n", v) //<nil>
	v = new(int)
	fmt.Printf("(*v) after new(): %v\n", (*v)) // 0
	fmt.Printf("v2 after new(): %v\n", v)      //0xc00012c030

}

func TestNewFunc() {
	pe, err := NewPerson("John", 28)
	if err == nil {
		fmt.Printf("pe.name: %v\n", pe.name)
	} else {
		fmt.Printf("err: %v\n", err)
	}

	testNew()
}
