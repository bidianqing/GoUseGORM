package main

import (
	"database/sql"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/demo?charset=utf8mb4"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	var users []User
	db.Raw("select * from tb_user where id = @Id", sql.Named("Id", 1)).Scan(&users)

	fmt.Println(users)
}

type User struct {
	Id   int
	Name string
}
