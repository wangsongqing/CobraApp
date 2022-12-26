// Package user 存放用户 Model 相关逻辑
package user

import (
	"CobraApp/app/models"
)

// User 用户模型
type User struct {
	models.BaseModel

	ID       int64  `json:"id"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"-"`
	Phone    string `json:"-"`
	Password string `json:"-"`

	models.CommonTimestampsField
}
