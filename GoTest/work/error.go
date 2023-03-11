package work

import (
	"errors"
	"fmt"
)

func Divide(a ,b int)(int, error) { //go语言没有try catch,提倡返回error
	if a == 0 {
	   return -1, errors.New("divide by zero!!!!")
	}
	return a/b, nil
}

//自定义error
type PathError struct{
	path string
	op string
	createTime string
	msg string
}

func (err PathError) Error() string{
	return err.path+ ": " + err.op + " " + err.createTime + err.msg
}

func NewPathError(a, b, c, d string) PathError  {
	return PathError{
		path :a,
		op: b,
		createTime: c,
		msg: d}
}

//使用recover阻断panic执行,这样就不会中止，只会输出错误
func Soo(a,b int){
	defer func() {
		if err := recover(); err != nil{
			fmt.Printf("err = %v\n", err)
		}
	}()
	panic(NewPathError(
		"path",
		"openfile",
		"2023-3-3",
		"message"))
}

