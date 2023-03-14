package test

import (
	"ginchat/models"
    "fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func TestGorm() {
	db, err := gorm.Open(mysql.Open("root:bugatti123t@tcp(127.0.0.1:3306)/GinChat?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&models.UserBasic{})

	// Create
	user := &models.UserBasic{}
	user.Name = "Tom"
	db.Create(user)

	// Read
	fmt.Println(db.First(user, 1))
	db.First(&user, 1) // find product with integer primary key
	db.First(&user, "Name = ?", "Tom") // find product with code D42

	// Update - update product's price to 200
	db.Model(user).Update("PassWord", "12345")
	// // Update - update multiple fields
	// db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	// db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - delete product
	// db.Delete(&product, 1)
}
