package main

import "fmt"

func Map1() {
	var m1 map[string]string
	m1 = make(map[string]string)
	fmt.Printf("m1: %v\n", m1)
	fmt.Printf("m1: %T\n", m1)
}

func Map2() {
	var m1 = map[string]string{"name": "tom", "age": "15"}
	fmt.Printf("m1: %v\n", m1)
	fmt.Printf("m1: %#v\n", m1)

	m2 := make(map[string]string)
	m2["name"] = "tom"
	fmt.Printf("m2: %v\n", m2)
}

func Map3() {
	var m1 = map[string]string{"name": "tom", "age": "15"}
	var key = "name"
	var value = m1[key]
	fmt.Printf("value: %v\n", value)
}

func Map4() {
	var m1 = map[string]string{"name": "tom", "age": "15"}
	var key1 = "name"
	var key2 = "name2"
	v, ok := m1[key1]
	v2, ok2 := m1[key2]
	fmt.Printf("v: %v\n", v)
	fmt.Printf("ok: %v\n", ok)
	fmt.Printf("v2: %v\n", v2)
	fmt.Printf("ok2: %v\n", ok2)
}

func DeleteEle() {
	m := map[string]int{}
	m["sdd"] = 1
	fmt.Printf("m: %v\n", m)
	delete(m, "sdd")
	fmt.Printf("m: %v\n", m)
}

func MapIterate() {
	var m1 = map[string]string{"name": "tom", "age": "15", "email": "shit@111.com"}
	for k, v := range m1 {
		fmt.Printf("k: %v\n", k)
		fmt.Printf("v: %v\n", v)
	}
}

func TestMap() {
	Map1()
	Map2()
	Map3()
	Map4()
	MapIterate()
	DeleteEle()
}
