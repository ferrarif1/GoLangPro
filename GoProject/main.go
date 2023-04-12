package main

import (
	"GoProject/user"
	"fmt"
	"math/rand"
	"strings"
)

// 使用欧几里得算法求最大公约数
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func main() {

	str := "zhongGuo_ZuiMeiLiDe_DiFang_ShiNaLi?"
	s1 := strings.Fields(str) //将会利用 1 个或多个空白符号来作为动态长度的分隔符将字符串分割成若干小块
	fmt.Println("-------s1------------", s1)
	for _, val := range s1 {
		fmt.Println(val + "\n")
	}
	s2 := strings.Split(str, "_") //拆分字符串
	fmt.Println("-------s2 len------------", len(s2))
	for i := 0; i < len(s2); i++ {
		fmt.Println(s2[i])
	}
	fmt.Println("-----------------")
	for _, val := range s2 {
		fmt.Println(val)
	}
	s3 := strings.Join(s2, ",") //拼接字符串
	fmt.Println("------s3-----------", s3)

	for i := 0; i < 5; i++ {
		a := rand.Intn(10) //生成10以内的随机数
		fmt.Println(a)
	}
	for i := 0; i < 5; i++ {
		a := rand.Int() //生产随机数
		fmt.Println(a)
	}

	user.Hello()
	// a := 200
	// p := &a
	// fmt.Printf("p: %v\n", p)
	// //TestString()

	// var str string
	// fmt.Print("请输入一段字符串：")
	// fmt.Scan(&str)

	// m := make(map[rune]int)
	// for _, r := range str {
	// 	m[r]++
	// }

	// var maxRune rune
	// var maxCount int
	// for r, count := range m {
	// 	if count > maxCount {
	// 		maxRune = r
	// 		maxCount = count
	// 	}
	// }

	//fmt.Printf("出现次数最多的字符是 %c，出现了 %d 次。\n", maxRune, maxCount)
	//控制台输入数据
	// var name string
	// var age string
	// fmt.Println("Please Input Name & age")
	// fmt.Scan(&name, &age)
	// fmt.Printf("name: %v\n", name)
	// fmt.Printf("age: %v\n", age)

	// //slice crud
	//TestSlice()

	// //map
	//TestMap()

	// //func
	// TestFunc()

	// //point array
	// TestPointArr()

	// //struct
	// TestStruct()

	// ProjTest()

	// InterfaceTest()

	// TestNewFunc()

	/*
	   并发编程：start***********
	*/
	// TestRoutine()

	//TestChannel()

	// TestWaitGroup()

	// TestRunTime()

	// TestMutex()

	// TestChannelIterate()

	// TestSelectSwitch()

	// TestTimer()

	//TestTicker()

	// TestAtomic()
	//TestAtomic2()

	/*
	   并发编程：end***********
	   golang标准库os模块：start***********
	*/

	//TestOSFile()
	//TestFile()
	// TestWriteFile()

	// TestIO()
	/*
	   golang标准库os模块：end***********
	*/

	/*
		牛客网题目：
	*/
	//TestNewCoder1()
	/*
		算法：
	*/
	//TestLRU()

	//TestShen()

	//TestInterface2()

}
