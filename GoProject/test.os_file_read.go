package main

import (
	"fmt"
	"io"
	"os"
)

/*
打开/关闭文件
*/
func openCloseFile() {
	//只读
	f, err := os.Open("b.txt")
	if err != nil {
		fmt.Printf("f.Name(): %v\n", f.Name())
		errclose := f.Close()
		if errclose == nil {
			fmt.Printf("errclose: %v\n", errclose)
		}
	} else {
		fmt.Printf("f err: %v\n", err)
	}

	//根据第二个参数，可读/写/创建
	//O_CREATE->没有就创建这个文件
	f2, err := os.OpenFile("f2.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Printf("f2.Name(): %v\n", f2.Name())
		errclose2 := f2.Close()
		if errclose2 == nil {
			fmt.Printf("errclose2: %v\n", errclose2)
		}
	} else {
		fmt.Printf("f2 err: %v\n", err)
	}
}

/*
创建文件
*/
func createFile2() {
	f, _ := os.Create("a2.txt") //等价于os.OpenFile("f2.txt", os.O_RDWR|os.O_CREATE, 0755)
	fmt.Printf("f.Name(): %v\n", f.Name())

	//第一个参数""表示在默认目录创建，temp作为文件前缀，文件没有后缀，创建结果：/var/folders/d6/q67f7j5x6dng41zf3k2kqk400000gn/T/temp112274055
	f2, _ := os.CreateTemp("", "temp")
	fmt.Printf("f2.Name(): %v\n", f2.Name())
}

/*
读
循环读 直到结束
*/
func readFile2() {
	//方法1
	fmt.Println("循环读:")
	f, _ := os.Open("b.txt")
	for {
		buf := make([]byte, 3)
		n, err := f.Read(buf)
		if err == io.EOF {
			break
		}
		fmt.Printf("n: %v\n", n)
		fmt.Printf("buf->string: %v\n", string(buf))
	}
	f.Close()

	//方法2 从某处开始读
	fmt.Println("从某处开始读:")
	f, _ = os.Open("b.txt")
	buff := make([]byte, 2)
	n, _ := f.ReadAt(buff, 3)
	fmt.Printf("n: %v\n", n)
	fmt.Printf("buff: %v\n", string(buff))

	//方法3 路径读
	fmt.Println("路径读:")
	de, _ := os.ReadDir("a")
	for _, v := range de {
		fmt.Printf("v.IsDir(): %v\n", v.IsDir())
		fmt.Printf("v.Name(): %v\n", v.Name())
	}

	//方法4 从某处开始读2
	fmt.Println("从某处开始读2:")
	buf2 := make([]byte, 10)
	f.Seek(3, 0)
	n2, _ := f.Read(buf2)
	fmt.Printf("n2: %v\n", n2)
	fmt.Printf("string(buf2): %v\n", string(buf2))
	f.Close()

}

func TestFile() {
	openCloseFile()
	createFile2()
	readFile2()
}
