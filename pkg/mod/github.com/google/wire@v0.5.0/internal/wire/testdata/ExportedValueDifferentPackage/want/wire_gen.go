// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"os"
)

// Injectors from wire.go:

func injectedFile() *os.File {
	file := _wireFileValue
	return file
}

var (
	_wireFileValue = os.Stdout
)
