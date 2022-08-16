package controllers

import (
	"CobraApp/app/models/user"
	"github.com/gookit/color"
)

func GetUserInfo(id int64) {
	userInfo := user.GetUserById(id)
	color.Red.Println(userInfo)
}
