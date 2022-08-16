package controllers

import (
	"CobraApp/app/models/user"
	"fmt"
)

func GetUserInfo(id int64) {
	userInfo := user.GetUserById(id)
	fmt.Println("userInfo:", userInfo)
}
