package main

import (
	"errors"
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

	ipAddress := "127.0.0.1"
	port := "3306"
	userName := "root"
	password := "123456"
	dbName := "study"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True", userName, password, ipAddress, port, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("err:", err)
		panic("failed to connect database")
	}

	fmt.Println("connect database success")

	fmt.Println("start transaction")

	te := db.Transaction(func(tx *gorm.DB) error {
		// Migrate the schema
		err := tx.AutoMigrate(&Product{})
		if err != nil {
			return err
		}

		tag := "test_004"

		// Create
		tx.Create(&Product{Code: tag, Price: 100})

		//// Read
		var product Product
		tx.First(&product, "code = ?", tag) // find product with code D42
		//
		// Update - update product's price to 200
		tx.Model(&product).Update("Price", 300)
		//// Update - update multiple fields
		tx.Model(&product).Updates(Product{Price: 200, Code: tag}) // non-zero fields
		tx.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": tag})
		//
		// Delete - delete product
		//db.Delete(&product, 1)
		return errors.New("test error")
		//return nil
	})

	fmt.Println("transaction error", te)

	fmt.Println("complete transaction")

}
