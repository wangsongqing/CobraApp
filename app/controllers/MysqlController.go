package controllers

import (
	"CobraApp/app/models/user"
	"CobraApp/pkg/database"
	"errors"
	"fmt"
	"github.com/gookit/color"
	"gorm.io/gorm"
)

type MysqlController struct {
}

// GetUserInfo 获取用户详情
func (mc *MysqlController) GetUserInfo() {
	userInfo := user.GetUserById(1)
	color.Red.Println(userInfo.Name)
}

// GetUserListById 查询排序
func (mc *MysqlController) GetUserListById() {
	//userList := user.GetUserList(0, 3)
	userList := user.GetOrderById(0)
	for _, v := range userList {
		fmt.Println(v.CreatedAt.Format("2006-01-02 15:04:05"))
	}
}

// Create 写入数据
func (mc *MysqlController) Create() {
	userModel := user.User{}
	userModel.Name = "小明"
	userModel.Email = "xiaoming@gmail.com"
	userModel.Phone = "18299999999"
	result := userModel.Create()
	if result > 0 {
		fmt.Println(result)
		fmt.Println("add success")
	}
}

// Creates 批量写入数据
func (mc *MysqlController) Creates() {
	userModel := []user.User{
		{Name: "王武", Email: "wangwu@qq.com", Phone: "12344444"},
		{Name: "王武01", Email: "wangwu01@qq.com", Phone: "1345555"},
	}

	if row := user.Creates(userModel); row <= 0 {
		fmt.Println("写入失败")
		return
	}

	fmt.Println("success")
}

// Update 更新单列
func (mc *MysqlController) Update() {
	result := user.Update(33, "明明01")
	if result > 0 {
		fmt.Println("update success")
	}
}

// Updates 更新多列
func (mc *MysqlController) Updates() {
	data := map[string]interface{}{
		"name":  "明天",
		"email": "mingtian@qq.com",
	}
	if row := user.Updates(33, data); row <= 0 {
		return
	}

	fmt.Println("success")
}

// Delete 删除数据
// DELETE FROM `users` WHERE `users`.`id` = 11
func (mc *MysqlController) Delete() {
	if row := user.Delete(11); row <= 0 {
		fmt.Println("删除失败")
		return
	}

	fmt.Println("success")
}

// Deletes 删除多条数据
// DELETE FROM `users` WHERE id > 8 and id < 11"
func (mc *MysqlController) Deletes() {
	if row := user.Deletes(8, 11); row <= 0 {
		fmt.Println("删除失败")
		return
	}

	fmt.Println("success")
}

// Transaction 执行事物
func (mc *MysqlController) Transaction() {
	tx := database.DB.Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		userData := user.User{
			Name:  "张三01",
			Email: "zhangsan@gmail.com",
			Phone: "1820001",
		}
		if err := tx.Create(&userData).Error; err != nil {
			// 返回任何错误都会回滚事务
			return err
		}

		if err := tx.Model(user.User{}).Where("id = ?", 32).Update("name", "wsq02").RowsAffected; err <= 0 {
			return errors.New("更新失败")
		}

		// 返回 nil 提交事务
		return nil
	})

	fmt.Println(tx)
}

// Transaction01 事物
func (mc *MysqlController) Transaction01() {
	// 开启事物
	tx := database.DB.Begin()
	userModel := user.User{
		ID: 36,
	}
	result := tx.Delete(&userModel)
	if result.RowsAffected <= 0 {
		tx.Rollback() // 回滚事务
		fmt.Println("删除失败")
		return
	}

	if row := tx.Model(user.User{}).Where("id = ?", 40).Updates(map[string]interface{}{"name": "李四", "email": "lisi@qq.com"}).RowsAffected; row <= 0 {
		tx.Rollback() // 回滚事务
		fmt.Println("更新失败")
		return
	}

	tx.Commit() // 提交事物
	fmt.Println("success")
}
