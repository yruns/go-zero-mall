package main

import (
	"fmt"
	"github.com/jinzhu/copier"
)

type User struct {
	Id   int64
	Name string
}

func main() {
	u1 := User{
		Id: 1,
	}

	u2 := User{
		Id:   2,
		Name: "yruns",
	}

	err := copier.CopyWithOption(&u2, u1, copier.Option{
		IgnoreEmpty: true,
	})
	fmt.Println(err)
	fmt.Println(u1)
	fmt.Println(u2)
}
