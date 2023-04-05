package main

import "fmt"

func TestInterface2() {
	/*
		定义了一个空接口x，并将其赋值为一个字符串类型的值"Hello"。接着使用类型断言将x转换为string类型，从而可以直接操作其值。
	*/
	var x interface{} = "Hello"
	s := x.(string)
	fmt.Println(s) // 输出：Hello

	/*
		需要注意的是，如果类型断言失败，会引发panic异常。为了避免这种情况，我们可以使用类型断言的另一种形式，即带有第二个返回值的类型断言，该返回值表示类型断言是否成功。例如：
	*/
	var x1 interface{} = 123
	if s1, ok := x1.(string); ok {
		fmt.Println(s1)
	} else {
		fmt.Println("类型断言失败")
	}

	var ori = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// moveDuplicatesToEnd(&ori,4)
	moveDuplicatesToEnd2(ori, 4)
	fmt.Printf("ori: %v\n", ori)
}

func moveDuplicatesToEnd(nums *[]int, target int) []int {
	i, j := 0, len(*nums)-1
	for i <= j {
		if (*nums)[i] == target {
			for i < j {
				(*nums)[i] = (*nums)[i+1]
				i++
			}
			(*nums)[i] = target
			break
		} else {
			i++
		}
	}
	return (*nums)
}

func moveDuplicatesToEnd2(nums []int, target int) []int {
	i, j := 0, len(nums)-1
	for i <= j {
		if nums[i] == target {
			for i < j {
				nums[i] = nums[i+1]
				i++
			}
			nums[i] = target
			break
		} else {
			i++
		}
	}
	return nums
}
