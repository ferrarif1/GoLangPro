package main

import (
	"fmt"
)

type Person struct {
	id   int
	name string
}

// 给结构体定义方法
func (recv Person) createSon(id int, name string) Person {
	recv.id = id
	recv.name = name
	return recv
}

func (recv *Person) update(id int, name string) Person {
	(*recv).id = id
	(*recv).name = name
	return *recv
}

func TestStruct() {
	//类型定义
	type MyInt int
	var i MyInt
	i = 100
	fmt.Printf("i: %T\n", i) //main.MyInt

	//类型别名
	type MyInt2 = int
	var i2 MyInt2
	i2 = 100
	fmt.Printf("i: %T\n", i2) //int

	var tom Person
	tom.name = "tom"

	kite := Person{id: 1, name: "kite"}
	fmt.Printf("kite: %v\n", kite)

	//匿名结构体
	var jerry struct {
		id   int
		name string
	}
	jerry.id = 132
	jerry.name = "jerry"

	sec := tom.createSon(7, "sec")
	fmt.Printf("tom1: %v\n", tom)
	fmt.Printf("sec: %v\n", sec)
	
	newtom := tom.update(4, "newtom")
	fmt.Printf("tom2: %v\n", tom)
	fmt.Printf("newtom: %v\n", newtom)

}
