package controllers

import (
	"CobraApp/app/models/user"
	"fmt"
)

func GetUserInfo() {
	userInfo := user.GetUserById(1)
	fmt.Println("userInfo:", userInfo)
}
