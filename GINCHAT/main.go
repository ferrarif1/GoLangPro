package main

import (
	"ginchat/router"
	utils "ginchat/utils"
)

func main() {

	utils.InitConfig()
	utils.InitMySQL()

	// DB := utils.InitMySQL()
	// data := make([]*models.UserBasic, 10)
	// DB.Find(&data)
	// for _, v := range data {
	// 	fmt.Println(v)
	// }

	/*
		链接mysql
	*/
	// test.TestGorm()

	/*
		    gin框架
			本地访问 http://localhost:8081/index 将返回service.GetIndex返回的值
	*/
	r := router.Router()
	r.Run(":8081")
}
