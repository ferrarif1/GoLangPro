package main

import "os"

func writeFile1() {
	f, _ := os.OpenFile("a2.txt", os.O_RDWR|os.O_TRUNC, 0755) //O_TRUNC表示覆盖
	f.Write([]byte("bitcoin wif"))
	f.Close()
}

func writeFile2() {
	f, _ := os.OpenFile("a2.txt", os.O_RDWR|os.O_APPEND, 0755) //O_TRUNC表示覆盖
	f.WriteString(" ethereum")
	f.Close()
}

func writeFile3() {
	f, _ := os.OpenFile("a2.txt", os.O_RDWR, 0755) //O_TRUNC表示覆盖
	f.WriteAt([]byte("tom & jerry"), 16)
	f.Close()
}

func TestWriteFile() {
	writeFile1()
	writeFile2()
	writeFile3()
}
