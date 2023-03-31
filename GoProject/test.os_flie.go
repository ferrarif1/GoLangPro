package main

import (
	"fmt"
	"os"
	"time"
)

// 创建文件
func createFile() {
	f, err := os.Create("a.txt")
	if err != nil {
		fmt.Printf("createFile err: %v\n", err)
	} else {
		fmt.Printf("createFile f.Name(): %v\n", f.Name())
	}
}

// 创建目录
func createDir() {
	err := os.MkdirAll("test/a/b", os.ModePerm)
	if err != nil {
		fmt.Printf("createDir err: %v\n", err)
	} else {
		fmt.Println("createDir success")
	}
}

// 删除目录
func removeDir() {
	err := os.RemoveAll("test")
	if err != nil {
		fmt.Printf("removeDir err: %v\n", err)
	} else {
		fmt.Println("removeDir success")
	}
}

// 获取工作目录
func getWd() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("getWd err: %v\n", err)
	} else {
		fmt.Printf("getWd dir: %v\n", dir)
	}
}

// 修改工作目录
func chWd() {
	//改成GoTest
	err := os.Chdir("/Users/zhangyuanyi/Downloads/GoLangPro/GoTest")
	if err != nil {
		fmt.Printf("chWd err: %v\n", err)
	} else {
		dir, _ := os.Getwd()
		fmt.Printf("chWd dir: %v\n", dir)
	}
	//改回来
	err2 := os.Chdir("/Users/zhangyuanyi/Downloads/GoLangPro/GoProject")
	if err != nil {
		fmt.Printf("ch back Wd err: %v\n", err2)
	} else {
		dir, _ := os.Getwd()
		fmt.Printf("ch back Wd dir: %v\n", dir)
	}
}

// 重命名
func rename() {
	err := os.Rename("a.txt", "b.txt")
	if err != nil {
		fmt.Printf("rename err: %v\n", err)
	} else {
		fmt.Println("rename a.txt")
	}
}

// 读
func readFile() {
	b, err := os.ReadFile("b.txt")
	if err != nil {
		fmt.Printf("readFile err: %v\n", err)
	} else {
		/*
			这里可以直接用b，或者b[:]来获得切片b的所有元素
		*/
		fmt.Printf("readFile b: %v\n", string(b))
	}
}

// 写
func writeFile() {
	s := "Hello"
	os.WriteFile("b.txt", []byte(s), os.ModePerm)
}

func TestOSFile() {
	createFile()
	createDir()
	time.Sleep(time.Second * 2)
	removeDir()
	getWd()
	chWd()
	rename()
	writeFile()
	readFile()
}
