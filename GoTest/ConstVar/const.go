package ConstVar

const(
	PI = 3.14
	E= 2.178
)

var City = "bj" //全局变量必须通过var/const定义，不能用:=

func Add(a, b int) float64{
	var c= 4.1
	d:=2.
	return c+d
}