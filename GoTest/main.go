package main

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"workspace/ConstVar"
	"workspace/Utils"
	"workspace/work"
)

var (
	A = 3 //大写，其他package也可访问
	b = 4 //小写，仅本package可访问
)

// 自定义类型
type signal uint8
type add func(a, b int) int
type ms map[string]string

// 给自定义类型ms添加方法
func (self ms) Hello() {
	fmt.Printf("%s\n", self["bob"])
}

func lit() {
	fmt.Printf("\n lit = %t\n", 04 == 4.00)
	fmt.Printf("%v\n", 3.5i)
	var r rune = '众'
	fmt.Printf("%d %x", r, r)
}

func str_func() {
	s := "Hhiebiexn"
	arr := strings.Split(s, "i")
	for _, ele := range arr {
		fmt.Printf(ele + "\n")
	}
	//字符串拼接
	s1 := "as"
	s2 := "asq"
	s3 := "aswe"

	S1 := s1 + s2 + s3
	S2 := strings.Join([]string{s1, s2, s3}, "-")

	sb := strings.Builder{}
	sb.WriteString(s1)
	sb.WriteString(s2)
	S4 := sb.String()

	fmt.Println(S1)
	fmt.Println(S2)
	fmt.Println(S4)
}

func Concat(arr []int) string {
	var sb strings.Builder
	for _, ele := range arr {
		sb.WriteString(strconv.Itoa(ele)) //strconv.Itoa:整形转字符串
		//sb.WriteString(strconv.FormatInt(int64(ele), 10))//与上面等价
		sb.WriteString(" ")
	}
	//拼接后末尾多了个空格，可用1）切片将其删去 或者 2）用Trim删掉
	s := sb.String()
	return s[0 : len(s)-1] //
	// return strings.Trim(sb.String()," ")
}

func array_func() {
	arr := [...]int{1, 2, 3}
	for i, ele := range arr {
		fmt.Printf("index = %d, ele = %d\n", i, ele)
	}

	arr2 := [...][3]uint{{1}, {2, 3, 4}} //自动推断为两行
	for row, array := range arr2 {
		for col, ele := range array {
			fmt.Printf("a[%d][%d] = %d \n", row, col, ele)
		}
	}
}

func slice_init() {
	fmt.Printf("*******slice_init*******\n")
	/*
			在Go语言中，数组和切片（slice）是两种不同的数据类型，它们的区别在以下几个方面：

		长度不同：数组的长度在定义时就已经确定，并且不可更改；切片的长度可以在运行时动态增加或缩减。

		内存分配方式不同：数组是在栈上分配内存，而切片则是在堆上分配内存。

		传递方式不同：数组是值传递，传递时会复制整个数组；而切片是引用传递，传递时只复制指向底层数组的指针、长度和容量等元信息。

		使用方式不同：数组的使用通常是通过下标来访问和操作数组元素，而切片则支持更灵活的方式，如切片操作、追加操作、复制操作等。

		声明方式不同：数组的声明方式是 var arr [N]T，其中 N 表示数组长度，T 表示元素类型；切片的声明方式是 var slice []T，其中 T 表示元素类型，长度和容量都为0。

		需要注意的是，在Go语言中，切片并不是一个完整的数组，而只是对一个底层数组的引用。因此，对切片的修改会影响到底层数组的内容，也会影响到其它引用该底层数组的切片或数组。
	*/
	s2d := [][]int{
		{1},
		{1, 3, 5},
	}
	fmt.Println(len(s2d))
	fmt.Println(len(s2d[0]))

	var s []int
	fmt.Printf("len = %d, cap = %d\n", len(s), cap(s))
	s = []int{}
	fmt.Printf("len = %d, cap = %d\n", len(s), cap(s))
	s = make([]int, 3, 5)
	fmt.Printf("len = %d, cap = %d\n", len(s), cap(s))
	s = make([]int, 3)
	fmt.Printf("len = %d, cap = %d\n", len(s), cap(s))
	s = []int{1, 2, 3, 4}
	fmt.Printf("len = %d, cap = %d\n", len(s), cap(s))

	s2 := []int{1, 2, 3, 4}
	s2 = append(s, 100)
	fmt.Println(s2)

	update_slice(&s2)
	fmt.Println(s2[0]) //注意这里是正常访问
}

func GenSlice(n int) []int {
	var arr []int
	arr = make([]int, 0, 10)
	for i := 0; i < 100; i++ {
		arr = append(arr, rand.Intn(128))
	}
	return arr
}

func CountUniqueSlice(arr []int) int {
	m := make(map[int]bool, len(arr))
	for _, ele := range arr {
		m[ele] = true
	}
	return len(m)
}

func update_slice(arr *[]int) {
	(*arr)[0] = 100 //注意这里arr是指针，它里面存的是一个地址，要先用（*arr）转换成真正的那个内存地址
	fmt.Printf("arr[0] upadate to:%d", (*arr)[0])
}

// 管道-环形队列
func channel_func() {
	var ch chan int
	ch = make(chan int, 8)
	fmt.Println("ch =")
	fmt.Println(ch)
	fmt.Printf("Type of ch is %T \n", ch)
	ch <- 1 // 放入元素 send
	ch <- 2
	ch <- 43
	ch <- 239
	v := <-ch //取出recv
	fmt.Println(v)

	close(ch) //关闭管道
	//若不close就遍历，如果有其他线程继续加元素，则会出差错，导致遍历出现问题

	//遍历1
	for ele := range ch {
		fmt.Printf("ele1 = %v \n", ele)
	}

	//遍历一遍后元素会出队，故遍历2 L为0
	//遍历2
	L := len(ch)
	fmt.Println(L)
	for i := 0; i < L; i++ {
		ele := <-ch
		fmt.Printf("ele2 = %v \n", ele)
	}
}

func func_map() {
	var m map[string]int
	m = make(map[string]int)
	m = make(map[string]int, 5)
	m2 := map[string]int{"key1": 1, "key2": 2}
	m["ade"] = 12

	delete(m, "ade")

	fmt.Println(len(m))

	if value, exists := m2["key1"]; exists {
		fmt.Println(value)
	} else {
		fmt.Println("error")
	}

	//遍历
	for key, value := range m {
		fmt.Println(key, value)
	}

	fmt.Println(m, m2)
}

func BinaryFormat(n int32) string {
	a := uint32(n)
	sb := strings.Builder{}
	c := uint32(math.Pow(2, 31))
	// fmt.Printf("c = %b \n",c)
	for i := 0; i < 32; i++ {
		// fmt.Printf("a = %b\n",a)
		// fmt.Printf("c = %b\n",c)
		// fmt.Printf("c&a = %b\n",c&a)
		if c&a > 0 {
			sb.WriteString("1")
		} else {
			sb.WriteString("0")
		}
		c >>= 1
		// fmt.Printf("第%d轮 %s\n", i, sb.String())
	}
	// fmt.Printf("%s\n", sb.String())
	return sb.String()
}

/*函数作为数据类型*/
// func func_arg(f func(a, b int)int, b int)int{
// 	a:= 2*b
// 	return f(a,b)
// }
//下面写法更简洁：
type foo func(a, b int) int

func func_arg(f foo, b int) int {
	a := 2 * b
	return f(a, b)
}

// 匿名函数
var sum = func(a, b int) int {
	return a + b
}

type addint func(a, b int) int

func op(f addint, a int) int {
	b := 2 * a
	return f(a, b)
}

/*
输出顺序：
*/
func basic() (i int) {
	i = 9
	fmt.Println("A") //先执行
	defer func(i int) {
		fmt.Println(i)
	}(i) //第7执行，输出9，因为i在defer注册时已经传入了
	defer func() {
		fmt.Println(i)
	}() //第6执行，输出5 ,这里的i执行时才会取值所以得到5 注意：这里需要调用func,故在最后加（）
	defer fmt.Println(i)   //第5执行 ，输出9，这个i在一开始就赋值了，所以还是9
	defer fmt.Println("B") //第4执行，defer是逆序的，后注册的先执行
	fmt.Println("C")       //先执行
	defer fmt.Println("D") //第3个执行
	fmt.Println("E")       //先执行
	return 5               //第2执行
}

/*
aaaaaaaa3
aaaaaaaa
panic: runtime error: integer divide by zero
*/
//panic类似于抛出异常中止程序
func defer_panic() {
	defer fmt.Println("aaaaaaaa")
	n := 0
	defer func() {
		fmt.Println(2 / n) //此处将有panic,但不会影响下一个defer执行
		defer fmt.Println("aaaaaaaa2")
	}()
	defer fmt.Println("aaaaaaaa3")
}

func main() {
	fmt.Println("Test fmt")

	var a int
	var b int
	a = 7
	b = 3
	c := a * b
	d := a / b
	e := (a == b)

	fmt.Printf("%d - %d - %t\n", c, d, e)
	fmt.Printf("%b", a)

	const (
		a1 = iota //0
		b1 = 30   //30
		c1        //30
		d1        //30
	)
	const (
		a2 = iota //0
		b2        //1
		c2        //2
		d2        //3
	)

	lit()

	//ConstVar
	var city = ConstVar.GetCity()
	fmt.Printf("city = %s\n", city)

	// BinaryFormat(0)
	// BinaryFormat(1)
	// BinaryFormat(2)
	BinaryFormat(3)

	Utils.Defaultvalue()
	Utils.Scope()
	var bbb0 = 100
	var bbb1 byte = 100
	fmt.Printf("Type of bbb0 is %T \n", bbb0)
	fmt.Printf("Type of bbb1 is %T \n", bbb1)

	var r rune = 'a'
	fmt.Printf("%d  %c \n", r, r)

	var e1 error
	e1 = errors.New("divide by zero")
	fmt.Printf("%v\n", e1)
	fmt.Printf("%+v\n", e1)
	fmt.Printf("%#v\n", e1)

	type User struct {
		Name string
		Age  int
	}
	us := User{"Admin", 12}
	fmt.Printf("%v\n", us)  //{Admin 12}
	fmt.Printf("%+v\n", us) //{Name:Admin Age:12}
	fmt.Printf("%#v\n", us) //main.User{Name:"Admin", Age:12}

	s := "张"
	arr := []rune(s)
	brr := []byte(s)
	for _, ele := range arr {
		fmt.Printf("%d\n", ele)
		fmt.Printf("%c\n", ele)
	}
	fmt.Printf("arr len %d, s len %d \n", len(arr), len(s))
	for _, ele := range brr { //for range遍历都会拷贝一份
		fmt.Printf("%d\n", ele)
		fmt.Printf("%c\n", ele)
	}
	fmt.Printf("arr len %d, s len %d \n", len(brr), len(s))

	crr := []byte{49, 228, 184, 128}
	fmt.Printf("%s \n", crr) //"1-"

	str_func()
	array_func()
	slice_init()
	channel_func()
	Utils.FuncB()

	work.Main2() // build workspace后才可用 用：1）go build 2）go run workspace
	work.Main3()

	q1 := append(crr, "abc"...) // string->[]byte
	fmt.Println(q1)

	fmt.Println(sum(1, 4))

	slc := make([]addint, 0, 5)
	slc = append(slc, sum, sum, sum) //slc里面放三个这种函数
	fmt.Println(slc)
	//闭包是引用了自由变量的函数
	work.Bibao()() //.Bibao()这层取到func b,再加个()以调用func b
	work.Bibao()() //每次运行func b都是新的 不会改变i
	work.Bibao()()

	bf := work.Bibao()
	bf() // bf是相同的，会改变i
	bf()
	//defer延迟执行
	basic()

	// panic(0) //主线程panic 会导致中止，后面不再执行
	// defer_panic()

	//error
	if res, err := work.Divide(0, 3); err != nil {
		fmt.Printf(err.Error())
	} else {
		fmt.Printf("res = %d\n", res)
	}

	//recover
	work.Soo(2, 3)

	//interface
	work.InterfaceTest()
	work.InterfaceTest2()

	//assert
	work.AssertTest()
}
