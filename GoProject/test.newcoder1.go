package main

import "fmt"

type CustomStruct struct{
    A int
}
type Options func(*CustomStruct)

/*
下面函数均返回一个Options类型的函数
*/
func WithAddvalue(value int) Options{
	return func(c *CustomStruct){
	    c.A += value
	}
}

func WithSubvalue(value int) Options{
	return func(c *CustomStruct){
	    c.A -= value
	}
}

func NewCustomStruct(initValue int, opts ...Options) *CustomStruct{
	res := &CustomStruct{
		A:initValue,
	}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func TestNewCoder1()  {
	stru := NewCustomStruct(5, WithAddvalue(-1),WithAddvalue(2),WithSubvalue(2))
	fmt.Printf("stru: %v\n", stru.A)
}