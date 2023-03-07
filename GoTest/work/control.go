package work

import(
	"fmt"
)

func switch_type()  {
	var num interface{} = 5.5
	switch num.(type) {
	case int:
		fmt.Printf("num is int %d",num)//num已被强制转换为int
	case float64:
		fmt.Printf("num is float64")
    default:
		fmt.Printf("num is other")
	}
   //与上面等价
	switch num.(type) {
	case int:
		value := num.(int)
		fmt.Printf("num is int %d",value)
	case float64:
		value := num.(float64)
		fmt.Printf("num is float64 %f",value)
    default:
		fmt.Printf("num is other")
	}
}

func square(a interface{}) interface{} {
	switch value := a.(type) {
	case int:
		return value * value
	case byte:
		return value * value
	case float32:
		return value * value
	default:
	    return nil
	}
}

func Main3(){
	 switch_type()
}