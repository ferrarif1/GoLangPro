package main

import (
	"bytes"
	"fmt"
)

func TestString() {
	var buffer bytes.Buffer
	buffer.WriteString("tom")
	buffer.WriteString("sdf")
	fmt.Printf("buffer: %v\n", buffer.String())
}