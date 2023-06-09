package main

import (
	"fmt"
	"mall/common/database"
	"mall/service/user/model"
)

type User struct {
	Id   int64
	Name string
}

func main() {
	db := database.InitGorm("root:root@tcp(mysql:3306)/go-zero?charset=utf8mb4&parseTime=true")
	var user model.User
	affected := db.Table("user").Where("id = 1").First(&user).RowsAffected
	fmt.Println(user)
	fmt.Println(affected)
}
