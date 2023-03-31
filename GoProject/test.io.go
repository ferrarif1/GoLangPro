package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func testCopy() {
	r := strings.NewReader("hello world")
	_, err := io.Copy(os.Stdout, r) //copy到控制台来输出, ⚠️ ！！！！需要用go run .才能输出！！！！，用vscode按钮无法正常输出
	if err != nil {
		log.Fatal(err)
	}
}

func readBuffer() {
	r := strings.NewReader("hello")
	buf := make([]byte, 10)
	r.Read(buf)
	fmt.Printf("string(buf): %v\n", string(buf))
}

func TestIO() {
	readBuffer()
	testCopy()
}
