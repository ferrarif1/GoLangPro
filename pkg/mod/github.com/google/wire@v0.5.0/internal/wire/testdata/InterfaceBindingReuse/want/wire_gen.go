// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

// Injectors from wire.go:

func injectFooBar() FooBar {
	bar := provideBar()
	fooBar := provideFooBar(bar, bar)
	return fooBar
}