package controllers

import (
	"CobraApp/app/models/user"
	"fmt"
	"github.com/gookit/color"
)

// GetUserInfo 获取用户详情
func GetUserInfo(id int64) {
	userInfo := user.GetUserById(id)
	color.Red.Println(userInfo)
}

// GetUserList 获取用户列表
func GetUserList() {
	userList := user.GetUserList(1, 3)
	for _, v := range userList {
		fmt.Printf("name:%v, email:%v \n", v.Name, v.Name)
	}
}
