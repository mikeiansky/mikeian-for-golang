package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql" // MySQL 驱动（仅导入，用于注册驱动）
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {

	ipAddress := "192.168.31.109"
	port := "3306"
	userName := "root"
	password := "123456"
	dbName := "test_db"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True", userName, password, ipAddress, port, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("err:", err)
		panic("failed to connect database")
	}

	fmt.Println("connect database success")
	// Migrate the schema
	err2 := db.AutoMigrate(&Product{})
	if err2 != nil {
		return
	}

	// Create
	p1 := &Product{Code: "D42", Price: 100}
	db.Create(p1)
	fmt.Println(p1, "p1.ID:", p1.ID)

	//// Read
	var product Product
	db.First(&product, 2)                 // find product with integer primary key
	db.First(&product, "code = ?", "F42") // find product with code D42
	//
	// Update - update product's price to 200
	db.Model(&product).Update("Price", 300)
	//// Update - update multiple fields
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
	//
	// Delete - delete product
	db.Delete(&product, 1)
}
