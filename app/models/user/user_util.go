package user

import (
	"CobraApp/pkg/database"
	"fmt"
	"gorm.io/gorm"
)

// Create 写入数据
func (user *User) Create() int64 {
	result := database.DB.Create(&user)
	return result.RowsAffected
}

// Creates 批量写入数据
func Creates(users []User) int64 {
	result := database.DB.Create(&users)
	return result.RowsAffected
}

// Update 更新单列
func Update(id int, name string) int64 {
	result := database.DB.Model(User{}).Where("id = ?", id).Update("name", name)
	return result.RowsAffected
}

func (user *User) BeforeUpdate(tx *gorm.DB) {
	fmt.Println("BeforeUpdate")
}

// Updates 更新多列
func Updates(id int, data interface{}) int64 {
	result := database.DB.Model(User{}).Where("id = ?", id).Updates(data)
	return result.RowsAffected
}

// Delete 删除单条数据
func Delete(id int64) int64 {
	users := User{ID: id}
	return database.DB.Limit(1).Delete(&users).RowsAffected
}

// Deletes 删除多条数据
func Deletes(idOne int, idTwo int) int64 {
	result := database.DB.Where("id > ? and id < ?", idOne, idTwo).Delete(User{})
	return result.RowsAffected
}

// IsEmailExist 判断 Email 是否已经存在
func IsEmailExist(email string) bool {
	var count int64
	database.DB.Model(User{}).Where("email = ?", email).Count(&count)
	return count > 0
}

// IsPhoneExist 判断 手机号 是否已经存在
func IsPhoneExist(phone string) bool {
	var count int64
	database.DB.Model(User{}).Where("phone = ?", phone).Count(&count)
	return count > 0
}

func GetUserById(id int64) (UserModel User) {
	database.DB.Model(User{}).Where("id = ?", id).Find(&UserModel)
	return
}

func GetUserList(idOne int, idTwo int) (UserModel []User) {
	database.DB.Model(User{}).Where("id > ? and id < ?", idOne, idTwo).Find(&UserModel)
	return
}

func GetOrderById(id int) (UserModel []User) {
	//database.DB.Model(User{}).Where("id > ?", id).First(&UserModel) // SELECT * FROM `users` WHERE id > 0 ORDER BY `users`.`id` LIMIT 1
	//database.DB.Model(User{}).Where("id > ?", id).Last(&UserModel) // SELECT * FROM `users` WHERE id > 0 ORDER BY `users`.`id` DESC LIMIT 1
	database.DB.Model(User{}).Where("id > ?", id).Order("created_at desc").Find(&UserModel) //SELECT * FROM `users` WHERE id > 0 ORDER BY created_at desc
	return
}
