package main

import(
	"time"
	"fmt"
)

type User struct{
	id int
	Score float32
	name, addr string
	enrollment time.Time
}
//匿名结构体
var stu struct{
	id int
	Score float32
	name, addr string
	enrollment time.Time
}

func (u User)hello(man string) {
    fmt.Printf("hi"+man+", my name is "+ u.name)
}

func(_ User)think(man string) {
    fmt.Printf("hi"+man)
}

func update_slice1(arr *[]User)  {
	(*arr)[0].name = "Not shit1"
}


func update_slice2(arr []User)  {
	arr[1].name = "Not shit2"
}

func update_slice3(arr []*User)  {
	arr[2].name = "Not shit3"
}

func Main2(){
	var ws User
	ws = User{Score:100, name:"zjsi"}
	ws.Score = 123
	ws.enrollment = time.Now()
    ws.think("somename")
	a:= ws.id + 24
	fmt.Printf("a = %d\n", a)

	s:=User{Score:100, name:"Shit"}
	arr := []User{s,s,s,s}//深拷贝 不影响原来的
	arr2 := []*User{&s,&s,&s,&s}
	update_slice1(&arr)
	fmt.Println("Sname = "+arr[0].name)//改了
	fmt.Println("Sname = "+s.name)//没改
	update_slice2(arr)
	fmt.Println("Sname = "+arr[1].name)//改了
	fmt.Println("Sname = "+s.name)//没改
	update_slice3(arr2)
	fmt.Println("Sname = "+arr2[2].name)//改了
	fmt.Println("Sname = "+s.name)//改了
}

