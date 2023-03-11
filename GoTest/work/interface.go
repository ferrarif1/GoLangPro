package work

import "fmt"

type Transporter interface{
	move(string, string)(int, error)
    say(int)
}

type Car struct{

}

func (_ Car) move(string, string)(int, error) {
       return 1,nil
}

func (car Car) say(int) {
     fmt.Println("car said")
}

func InterfaceTest()  {
	var t Transporter
	car := Car{}
	t = car
	t.say(1)
    
}
/*
interface 赋值
*/
type Ship struct {}

func (_ Ship) move(string, string)(int, error) {
	return 1,nil
}

func (car Ship) say(int) {
  fmt.Println("ship said")
}



func (car Car)whistle(n int)int{
    return 1
}

func (ship *Ship)whistle(n int)int{
	return 2
}

func InterfaceTest2(){
	car := Car{}
	ship := Ship{}
	var tra Transporter
	tra = car
	tra = &car//值实现的方法，指针也实现了
	tra = &ship//只有指针实现了
	tra.say(0)//直接调用，不必*tra 
}
