package links

import (
	"CobraApp/app/models"
	"CobraApp/pkg/database"
)

// Links User 用户模型
type Links struct {
	models.BaseModel

	Name string `json:"name,omitempty"`
	Url  string `json:"url,omitempty"`

	models.CommonTimestampsField
}

func (Links *Links) Create() {
	database.DB.Create(&Links)
}
