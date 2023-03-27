package main

import "fmt"

//构造函数

func NewPerson(name string, age int) (*Person, error) {
	if name == "" {
		return nil, fmt.Errorf("name cannot be empty")
	}
	if age < 0 {
		return nil, fmt.Errorf("age cannot be negative")
	}
	return &Person{name: name, age: age}, nil
}

func TestNewFunc() {
	pe, err := NewPerson("John", 28)
	if err == nil {
		fmt.Printf("pe.name: %v\n", pe.name)
	} else {
		fmt.Printf("err: %v\n", err)
	}

}
