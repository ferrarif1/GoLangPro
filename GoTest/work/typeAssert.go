package work

import "fmt"
//类型断言


func assert(i interface{})  {
	switch v:= i.(type){
	case int:
		// fmt.Printf("\n%d\n", v)
		// if v, ok := i.(int); ok {//断言是int
		// 	fmt.Printf("\n%d\n", v)
		// }
		v1 := i.(int)
		fmt.Printf("\n%d\n", v1)
	case float64:
		fmt.Printf("%f\n", v)
    default:
		fmt.Printf("%T - %v\n", v, v)
	}
}



func AssertTest()  {
	var i interface{}

    var a int
	var b float64
	var c byte

	i = a
    assert(i)

	i = b
    assert(i)

	i = c
    assert(i)

}