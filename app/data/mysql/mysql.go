package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql" // MySQL 驱动（仅导入，用于注册驱动）
)

// 定义一个用户结构体，用于映射查询结果（可选，用于读取数据展示）
type User struct {
	ID    int
	Name  string
	Email string
	Age   int
}

func main() {
	fmt.Println("app start ... ")
	// ✅ 1. 连接 MySQL 数据库
	// 格式：用户名:密码@tcp(主机:端口)/数据库名?charset=utf8mb4&parseTime=True
	ipAddress := "192.168.31.109"
	port := "3306"
	userName := "root"
	password := "123456"
	dbName := "test_db"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True", userName, password, ipAddress, port, dbName)
	//dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True", userName, password, ipAddress, port, dbName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("无法连接到 MySQL: %v", err)
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db) // 确保程序退出前关闭数据库连接

	// ✅ 2. 测试连接是否正常
	err = db.Ping()
	if err != nil {
		log.Fatalf("MySQL 连接测试失败: %v", err)
	}
	fmt.Println("✅ MySQL 连接成功！")

	//// ✅ 3. 创建表（如果表不存在）—— 初始化数据表 users
	//createTableSQL := `
	//CREATE TABLE IF NOT EXISTS users (
	//	id INT AUTO_INCREMENT PRIMARY KEY,
	//	name VARCHAR(100) NOT NULL,
	//	email VARCHAR(100),
	//	age INT
	//);`
	//_, err = db.Exec(createTableSQL)
	//if err != nil {
	//	log.Fatalf("创建表失败: %v", err)
	//}
	//fmt.Println("✅ 数据表 users 已准备（或已存在）")
	fmt.Println(time.Now())
	// ✅ 4. 插入数据（INSERT）
	insertSQL := `INSERT INTO time_test (create_time, create_at) VALUES (?, ?)`
	ret, err := db.Exec(insertSQL, time.Now().UTC().Format(time.DateTime), time.Now().UTC().Format(time.DateTime))
	if err != nil {
		log.Fatalf("插入数据失败: %v", err)
	}

	// 3. 读取时间数据
	var (
		createTime time.Time
		createAt   time.Time
	)
	id, _ := ret.LastInsertId()
	selectSQL := fmt.Sprintf(`SELECT create_time, create_at FROM time_test WHERE id = %v`, id)
	err = db.QueryRow(selectSQL).Scan(&createTime, &createAt)
	if err != nil {
		log.Fatalf("读取数据失败: %v", err)
	}

	// 4. 打印结果对比
	fmt.Println("原始插入时间 (UTC):", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println("读取到的时间 (本地时区):", createTime.Format("2006-01-02 15:04:05"))
	fmt.Println("createTime 时区信息:", createTime.Location().String())
	fmt.Println("createTime:", createTime)
	fmt.Println("createAt 时区信息:", createAt.Location().String())
	fmt.Println("createAt:", createAt)

	//// 获取插入的 ID（可选）
	//id, _ := result.LastInsertId()
	//fmt.Printf("✅ 插入成功，用户 ID: %d\n", id)
	//
	//// ✅ 5. 查询数据（SELECT）
	//rows, err := db.Query("SELECT id, name, email, age FROM users")
	//if err != nil {
	//	log.Fatalf("查询数据失败: %v", err)
	//}
	//defer func(rows *sql.Rows) {
	//	err := rows.Close()
	//	if err != nil {
	//		panic(err)
	//	}
	//}(rows)
	//
	//fmt.Println("\n📋 当前所有用户数据：")
	//for rows.Next() {
	//	var u User
	//	err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Age)
	//	if err != nil {
	//		log.Fatalf("扫描行数据失败: %v", err)
	//	}
	//	fmt.Printf("ID: %d, Name: %s, Email: %s, Age: %d\n", u.ID, u.Name, u.Email, u.Age)
	//}
	//
	//// 检查遍历过程中是否有错误
	//if err = rows.Err(); err != nil {
	//	log.Fatalf("遍历结果集出错: %v", err)
	//}
	//
	//// ✅ 6. 更新数据（UPDATE）
	//updateSQL := `UPDATE users SET age = ? WHERE name = ?`
	//_, err = db.Exec(updateSQL, 26, "Alice")
	//if err != nil {
	//	log.Fatalf("更新数据失败: %v", err)
	//}
	//fmt.Println("\n✅ 更新成功：将 Alice 的年龄更新为 26")
	//
	//// ✅ 7. 再次查询，查看更新后的数据
	//rows, err = db.Query("SELECT id, name, email, age FROM users")
	//if err != nil {
	//	log.Fatalf("再次查询失败: %v", err)
	//}
	//defer func(rows *sql.Rows) {
	//	err := rows.Close()
	//	if err != nil {
	//		panic(err)
	//	}
	//}(rows)
	//
	//fmt.Println("\n📋 更新后的用户数据：")
	//for rows.Next() {
	//	var u User
	//	err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Age)
	//	if err != nil {
	//		log.Fatalf("扫描行数据失败: %v", err)
	//	}
	//	fmt.Printf("ID: %d, Name: %s, Email: %s, Age: %d\n", u.ID, u.Name, u.Email, u.Age)
	//}
	//if err = rows.Err(); err != nil {
	//	log.Fatalf("遍历结果集出错: %v", err)
	//}
	fmt.Println("app complete ... ")
}
