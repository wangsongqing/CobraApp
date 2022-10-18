package links

import (
	"CobraApp/app/models"
)

// Links User 用户模型
type Links struct {
	models.BaseModel

	Name string `json:"name,omitempty"`
	Url  string `json:"url,omitempty"`

	models.CommonTimestampsField
}
