package main

import (
	app "ginchat/router"
)

func main() {
	/*
		链接mysql
	*/
	// test.TestGorm()
	
	/*
	    gin框架
		本地访问 http://localhost:8081/index 将返回service.GetIndex返回的值
	*/
	r := app.Router()
	r.Run(":8081")
}
