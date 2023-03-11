package work

import(
	"fmt"
)
//闭包，返回一个函数
func Bibao() func() {
	i:= 10
	fmt.Printf("%p\n",&i)
	b:= func() {
		fmt.Printf("i addr = %p \n  ",&i)
		i--
		fmt.Println(i)
	}
	return b
}